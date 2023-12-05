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

// 队列参数结构体
type AsyncQueue struct {
	// 队列公共参数
	QueueName  string        // 队列名称
	Priorities []int64       // 优先级配置
	RedisConn  *redis.Client // Redis 连接

	// 消费者参数
	Handlers    []Handler // 任务处理方法
	Concurrency int64     // 并发执行任务数
	RetryMax    int64     // 最大重试次数
	Timeout     int64     // 任务处理超时秒数

	lock     sync.Mutex // protects the fields
	consumed bool       // 是否已开启消费
}

func NewAsyncQueue(options ...OptionFunc) *AsyncQueue {
	c := &AsyncQueue{}
	for _, f := range options {
		f(c)
	}

	c.initQueueConfig()
	return c
}

// 初始化默认配置
func (a *AsyncQueue) initQueueConfig() {
	if len(a.Priorities) == 0 {
		a.Priorities = GetDefaultPriorities()
	}
	if a.Concurrency == 0 {
		a.Concurrency = DefaultConcurrency
	}
	if a.RetryMax == 0 {
		a.RetryMax = DefaultRetryMax
	}
	if a.Timeout == 0 {
		a.Timeout = DefaultTimeout
	}
}

// 启动消费
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

	// TODO 启动守护协程

	// 根据并发数启动消费者
	for i := int64(0); i < a.Concurrency; i++ {
		go func(consumerID int64) {
			a.consumeApi(consumerID)
		}(i)
	}
	return nil
}

// 发布任务
func (a *AsyncQueue) Publish(task *TaskInfo) error {
	// TODO 校验字段

	now := time.Now().Unix()

	initTaskInfo(task, TaskTypeNormal)

	msg := TaskInfoToMessage(task)
	if msg.ProcessAt > now {
		if err := a.Schedule(msg); err != nil {
			return fmt.Errorf("publish schedule failed: %w", err)
		}
	} else {
		if err := a.Pending(msg); err != nil {
			return fmt.Errorf("publish pending failed: %w", err)
		}
	}

	return nil
}

// 发布链表任务
func (a *AsyncQueue) PublishList(tasks []*TaskInfo) error {
	now := time.Now().Unix()

	for i := 0; i < len(tasks); i++ {
		initTaskInfo(tasks[i], TaskTypeDAG)
		if i > 0 {
			tasks[i].PreTaskIDs = []string{tasks[i-1].ID}
		}
	}

	for i := 0; i < len(tasks); i++ {
		msg := TaskInfoToMessage(tasks[i])
		if msg.ProcessAt > now {
			if err := a.Schedule(msg); err != nil {
				return fmt.Errorf("publishList schedule failed: %w", err)
			}
		} else {
			if err := a.DAGing(msg); err != nil {
				return fmt.Errorf("publishList waitpre failed: %w", err)
			}
		}
	}

	return nil
}

// 发布DAG任务
func (a *AsyncQueue) PublishDAG(tasks []TaskInfo) error {
	return nil
}

func (a *AsyncQueue) AddHandler(f Handler) {
	if f != nil {
		a.Handlers = append(a.Handlers, f)
	}
}

func (a *AsyncQueue) consumeApi(consumerID int64) {
	tick := time.NewTicker(2 * time.Second)
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
		// 从高优先级队列获取任务
		resp, err = a.Active(priority)
		if err != nil {
			log.Error("consumeTask GetMessages error", zap.Error(err), zap.Int64("consumerID", consumerID))
			return
		}
		if resp != nil {
			break
		}
	}

	if resp == nil {
		//log.Debug("consumeTask end empty", zapInt64()g("consumerID", consumerID))
		return
	}

	//if resp.Type == TaskTypeDAG && resp.PreTaskID != "" {
	//	status, err := a.GetTaskStatus(resp.PreTaskID)
	//	if err != nil {
	//		log.Error("consumeTask GetTaskStatus error", zap.Error(err), zap.Int64("consumerID", consumerID))
	//		return
	//	}
	//	if status != "success" {
	//		a.Requeue(resp)
	//		log.Debug("Requeue", zap.Int64("consumerID", consumerID), zap.Any("resp", resp))
	//		return
	//	}
	//}

	task := &TaskInfo{
		TaskMeta: TaskMeta{},
		Payload:  resp.Payload,
	}

	for _, handler := range a.Handlers {
		handler.ProcessTask(task)
	}

	a.Done(resp)
	log.Debug("consumeTask end finish", zap.Int64("consumerID", consumerID))
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
