package async_queue

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

type Cache struct {
	QueueName string        // 队列名称
	RedisConn *redis.Client // Redis 连接
}

func NewCache(queueName string, redisConn *redis.Client) *Cache {
	return &Cache{
		QueueName: queueName,
		RedisConn: redisConn,
	}
}

// KEYS[1] -> AsyncQueue:{<qname>}:task:<task_id>
// KEYS[2] -> AsyncQueue:{<qname>}:p:<priority>:schedule
// -------
// ARGV[1] -> task Payload
// ARGV[2] -> task ID
// ARGV[3] -> current timestamp seconds
// ARGV[4] -> task Type
// ARGV[5] -> task PreTaskIDs
// ARGV[6] -> task Priority
// ARGV[7] -> task ProcessAt in Unix time
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
           "id", ARGV[2],
           "state", "schedule",
           "schedule_since", ARGV[3],
           "type", ARGV[4],
           "preTaskIDs", ARGV[5],
           "priority", ARGV[6],
           "postTaskIDs", ARGV[8])
redis.call("ZADD", KEYS[2], ARGV[7], ARGV[2])
return 1
`)

// Schedule adds the task to the schedule set to be processed in the future.
func (c *Cache) Schedule(msg *TaskMessage) error {
	keys := []string{
		c.TaskKey(msg.ID),
		c.ScheduleKey(msg.Priority),
	}
	argv := []interface{}{
		msg.Payload,
		msg.ID,
		time.Now().Unix(),
		msg.Type,
		msg.PreTaskIDs,
		msg.Priority,
		msg.ProcessAt,
		msg.PostTaskIDs,
	}
	n, err := scheduleCmd.Run(c.RedisConn, keys, argv...).Int64()
	if err != nil {
		return err
	}
	if n == 0 {
		return errors.New("msg exists")
	}
	return nil
}

// dagingCmd daging a given task message.
//
// Input:
// KEYS[1] -> AsyncQueue:{<qname>}:task:<task_id>
// KEYS[2] -> AsyncQueue:{<qname>}:p:<priority>:daging
// --
// ARGV[1] -> task Payload
// ARGV[2] -> task ID
// ARGV[3] -> current timestamp seconds
// ARGV[4] -> task Type
// ARGV[5] -> task PreTaskIDs
// ARGV[6] -> task Priority
//
// Output:
// Returns 1 if successfully enqueued
// Returns 0 if task ID already exists
var dagingCmd = redis.NewScript(`
if redis.call("EXISTS", KEYS[1]) == 1 then
	return 0
end
redis.call("HSET", KEYS[1],
           "msg", ARGV[1],
           "id", ARGV[2],
           "state", "daging",
           "daging_since", ARGV[3],
           "type", ARGV[4],
           "preTaskIDs", ARGV[5],
           "priority", ARGV[6],
           "indegree", ARGV[7],
           "postTaskIDs", ARGV[8])
redis.call("ZADD", KEYS[2], ARGV[7], ARGV[2])
return 1
`)

// DAGing adds the given task to the daging list of the queue.
func (c *Cache) DAGing(msg *TaskMessage) error {
	keys := []string{
		c.TaskKey(msg.ID),
		c.DAGingKey(msg.Priority),
	}
	argv := []interface{}{
		msg.Payload,
		msg.ID,
		time.Now().Unix(),
		msg.Type,
		msg.PreTaskIDs,
		msg.Priority,
		msg.InDegree,
		msg.PostTaskIDs,
	}
	n, err := dagingCmd.Run(c.RedisConn, keys, argv...).Int64()
	if err != nil {
		return err
	}
	if n == 0 {
		return errors.New("msg exists")
	}
	return nil
}

// pendingCmd pending a given task message.
//
// Input:
// KEYS[1] -> AsyncQueue:{<qname>}:task:<task_id>
// KEYS[2] -> AsyncQueue:{<qname>}:p:<priority>:pending
// --
// ARGV[1] -> task Payload
// ARGV[2] -> task ID
// ARGV[3] -> current timestamp seconds
// ARGV[4] -> task Type
// ARGV[5] -> task PreTaskIDs
// ARGV[6] -> task Priority
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
           "preTaskIDs", ARGV[5],
           "priority", ARGV[6],
           "postTaskIDs", ARGV[7])
redis.call("LPUSH", KEYS[2], ARGV[2])
return 1
`)

// Pending adds the given task to the pending list of the queue.
func (c *Cache) Pending(msg *TaskMessage) error {
	keys := []string{
		c.TaskKey(msg.ID),
		c.PendingKey(msg.Priority),
	}
	argv := []interface{}{
		msg.Payload,
		msg.ID,
		time.Now().Unix(),
		msg.Type,
		msg.PreTaskIDs,
		msg.Priority,
		msg.PostTaskIDs,
	}
	n, err := pendingCmd.Run(c.RedisConn, keys, argv...).Int64()
	if err != nil {
		return err
	}
	if n == 0 {
		return errors.New("msg exists")
	}
	return nil
}

// Input:
// KEYS[1] -> AsyncQueue:{<qname>}:p:<priority>:pending
// KEYS[2] -> AsyncQueue:{<qname>}:paused
// KEYS[3] -> AsyncQueue:{<qname>}:Active
// KEYS[4] -> AsyncQueue:{<qname>}:lease
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
local id = redis.call("RPOP", KEYS[1])
if id then
	redis.call("ZADD", KEYS[2], ARGV[2], id)
	local key = ARGV[1] .. id
	redis.call("HSET", key, "state", "active")
	redis.call("HDEL", key, "pending_since")
	return redis.call("HGETALL", key)
end
return nil`)

// Active queries given queues in order and pops a task message
// off a queue if one exists and returns the message and its lease expiration time.
// Dequeue skips a queue if the queue is paused.
// If all queues are empty, ErrNoProcessableTask error is returned.
func (c *Cache) Active(priority int64) (msg *TaskMessage, err error) {
	keys := []string{
		c.PendingKey(priority),
		c.ActiveKey(),
	}
	argv := []interface{}{
		c.TaskKeyPrefix(),
		time.Now().Unix(),
	}
	res, err := activeCmd.Run(c.RedisConn, keys, argv...).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	resMap := make(map[string]string)
	resArr, ok := res.([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected res from lua: %v", res)
	}
	for i := 1; i < len(resArr); i += 2 {
		k, ok := resArr[i-1].(string)
		if !ok {
			return nil, fmt.Errorf("unexpected res key from lua: %v", resArr[i-1])
		}
		v, ok := resArr[i].(string)
		if !ok {
			return nil, fmt.Errorf("unexpected res value from lua: %v", resArr[i])
		}
		resMap[k] = v
	}

	msg = &TaskMessage{
		ID:          resMap["id"],
		PreTaskIDs:  resMap["preTaskIDs"],
		PostTaskIDs: resMap["postTaskIDs"],
		Type:        resMap["type"],
		Payload:     []byte(resMap["msg"]),
	}
	if len(resMap["priority"]) > 0 {
		msg.Priority, _ = strconv.ParseInt(resMap["priority"], 10, 64)
	}
	if len(resMap["retry"]) > 0 {
		msg.Retry, _ = strconv.Atoi(resMap["retry"])
	}
	if len(resMap["indegree"]) > 0 {
		msg.InDegree, _ = strconv.ParseInt(resMap["indegree"], 10, 64)
	}

	return msg, nil
}

// KEYS[1] -> AsyncQueue:{<qname>}:active
// KEYS[2] -> AsyncQueue:{<qname>}:lease
// KEYS[3] -> AsyncQueue:{<qname>}:task:<task_id>
// KEYS[4] -> AsyncQueue:{<qname>}:processed:<yyyy-mm-dd>
// KEYS[5] -> AsyncQueue:{<qname>}:processed
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
func (c *Cache) Done(msg *TaskMessage) error {
	keys := []string{
		c.ActiveKey(),
		c.TaskKey(msg.ID),
	}
	argv := []interface{}{
		msg.ID,
		60 * 10,
	}
	return doneCmd.Run(c.RedisConn, keys, argv...).Err()
}

// KEYS[1] -> AsyncQueue:{<qname>}:active
// KEYS[2] -> AsyncQueue:{<qname>}:lease
// KEYS[3] -> AsyncQueue:{<qname>}:task:<task_id>
// KEYS[4] -> AsyncQueue:{<qname>}:processed:<yyyy-mm-dd>
// KEYS[5] -> AsyncQueue:{<qname>}:processed
// -------
// ARGV[1] -> task ID
// ARGV[2] -> stats expiration timestamp
// ARGV[3] -> max int64 value
var successCmd = redis.NewScript(`
if redis.call("ZREM", KEYS[1], ARGV[1]) == 0 then
  return redis.error_reply("NOT FOUND")
end
redis.call("HSET", KEYS[2], "state", "success")
redis.call("EXPIRE", KEYS[2], ARGV[2])
if ARGV[3] == "dag" and ARGV[4] ~= "" then
	for postID in (ARGV[4] .. ","):gmatch("(.-)" .. ",") do
		local key = ARGV[5] .. postID
		local indegree = redis.call("HINCRBY", key, "indegree", -1)
		redis.call("ZADD", KEYS[3], indegree, postID)
	end
end
return redis.status_reply("OK")
`)

// Success removes the task from active queue and deletes the task.
// It removes a uniqueness lock acquired by the task, if any.
func (c *Cache) Success(msg *TaskMessage) error {
	keys := []string{
		c.ActiveKey(),
		c.TaskKey(msg.ID),
		c.DAGingKey(msg.Priority),
	}
	argv := []interface{}{
		msg.ID,
		60 * 10,
		msg.Type,
		msg.PostTaskIDs,
		c.TaskKeyPrefix(),
	}
	// 子任务
	return successCmd.Run(c.RedisConn, keys, argv...).Err()
}

var retryCmd = redis.NewScript(`
if redis.call("ZREM", KEYS[1], ARGV[1]) == 0 then
  return redis.error_reply("NOT FOUND")
end
redis.call("ZADD", KEYS[2], ARGV[2], ARGV[1])
redis.call("HSET", KEYS[3], "state", "retry")
if tonumber(ARGV[3]) == 1 then
	redis.call("HSET", KEYS[3], "retry", ARGV[4])
end
return redis.status_reply("OK")
`)

// Retry removes the task from active queue and deletes the task.
// It removes a uniqueness lock acquired by the task, if any.
func (c *Cache) Retry(msg *TaskMessage, processAt int64, isFailure bool) error {
	if isFailure {
		msg.Retry++
	}

	keys := []string{
		c.ActiveKey(),
		c.ScheduleKey(msg.Priority),
		c.TaskKey(msg.ID),
	}
	argv := []interface{}{
		msg.ID,
		processAt,
		isFailure,
		msg.Retry,
	}
	return retryCmd.Run(c.RedisConn, keys, argv...).Err()
}

// KEYS[1] -> AsyncQueue:{<qname>}:active
// KEYS[2] -> AsyncQueue:{<qname>}:lease
// KEYS[3] -> AsyncQueue:{<qname>}:p:<priority>:pending
// KEYS[4] -> AsyncQueue:{<qname>}:task:<task_id>
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
func (c *Cache) Requeue(msg *TaskMessage) error {
	keys := []string{
		c.ActiveKey(),
		c.PendingKey(msg.Priority),
		c.TaskKey(msg.ID),
	}
	argv := []interface{}{
		msg.ID,
	}
	return requeueCmd.Run(c.RedisConn, keys, argv...).Err()
}

// KEYS[1] -> source queue (e.g. asynq:{<qname>:scheduled or asynq:{<qname>}:retry})
// KEYS[2] -> asynq:{<qname>}:pending
// ARGV[1] -> current unix time in seconds
// ARGV[2] -> task key prefix
// ARGV[3] -> current unix time in nsec
// ARGV[4] -> group key prefix
// Note: Script moves tasks up to 100 at a time to keep the runtime of script short.
var forwardCmd = redis.NewScript(`
local ids = redis.call("ZRANGEBYSCORE", KEYS[1], "-inf", ARGV[1], "LIMIT", 0, 100)
for _, id in ipairs(ids) do
	local taskKey = ARGV[2] .. id
	local type = redis.call("HGET", taskKey, "type")
	if type == 'dag' then
	    redis.call("ZADD", KEYS[2], ARGV[1], id)
	    redis.call("ZREM", KEYS[1], id)
		redis.call("HSET", taskKey, 
				   "state", "daging",
				   "daging_since", ARGV[1])
	else
		redis.call("LPUSH", KEYS[3], id)
	    redis.call("ZREM", KEYS[1], id)
		redis.call("HSET", taskKey,
				   "state", "pending",
				   "pending_since", ARGV[1])
	end
end
return table.getn(ids)`)

// Forward moves tasks with a score less than the current unix time from the delayed (i.e. scheduled | retry) zset
// to the pending list or group set.
// It returns the number of tasks moved.
func (c *Cache) Forward(priority int64) (int, error) {
	keys := []string{
		c.ScheduleKey(priority),
		c.DAGingKey(priority),
		c.PendingKey(priority),
	}
	argv := []interface{}{
		time.Now().Unix(),
		c.TaskKeyPrefix(),
	}
	res, err := forwardCmd.Run(c.RedisConn, keys, argv...).Result()
	if err != nil {
		return 0, err
	}
	// TODO
	if priority == 10 {
		fmt.Printf("Forward res:%+v\n", res)
	}
	return 0, nil
}

// KEYS[1] -> source queue (e.g. asynq:{<qname>:scheduled or asynq:{<qname>}:retry})
// KEYS[2] -> asynq:{<qname>}:pending
// ARGV[1] -> current unix time in seconds
// ARGV[2] -> task key prefix
// ARGV[3] -> current unix time in nsec
// ARGV[4] -> group key prefix
// Note: Script moves tasks up to 100 at a time to keep the runtime of script short.
var dagforwardCmd = redis.NewScript(`
local ids = redis.call("ZRANGEBYSCORE", KEYS[1], "-inf", 0, "LIMIT", 0, 100)
for _, id in ipairs(ids) do
	local taskKey = ARGV[2] .. id
	redis.call("LPUSH", KEYS[2], id)
	redis.call("ZREM", KEYS[1], id)
	redis.call("HSET", taskKey,
			   "state", "pending",
			   "pending_since", ARGV[1])
end
return table.getn(ids)`)

// DAGForward moves tasks with a score less than the current unix time from the delayed (i.e. scheduled | retry) zset
// to the pending list or group set.
// It returns the number of tasks moved.
func (c *Cache) DAGForward(priority int64) (int, error) {
	keys := []string{
		c.DAGingKey(priority),
		c.PendingKey(priority),
	}
	argv := []interface{}{
		time.Now().Unix(),
		c.TaskKeyPrefix(),
	}
	res, err := dagforwardCmd.Run(c.RedisConn, keys, argv...).Result()
	if err != nil {
		return 0, err
	}
	// TODO
	if priority == 10 {
		fmt.Printf("DAGForward res:%+v\n", res)
	}
	return 0, nil
}

func (c *Cache) GetTaskStatus(id string) (string, error) {
	state, err := c.RedisConn.HGet(c.TaskKey(id), "state").Result()
	if err != nil {
		return "", err
	}
	return state, nil
}

func (c *Cache) QueueKeyPrefix() string {
	return fmt.Sprintf("AsyncQueue:{%s}", c.QueueName)
}

func (c *Cache) TaskKey(id string) string {
	return fmt.Sprintf("%s:task:%s", c.QueueKeyPrefix(), id)
}

func (c *Cache) TaskKeyPrefix() string {
	return fmt.Sprintf("%s:task:", c.QueueKeyPrefix())
}

func (c *Cache) ScheduleKey(priority int64) string {
	return fmt.Sprintf("%s:p:%d:schedule", c.QueueKeyPrefix(), priority)
}

func (c *Cache) DAGingKey(priority int64) string {
	return fmt.Sprintf("%s:p:%d:daging", c.QueueKeyPrefix(), priority)
}

func (c *Cache) PendingKey(priority int64) string {
	return fmt.Sprintf("%s:p:%d:pending", c.QueueKeyPrefix(), priority)
}

func (c *Cache) ActiveKey() string {
	return fmt.Sprintf("%s:active", c.QueueKeyPrefix())
}
