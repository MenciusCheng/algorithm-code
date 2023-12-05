package async_queue

import (
	"github.com/MenciusCheng/algorithm-code/utils"
	"strings"
)

// TaskInfo 任务信息
type TaskInfo struct {
	// 元数据
	TaskMeta

	// Payload 业务数据
	Payload []byte
}

type TaskMeta struct {
	// 唯一任务ID
	ID string

	// 任务类型
	Type string

	// 优先级
	Priority int64

	// 计划开始执行时间戳（秒级）
	ProcessAt int64

	// 前置任务ID列表
	PreTaskIDs []string

	// Retry 已重试次数
	Retry int
}

// TaskMessage is the internal representation of a task with additional metadata fields.
// Serialized data of this type gets written to redis.
type TaskMessage struct {
	// ID is a unique identifier for each task.
	ID string

	PreTaskID string

	ProcessAt int64

	// Type indicates the kind of the task to be performed.
	Type string

	// 优先级
	Priority int64

	// Payload holds data needed to process the task.
	Payload []byte

	// Queue is a name this message should be enqueued to.
	Queue string

	// Retry is the max number of retry for this task.
	Retry int

	// Retried is the number of times we've retried this task so far.
	Retried int

	// ErrorMsg holds the error message from the last failure.
	ErrorMsg string

	// Time of last failure in Unix time,
	// the number of seconds elapsed since January 1, 1970 UTC.
	//
	// Use zero to indicate no last failure
	LastFailedAt int64

	// Timeout specifies timeout in seconds.
	// If task processing doesn't complete within the timeout, the task will be retried
	// if retry count is remaining. Otherwise it will be moved to the archive.
	//
	// Use zero to indicate no timeout.
	Timeout int64

	// Deadline specifies the deadline for the task in Unix time,
	// the number of seconds elapsed since January 1, 1970 UTC.
	// If task processing doesn't complete before the deadline, the task will be retried
	// if retry count is remaining. Otherwise it will be moved to the archive.
	//
	// Use zero to indicate no deadline.
	Deadline int64

	// UniqueKey holds the redis key used for uniqueness lock for this task.
	//
	// Empty string indicates that no uniqueness lock was used.
	UniqueKey string

	// GroupKey holds the group key used for task aggregation.
	//
	// Empty string indicates no aggregation is used for this task.
	GroupKey string

	// Retention specifies the number of seconds the task should be retained after completion.
	Retention int64

	// CompletedAt is the time the task was processed successfully in Unix time,
	// the number of seconds elapsed since January 1, 1970 UTC.
	//
	// Use zero to indicate no value.
	CompletedAt int64
}

func initTaskInfo(task *TaskInfo, taskType string) {
	task.Type = taskType
	if task.ID == "" {
		task.ID = utils.GetUUIDV4()
	}
	// TODO 校验优先级是否在配置内
	if task.Priority <= 0 {
		task.Priority = QueuePriorityDefault
	}
}

func TaskInfoToMessage(task *TaskInfo) *TaskMessage {
	msg := &TaskMessage{
		ID:       task.ID,
		Type:     task.Type,
		Priority: task.Priority,
		Payload:  task.Payload,
	}
	if len(task.PreTaskIDs) > 0 {
		msg.PreTaskID = strings.Join(task.PreTaskIDs, ",")
	}
	return msg
}
