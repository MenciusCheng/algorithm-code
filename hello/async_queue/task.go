package async_queue

import (
	"github.com/MenciusCheng/algorithm-code/utils"
	"strings"
	"time"
)

// TaskInfo 任务信息
type TaskInfo struct {
	// Payload 业务数据
	Payload []byte

	// TaskMeta 元数据
	TaskMeta
}

type TaskMeta struct {
	// ID 唯一任务ID
	ID string

	// Type 任务类型
	Type string

	// Priority 优先级
	Priority int64

	// ProcessAt 计划开始执行时间戳（秒级）
	ProcessAt int64

	// PreTaskIDs 前置任务ID列表
	PreTaskIDs []string

	// InDegree 入度，前置任务数
	InDegree int

	// PostTaskIDs 后置任务ID列表
	PostTaskIDs []string

	// Retry 已重试次数
	Retry int
}

func NewTaskInfo(payload []byte, taskType string, options ...OptionTaskFunc) *TaskInfo {
	task := &TaskInfo{
		Payload: payload,
		TaskMeta: TaskMeta{
			ID:       utils.GetUUIDV4(),
			Type:     taskType,
			Priority: QueuePriorityDefault,
		},
	}
	for _, f := range options {
		f(task)
	}

	return task
}

// TaskMessage is the internal representation of a task with additional metadata fields.
// Serialized data of this type gets written to redis.
type TaskMessage struct {
	// Payload 业务数据
	Payload []byte

	// ID 唯一任务ID
	ID string

	// Type 任务类型
	Type string

	// Priority 优先级
	Priority int64

	// ProcessAt 计划开始执行时间戳（秒级）
	ProcessAt int64

	// PreTaskIDs 前置任务ID列表, 逗号分割
	PreTaskIDs string

	// InDegree 入度，前置任务数
	InDegree int

	// PostTaskIDs 后置任务ID列表, 逗号分割
	PostTaskIDs string

	// Retry 已重试次数
	Retry int
}

func TaskInfoToMessage(task *TaskInfo) *TaskMessage {
	msg := &TaskMessage{
		Payload:   task.Payload,
		ID:        task.ID,
		Type:      task.Type,
		Priority:  task.Priority,
		ProcessAt: task.ProcessAt,
		InDegree:  task.InDegree,
	}
	if len(task.PreTaskIDs) > 0 {
		msg.PreTaskIDs = strings.Join(task.PreTaskIDs, TaskIDSep)
	}
	if len(task.PostTaskIDs) > 0 {
		msg.PostTaskIDs = strings.Join(task.PostTaskIDs, TaskIDSep)
	}
	return msg
}

func TaskMessageToInfo(msg *TaskMessage) *TaskInfo {
	taskInfo := &TaskInfo{
		Payload: msg.Payload,
		TaskMeta: TaskMeta{
			ID:          msg.ID,
			Type:        msg.Type,
			Priority:    msg.Priority,
			ProcessAt:   msg.ProcessAt,
			PreTaskIDs:  nil,
			InDegree:    msg.InDegree,
			PostTaskIDs: nil,
			Retry:       msg.Retry,
		},
	}
	if len(msg.PreTaskIDs) > 0 {
		taskInfo.PreTaskIDs = strings.Split(msg.PreTaskIDs, TaskIDSep)
	}
	if len(msg.PostTaskIDs) > 0 {
		taskInfo.PostTaskIDs = strings.Split(msg.PostTaskIDs, TaskIDSep)
	}
	return taskInfo
}

type OptionTaskFunc func(*TaskInfo)

// ConfigTaskPriority 配置任务优先级
func ConfigTaskPriority(priority int64) OptionTaskFunc {
	return func(t *TaskInfo) {
		t.Priority = priority
	}
}

// ConfigTaskProcessAt 配置任务计划开始执行时间戳（秒级）
func ConfigTaskProcessAt(processAt int64) OptionTaskFunc {
	return func(t *TaskInfo) {
		t.ProcessAt = processAt
	}
}

// ConfigTaskDelaySecond 配置任务延时执行秒数
func ConfigTaskDelaySecond(second int64) OptionTaskFunc {
	return func(t *TaskInfo) {
		ts := time.Now().Add(time.Duration(second) * time.Second)
		t.ProcessAt = ts.Unix()
	}
}
