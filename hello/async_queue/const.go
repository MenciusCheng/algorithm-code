package async_queue

const (
	// 默认并发执行任务数
	DefaultConcurrency = 1
	// 默认任务处理超时秒数
	DefaultTimeout = 60 * 60
)

// 队列优先级
const (
	QueuePriorityLow     = 0
	QueuePriorityDefault = 3
	QueuePriorityHigh    = 6
)

// 任务类型
const (
	TaskTypeNormal = "normal"
	TaskTypeDAG    = "dag"
)

// 默认优先级队列配置
func GetDefaultPriorities() []int64 {
	return []int64{QueuePriorityHigh, QueuePriorityDefault, QueuePriorityLow}
}
