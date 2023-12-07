package async_queue

import (
	"errors"
	"fmt"
	"github.com/MenciusCheng/algorithm-code/utils/log"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"sync"
	"time"
)

// AsyncQueue 异步队列结构体
type AsyncQueue struct {
	// 队列公共参数
	QueueName  string        // 唯一队列名称
	RedisConn  *redis.Client // Redis 连接
	Priorities []int64       // 优先级配置

	// 消费者参数
	Handlers         []Handler // 任务处理方法
	Concurrency      int64     // 并发执行任务数
	RetryMax         int64     // 最大重试次数
	Timeout          int64     // 任务处理超时秒数，超过该时间则失败重试
	Retention        int64     // 任务完成后状态保留秒数
	FailureRetention int64     // 任务失败后状态保留秒数
	Deadline         int64     // 任务最大期限秒数，保留字段暂无使用

	Cache    *Cache     // Redis 管理任务和队列
	lock     sync.Mutex // protects the fields
	consumed bool       // 是否已开启消费
}

func NewAsyncQueue(queueName string, redisConn *redis.Client, options ...OptionFunc) *AsyncQueue {
	c := &AsyncQueue{
		QueueName:        queueName,
		RedisConn:        redisConn,
		Priorities:       GetDefaultPriorities(),
		Concurrency:      DefaultConcurrency,
		Timeout:          DefaultTimeout,
		Retention:        DefaultRetention,
		FailureRetention: DefaultFailureRetention,
	}
	for _, f := range options {
		f(c)
	}

	c.Cache = NewCache(queueName, redisConn)
	c.Cache.Retention = c.Retention
	c.Cache.FailureRetention = c.FailureRetention

	return c
}

// Publish 发布普通任务
func (a *AsyncQueue) Publish(payload []byte, options ...OptionTaskFunc) (*TaskInfo, error) {
	now := time.Now().Unix()
	task := NewTaskInfo(payload, TaskTypeNormal, options...)

	// 校验优先级是否在配置内
	if !a.IsValidPriority(task.Priority) {
		return nil, fmt.Errorf("publish invalid priority:%d", task.Priority)
	}

	msg := TaskInfoToMessage(task)
	if msg.ProcessAt > now {
		if err := a.Cache.Schedule(msg); err != nil {
			return nil, fmt.Errorf("publish schedule failed: %w", err)
		}
	} else {
		if err := a.Cache.Pending(msg); err != nil {
			return nil, fmt.Errorf("publish pending failed: %w", err)
		}
	}

	return task, nil
}

// PublishListDAG 发布单链表关系的DAG任务
func (a *AsyncQueue) PublishListDAG(payloads [][]byte, options ...OptionTaskFunc) ([]*TaskInfo, error) {
	// 构建任务链表
	tasks := make([]*TaskInfo, 0, len(payloads))
	for _, payload := range payloads {
		task := NewTaskInfo(payload, TaskTypeDAG, options...)
		tasks = append(tasks, task)
	}
	for i := 1; i < len(tasks); i++ {
		tasks[i].PreTaskIDs = []string{tasks[i-1].ID}
		tasks[i].InDegree = len(tasks[i].PreTaskIDs)
		tasks[i-1].PostTaskIDs = []string{tasks[i].ID}
	}

	// check
	taskIDMap := make(map[string]bool)
	for _, task := range tasks {
		if taskIDMap[task.ID] {
			return nil, fmt.Errorf("publishListDAG id conflict:%s", task.ID)
		}
		taskIDMap[task.ID] = true

		// 校验优先级是否在配置内
		if !a.IsValidPriority(task.Priority) {
			return nil, fmt.Errorf("publishListDAG invalid priority:%d", task.Priority)
		}
	}

	// 构建并发布DAG任务
	dag := NewDAG()
	for i := 0; i < len(tasks); i++ {
		dag.AddVertex(tasks[i].ID, tasks[i])
	}
	for i := 1; i < len(tasks); i++ {
		dag.AddEdge(tasks[i-1].ID, tasks[i].ID)
	}
	publishTasks, err := a.publishDAG(dag)
	if err != nil {
		return nil, fmt.Errorf("publishDAG failed:%w", err)
	}

	return publishTasks, nil
}

// publishDAG 发布DAG任务
func (a *AsyncQueue) publishDAG(dag *DAG) ([]*TaskInfo, error) {
	if dag == nil || len(dag.Vertexes) == 0 {
		return nil, errors.New("dag task is empty")
	}
	if !dag.CanFinish() {
		return nil, errors.New("invalid dag task loop")
	}

	vertices := dag.BFS()
	tasks := make([]*TaskInfo, 0, len(vertices))
	for _, vertex := range vertices {
		task, ok := vertex.Value.(*TaskInfo)
		if !ok {
			return nil, fmt.Errorf("invalid vertex value")
		}
		tasks = append(tasks, task)
	}

	now := time.Now().Unix()
	for _, task := range tasks {
		msg := TaskInfoToMessage(task)
		if msg.ProcessAt > now {
			if err := a.Cache.Schedule(msg); err != nil {
				return nil, fmt.Errorf("publishDAG schedule failed: %w", err)
			}
		} else {
			if err := a.Cache.DAGing(msg); err != nil {
				return nil, fmt.Errorf("publishDAG daging failed: %w", err)
			}
		}
	}

	return tasks, nil
}

// StartConsuming 启动消费
func (a *AsyncQueue) StartConsuming() error {
	a.lock.Lock()
	defer a.lock.Unlock()

	if a.consumed {
		return errors.New("multiple startConsuming not allowed")
	}
	a.consumed = true

	if a.Concurrency <= 0 {
		return errors.New("concurrency is zero")
	}

	// 根据并发数启动消费者
	for i := int64(0); i < a.Concurrency; i++ {
		go func(consumerID int64) {
			a.consumeApi(consumerID)
		}(i)
	}

	// 轮询到达执行时间的延时队列
	go func() {
		// TODO panic 检查
		a.forwardAll()
	}()

	go func() {
		a.dagForwardAll()
	}()

	return nil
}

// 校验优先级是否在配置内
func (a *AsyncQueue) IsValidPriority(priority int64) bool {
	for _, item := range a.Priorities {
		if item == priority {
			return true
		}
	}
	return false
}

func (a *AsyncQueue) AddHandler(f Handler) {
	if f != nil {
		a.Handlers = append(a.Handlers, f)
	}
}

func (a *AsyncQueue) consumeApi(consumerID int64) {
	tick := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-tick.C:
			a.consumeTask(consumerID)
		}
	}
}

func (a *AsyncQueue) consumeTask(consumerID int64) {
	var resp *TaskMessage
	var err error

	for _, priority := range a.Priorities {
		// 先从高优先级队列获取任务
		resp, err = a.Cache.Active(priority)
		if err != nil {
			log.Error("consumeTask GetMessages error", zap.Error(err), zap.Int64("consumerID", consumerID))
			return
		}
		if resp != nil {
			break
		}
	}
	// 无任务消息，直接返回
	if resp == nil {
		return
	}

	task := TaskMessageToInfo(resp)
	for _, handler := range a.Handlers {
		handler.ProcessTask(task)
	}

	err = a.Cache.Success(resp)
	if err != nil {
		log.Error("Success err", zap.Error(err))
		return
	}
	log.Debug("consumeTask end finish", zap.Int64("consumerID", consumerID))
}

func (a *AsyncQueue) forwardAll() {
	tick := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-tick.C:
			for _, priority := range a.Priorities {
				a.Cache.Forward(priority)
			}
		}
	}
}

func (a *AsyncQueue) dagForwardAll() {
	tick := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-tick.C:
			for _, priority := range a.Priorities {
				a.Cache.DAGForward(priority)
			}
		}
	}
}

type Handler interface {
	ProcessTask(*TaskInfo) error
}

// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers. If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// Handler that calls f.
type HandlerFunc func(*TaskInfo) error

// ProcessTask calls f(t)
func (f HandlerFunc) ProcessTask(t *TaskInfo) error {
	return f(t)
}
