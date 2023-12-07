package async_queue

const (
	// 默认并发执行任务数
	DefaultConcurrency = 1
	// 默认任务处理超时秒数
	DefaultTimeout = 60 * 60
	// 默认任务完成后状态保留秒数
	DefaultRetention = 0
	// 默认任务失败后状态保留秒数
	DefaultFailureRetention = 60 * 60 * 24
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

// GetDefaultPriorities 默认优先级队列配置，从高到低遍历
func GetDefaultPriorities() []int64 {
	return []int64{QueuePriorityHigh, QueuePriorityDefault, QueuePriorityLow}
}

// 任务ID列表分隔符
const TaskIDSep = ","
