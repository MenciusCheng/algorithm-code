package async_queue

import (
	"errors"
	"github.com/MenciusCheng/algorithm-code/utils"
	"github.com/MenciusCheng/algorithm-code/utils/log"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"sync"
	"time"
)

// 队列参数结构体
type AsyncQueue struct {
	OneLevelName           string        // 一级名称
	TwoLevelName           string        // 二级名称
	ThreeLevelName         string        // 三级名称
	MaxQueueConcurrencyNum int64         // 队列数
	RedisConn              *redis.Client // Redis 连接
	QueueKey               string        // 队列Key，根据一二三级名称计算
	handlers               []Handler     // 消息处理方法
	BatchSize              int64         // 一次取出处理的任务数，默认为10

	QueueName string // 队列名称

	lock     sync.Mutex // protects the fields
	consumed bool       // 是否已开启消费
}

// 初始化维度的队列名，按照配置生成队列数，对应按配置生成消费协程
func (a *AsyncQueue) initQueueConcurrency() {
	if a.QueueKey == "" {
		a.QueueKey = GetQueueKey(a)
	}
	if a.BatchSize == 0 {
		a.BatchSize = DefaultBatchSize
	}
}

// 初始化调用队列数据
func (a *AsyncQueue) InitQueue() error {
	if a.MaxQueueConcurrencyNum <= 0 {
		return errors.New("MaxQueueConcurrencyNum is zero")
	}
	// queueID 范围为 [1, MaxQueueConcurrencyNum]
	for i := int64(1); i <= a.MaxQueueConcurrencyNum; i++ {
		go func(queueID int64) {
			a.consumeApi(queueID)
		}(i)
	}
	return nil
}

// 发布任务
func (a *AsyncQueue) Publish(task Task) error {
	now := time.Now().Unix()

	msg := &TaskMessage{
		ID:      utils.GetUUIDV4(),
		Payload: task.Payload,
	}

	var err error
	if task.ProcessAt > now {
		err = a.Schedule(msg, task.ProcessAt)
	} else {
		err = a.Pending(msg)
	}
	if err != nil {
		return err
	}

	return nil
}

// 发布任务
func (a *AsyncQueue) PublishDAG(tasks []Task) error {
	msgList := make([]*TaskMessage, 0, len(tasks))
	for i := 0; i < len(tasks); i++ {
		msg := &TaskMessage{
			ID:      utils.GetUUIDV4(),
			Payload: tasks[i].Payload,
			Type:    TaskTypeDAG,
		}
		if i > 0 {
			msg.PreTaskID = msgList[i-1].ID
		}
		msgList = append(msgList, msg)
	}

	// TODO 改成批量
	for _, msg := range msgList {
		err := a.Pending(msg)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *AsyncQueue) AddHandler(f Handler) {
	if f != nil {
		a.handlers = append(a.handlers, f)
	}
}

func (a *AsyncQueue) consumeApi(queueID int64) {
	tick := time.NewTicker(2 * time.Second)
	for {
		select {
		case <-tick.C:
			a.consumeTask(queueID)
		}
	}
}

func (a *AsyncQueue) consumeTask(queueID int64) {
	queueKey := GetQueueConcurrencyKey(a.QueueKey, queueID)
	//log.Debug("consumeTask start", zap.String("queueKey", queueKey))

	// 从队列读出数据
	if !a.checkQueueID(queueID) {
		log.Error("consumeTask checkQueueID error", zap.Int64("queueID", queueID), zap.String("queueKey", queueKey))
		return
	}
	resp, err := a.Active()
	if err != nil {
		log.Error("consumeTask GetMessages error", zap.Error(err), zap.String("queueKey", queueKey))
		return
	}
	if resp == nil {
		//log.Debug("consumeTask end empty", zap.String("queueKey", queueKey))
		return
	}

	if resp.Type == TaskTypeDAG && resp.PreTaskID != "" {
		status, err := a.GetTaskStatus(resp.PreTaskID)
		if err != nil {
			log.Error("consumeTask GetTaskStatus error", zap.Error(err), zap.String("queueKey", queueKey))
			return
		}
		if status != "success" {
			a.Requeue(resp)
			log.Debug("Requeue", zap.String("queueKey", queueKey), zap.Any("resp", resp))
			return
		}
	}

	param := HandlerParam{
		Payload: resp.Payload,
	}
	for _, handler := range a.handlers {
		handler(&param)
	}

	a.Done(resp)
	log.Debug("consumeTask end finish", zap.String("queueKey", queueKey))
}

func (a *AsyncQueue) checkQueueID(queueID int64) bool {
	if queueID <= 0 || queueID > a.MaxQueueConcurrencyNum {
		return false
	}
	return true
}

type HandlerParam struct {
	Payload []byte
}

type Handler func(*HandlerParam)
