package async_queue

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

// KEYS[1] -> asynq:{<qname>}:t:<task_id>
// KEYS[2] -> asynq:{<qname>}:scheduled
// -------
// ARGV[1] -> task message data
// ARGV[2] -> process_at time in Unix time
// ARGV[3] -> task ID
//
// Output:
// Returns 1 if successfully enqueued
// Returns 0 if task ID already exists
var scheduleCmd = redis.NewScript(`
if redis.call("EXISTS", KEYS[1]) == 1 then
	return 0
end
redis.call("HSET", KEYS[1],
           "msg", ARGV[1],
           "state", "scheduled")
redis.call("ZADD", KEYS[2], ARGV[2], ARGV[3])
return 1
`)

// Schedule adds the task to the scheduled set to be processed in the future.
func (a *AsyncQueue) Schedule(msg *TaskMessage, processAt int64) error {
	keys := []string{
		a.TaskKey(msg.ID),
		a.ScheduleKey(),
	}
	argv := []interface{}{
		msg.Payload,
		processAt,
		msg.ID,
	}
	n, err := scheduleCmd.Run(a.RedisConn, keys, argv...).Int64()
	if err != nil {
		return err
	}
	if n == 0 {
		return errors.New("msg exists")
	}
	return nil
}

// pendingCmd enqueues a given task message.
//
// Input:
// KEYS[1] -> asynq:{<qname>}:t:<task_id>
// KEYS[2] -> asynq:{<qname>}:pending
// --
// ARGV[1] -> task message data
// ARGV[2] -> task ID
// ARGV[3] -> current unix time in nsec
//
// Output:
// Returns 1 if successfully enqueued
// Returns 0 if task ID already exists
var pendingCmd = redis.NewScript(`
if redis.call("EXISTS", KEYS[1]) == 1 then
	return 0
end
redis.call("HSET", KEYS[1],
           "msg", ARGV[1],
           "id", ARGV[2],
           "state", "pending",
           "pending_since", ARGV[3],
           "type", ARGV[4],
           "preTaskID", ARGV[5])
redis.call("LPUSH", KEYS[2], ARGV[2])
return 1
`)

// Pending adds the given task to the pending list of the queue.
func (a *AsyncQueue) Pending(msg *TaskMessage) error {
	keys := []string{
		a.TaskKey(msg.ID),
		a.PendingKey(),
	}
	argv := []interface{}{
		msg.Payload,
		msg.ID,
		time.Now().Unix(),
		msg.Type,
		msg.PreTaskID,
	}
	n, err := pendingCmd.Run(a.RedisConn, keys, argv...).Int64()
	if err != nil {
		return err
	}
	if n == 0 {
		return errors.New("msg exists")
	}
	return nil
}

// Input:
// KEYS[1] -> asynq:{<qname>}:pending
// KEYS[2] -> asynq:{<qname>}:paused
// KEYS[3] -> asynq:{<qname>}:Active
// KEYS[4] -> asynq:{<qname>}:lease
// --
// ARGV[1] -> initial lease expiration Unix time
// ARGV[2] -> task key prefix
//
// Output:
// Returns nil if no processable task is found in the given queue.
// Returns an encoded TaskMessage.
//
// Note: activeCmd checks whether a queue is paused first, before
// calling RPOPLPUSH to pop a task from the queue.
var activeCmd = redis.NewScript(`
local id = redis.call("RPOPLPUSH", KEYS[1], KEYS[2])
if id then
	local key = ARGV[1] .. id
	redis.call("HSET", key, "state", "Active")
	redis.call("HDEL", key, "pending_since")
	return redis.call("HGETALL", key)
end
return nil`)

// Active queries given queues in order and pops a task message
// off a queue if one exists and returns the message and its lease expiration time.
// Dequeue skips a queue if the queue is paused.
// If all queues are empty, ErrNoProcessableTask error is returned.
func (a *AsyncQueue) Active() (msg *TaskMessage, err error) {
	msg = &TaskMessage{}

	keys := []string{
		a.PendingKey(),
		a.ActiveKey(),
	}
	argv := []interface{}{
		a.TaskKeyPrefix(),
	}
	res, err := activeCmd.Run(a.RedisConn, keys, argv...).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if arr, ok := res.([]interface{}); ok {
		for i := 1; i < len(arr); i += 2 {
			m[arr[i-1].(string)] = arr[i].(string)
		}
	}

	msg.ID = m["id"]
	msg.Payload = []byte(m["msg"])
	msg.Type = m["type"]
	msg.PreTaskID = m["preTaskID"]

	return msg, nil
}

// KEYS[1] -> asynq:{<qname>}:active
// KEYS[2] -> asynq:{<qname>}:lease
// KEYS[3] -> asynq:{<qname>}:t:<task_id>
// KEYS[4] -> asynq:{<qname>}:processed:<yyyy-mm-dd>
// KEYS[5] -> asynq:{<qname>}:processed
// -------
// ARGV[1] -> task ID
// ARGV[2] -> stats expiration timestamp
// ARGV[3] -> max int64 value
var doneCmd = redis.NewScript(`
if redis.call("LREM", KEYS[1], 0, ARGV[1]) == 0 then
  return redis.error_reply("NOT FOUND")
end
redis.call("HSET", KEYS[2], "state", "success")
redis.call("EXPIRE", KEYS[2], ARGV[2])
return redis.status_reply("OK")
`)

/*
if redis.call("DEL", KEYS[2]) == 0 then
  return redis.error_reply("NOT FOUND")
end
*/

// Done removes the task from active queue and deletes the task.
// It removes a uniqueness lock acquired by the task, if any.
func (a *AsyncQueue) Done(msg *TaskMessage) error {
	keys := []string{
		a.ActiveKey(),
		a.TaskKey(msg.ID),
	}
	argv := []interface{}{
		msg.ID,
		60 * 10,
	}
	return doneCmd.Run(a.RedisConn, keys, argv...).Err()
}

// KEYS[1] -> asynq:{<qname>}:active
// KEYS[2] -> asynq:{<qname>}:lease
// KEYS[3] -> asynq:{<qname>}:pending
// KEYS[4] -> asynq:{<qname>}:t:<task_id>
// ARGV[1] -> task ID
// Note: Use RPUSH to push to the head of the queue.
var requeueCmd = redis.NewScript(`
if redis.call("LREM", KEYS[1], 0, ARGV[1]) == 0 then
  return redis.error_reply("NOT FOUND")
end
redis.call("LPUSH", KEYS[2], ARGV[1])
redis.call("HSET", KEYS[3], "state", "pending")
return redis.status_reply("OK")`)

// Requeue moves the task from active queue to the specified queue.
func (a *AsyncQueue) Requeue(msg *TaskMessage) error {
	keys := []string{
		a.ActiveKey(),
		a.PendingKey(),
		a.TaskKey(msg.ID),
	}
	argv := []interface{}{
		msg.ID,
	}
	return requeueCmd.Run(a.RedisConn, keys, argv...).Err()
}

func (a *AsyncQueue) GetTaskStatus(id string) (string, error) {
	state, err := a.RedisConn.HGet(a.TaskKey(id), "state").Result()
	if err != nil {
		return "", err
	}
	return state, nil
}

func (a *AsyncQueue) TaskKeyPrefix() string {
	return fmt.Sprintf("AsyncQueue:{%s}:task:", a.QueueName)
}

func (a *AsyncQueue) TaskKey(id string) string {
	return fmt.Sprintf("AsyncQueue:{%s}:task:%s", a.QueueName, id)
}

func (a *AsyncQueue) ScheduleKey() string {
	return fmt.Sprintf("AsyncQueue:{%s}:schedule", a.QueueName)
}

func (a *AsyncQueue) PendingKey() string {
	return fmt.Sprintf("AsyncQueue:{%s}:pending", a.QueueName)
}

func (a *AsyncQueue) ActiveKey() string {
	return fmt.Sprintf("AsyncQueue:{%s}:active", a.QueueName)
}
