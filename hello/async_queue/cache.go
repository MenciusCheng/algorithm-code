package async_queue

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

type Cache struct {
	QueueName        string        // 队列名称
	RedisConn        *redis.Client // Redis 连接
	Retention        int64         // 任务完成后状态保留秒数
	FailureRetention int64         // 任务失败后状态保留秒数
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
// ARGV[1] -> current timestamp seconds
// ARGV[2] -> task Payload
// ARGV[3] -> task ID
// ARGV[4] -> task Type
// ARGV[5] -> task Priority
// ARGV[6] -> task PreTaskIDs
// ARGV[7] -> task InDegree
// ARGV[8] -> task PostTaskIDs
// ARGV[9] -> task ProcessAt in Unix time
//
// Output:
// Returns 1 if successfully enqueued
// Returns 0 if task ID already exists
var scheduleCmd = redis.NewScript(`
if redis.call("EXISTS", KEYS[1]) == 1 then
	return 0
end
redis.call("HSET", KEYS[1],
           "state", "schedule",
           "schedule_since", ARGV[1],
           "payload", ARGV[2],
           "id", ARGV[3],
           "type", ARGV[4],
           "priority", ARGV[5],
           "preTaskIDs", ARGV[6],
           "inDegree", ARGV[7],
           "postTaskIDs", ARGV[8])
redis.call("ZADD", KEYS[2], ARGV[9], ARGV[3])
return 1
`)

// Schedule 添加任务到延时队列
func (c *Cache) Schedule(msg *TaskMessage) error {
	keys := []string{
		c.TaskKey(msg.ID),
		c.ScheduleKey(msg.Priority),
	}
	argv := []interface{}{
		time.Now().Unix(),
		msg.Payload,
		msg.ID,
		msg.Type,
		msg.Priority,
		msg.PreTaskIDs,
		msg.InDegree,
		msg.PostTaskIDs,
		msg.ProcessAt,
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

// Input:
// KEYS[1] -> AsyncQueue:{<qname>}:task:<task_id>
// KEYS[2] -> AsyncQueue:{<qname>}:p:<priority>:daging
// --
// ARGV[1] -> current timestamp seconds
// ARGV[2] -> task Payload
// ARGV[3] -> task ID
// ARGV[4] -> task Type
// ARGV[5] -> task Priority
// ARGV[6] -> task PreTaskIDs
// ARGV[7] -> task InDegree
// ARGV[8] -> task PostTaskIDs
//
// Output:
// Returns 1 if successfully enqueued
// Returns 0 if task ID already exists
var dagingCmd = redis.NewScript(`
if redis.call("EXISTS", KEYS[1]) == 1 then
	return 0
end
redis.call("HSET", KEYS[1],
           "state", "daging",
           "daging_since", ARGV[1],
           "payload", ARGV[2],
           "id", ARGV[3],
           "type", ARGV[4],
           "priority", ARGV[5],
           "preTaskIDs", ARGV[6],
           "inDegree", ARGV[7],
           "postTaskIDs", ARGV[8])
redis.call("ZADD", KEYS[2], ARGV[7], ARGV[3])
return 1
`)

// DAGing 添加任务到DAG队列
func (c *Cache) DAGing(msg *TaskMessage) error {
	keys := []string{
		c.TaskKey(msg.ID),
		c.DAGingKey(msg.Priority),
	}
	argv := []interface{}{
		time.Now().Unix(),
		msg.Payload,
		msg.ID,
		msg.Type,
		msg.Priority,
		msg.PreTaskIDs,
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

// Input:
// KEYS[1] -> AsyncQueue:{<qname>}:task:<task_id>
// KEYS[2] -> AsyncQueue:{<qname>}:p:<priority>:pending
// --
// ARGV[1] -> current timestamp seconds
// ARGV[2] -> task Payload
// ARGV[3] -> task ID
// ARGV[4] -> task Type
// ARGV[5] -> task Priority
//
// Output:
// Returns 1 if successfully enqueued
// Returns 0 if task ID already exists
var pendingCmd = redis.NewScript(`
if redis.call("EXISTS", KEYS[1]) == 1 then
	return 0
end
redis.call("HSET", KEYS[1],
           "state", "pending",
           "pending_since", ARGV[1],
           "payload", ARGV[2],
           "id", ARGV[3],
           "type", ARGV[4],
           "priority", ARGV[5])
redis.call("LPUSH", KEYS[2], ARGV[3])
return 1
`)

// Pending 添加任务到挂起队列
func (c *Cache) Pending(msg *TaskMessage) error {
	keys := []string{
		c.TaskKey(msg.ID),
		c.PendingKey(msg.Priority),
	}
	argv := []interface{}{
		time.Now().Unix(),
		msg.Payload,
		msg.ID,
		msg.Type,
		msg.Priority,
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
// KEYS[2] -> AsyncQueue:{<qname>}:active
// --
// ARGV[1] -> task key prefix
// ARGV[2] -> current timestamp seconds
//
// Output:
// Returns nil if no processable task is found in the given queue.
// Returns Hash
var activeCmd = redis.NewScript(`
local id = redis.call("RPOP", KEYS[1])
if id then
	redis.call("ZADD", KEYS[2], ARGV[2], id)
	local key = ARGV[1] .. id
	redis.call("HSET", key, "state", "active")
	redis.call("HSET", key, "active_since", ARGV[2])
	return redis.call("HGETALL", key)
end
return nil`)

// Active 从挂起队列获取任务，并添加到处理队列
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

	// Redis Hash 转为 Map
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
		Payload:     []byte(resMap["payload"]),
		ID:          resMap["id"],
		Type:        resMap["type"],
		PreTaskIDs:  resMap["preTaskIDs"],
		PostTaskIDs: resMap["postTaskIDs"],
	}
	if len(resMap["priority"]) > 0 {
		msg.Priority, _ = strconv.ParseInt(resMap["priority"], 10, 64)
	}
	if len(resMap["retry"]) > 0 {
		msg.Retry, _ = strconv.ParseInt(resMap["retry"], 10, 64)
	}
	if len(resMap["inDegree"]) > 0 {
		msg.InDegree, _ = strconv.Atoi(resMap["inDegree"])
	}

	return msg, nil
}

// KEYS[1] -> AsyncQueue:{<qname>}:active
// KEYS[2] -> AsyncQueue:{<qname>}:task:<task_id>
// KEYS[3] -> AsyncQueue:{<qname>}:p:<priority>:daging
// -------
// ARGV[1] -> task ID
// ARGV[2] -> retention seconds
// ARGV[3] -> task Type
// ARGV[4] -> task PostTaskIDs
// ARGV[5] -> task key prefix
var successCmd = redis.NewScript(`
if redis.call("ZREM", KEYS[1], ARGV[1]) == 0 then
	return redis.error_reply("NOT FOUND")
end
redis.call("HSET", KEYS[2], "state", "success")
redis.call("EXPIRE", KEYS[2], ARGV[2])
if ARGV[3] == "dag" and ARGV[4] ~= "" then
	for taskID in (ARGV[4] .. ","):gmatch("(.-)" .. ",") do
		if redis.call("ZREM", KEYS[3], taskID) == 1 then
			local key = ARGV[5] .. taskID
			local inDegree = redis.call("HINCRBY", key, "inDegree", -1)
			redis.call("ZADD", KEYS[3], inDegree, taskID)
		end
	end
end
return redis.status_reply("OK")
`)

// Success 任务处理成功，从处理队列中移除任务
// 如果是DAG任务，则对子任务的入度减少1
func (c *Cache) Success(msg *TaskMessage) error {
	keys := []string{
		c.ActiveKey(),
		c.TaskKey(msg.ID),
		c.DAGingKey(msg.Priority),
	}
	argv := []interface{}{
		msg.ID,
		c.Retention,
		msg.Type,
		msg.PostTaskIDs,
		c.TaskKeyPrefix(),
	}
	return successCmd.Run(c.RedisConn, keys, argv...).Err()
}

// KEYS[1] -> AsyncQueue:{<qname>}:active
// KEYS[2] -> AsyncQueue:{<qname>}:task:<task_id>
// KEYS[3] -> AsyncQueue:{<qname>}:p:<priority>:daging
// KEYS[4] -> AsyncQueue:{<qname>}:failure
// -------
// ARGV[1] -> task ID
// ARGV[2] -> failure retention seconds
// ARGV[3] -> task Type
// ARGV[4] -> task PostTaskIDs
// ARGV[5] -> task key prefix
// ARGV[6] -> current timestamp seconds
var failureCmd = redis.NewScript(`
if redis.call("ZREM", KEYS[1], ARGV[1]) == 0 then
	return redis.error_reply("NOT FOUND")
end
redis.call("HSET", KEYS[2], "state", "failure")
redis.call("EXPIRE", KEYS[2], ARGV[2])
if tonumber(ARGV[2]) > 0 then
	redis.call("ZADD", KEYS[4], ARGV[6], ARGV[1])
	redis.call("EXPIRE", KEYS[4], ARGV[2])
end
if ARGV[3] == "dag" and ARGV[4] ~= "" then
	for taskID in (ARGV[4] .. ","):gmatch("(.-)" .. ",") do
		if redis.call("ZREM", KEYS[3], taskID) == 1 then
			local taskKey = ARGV[5] .. taskID
			local inDegree = redis.call("HINCRBY", taskKey, "inDegree", -1)
			redis.call("ZADD", KEYS[3], inDegree, taskID)
		end
	end
end
return redis.status_reply("OK")
`)

// Failure 任务处理失败，从处理队列中移除任务，并写入失败队列，不再重试
// 如果是DAG任务，则对子任务的入度减少1
func (c *Cache) Failure(msg *TaskMessage) error {
	keys := []string{
		c.ActiveKey(),
		c.TaskKey(msg.ID),
		c.DAGingKey(msg.Priority),
		c.FailureKey(),
	}
	argv := []interface{}{
		msg.ID,
		c.FailureRetention,
		msg.Type,
		msg.PostTaskIDs,
		c.TaskKeyPrefix(),
		time.Now().Unix(),
	}
	return failureCmd.Run(c.RedisConn, keys, argv...).Err()
}

// KEYS[1] -> AsyncQueue:{<qname>}:active
// KEYS[2] -> AsyncQueue:{<qname>}:p:<priority>:schedule
// KEYS[3] -> AsyncQueue:{<qname>}:task:<task_id>
// -------
// ARGV[1] -> task ID
// ARGV[2] -> task ProcessAt in Unix time
// ARGV[3] -> task Retry count
var retryCmd = redis.NewScript(`
if redis.call("ZREM", KEYS[1], ARGV[1]) == 0 then
	return redis.error_reply("NOT FOUND")
end
redis.call("ZADD", KEYS[2], ARGV[2], ARGV[1])
redis.call("HSET", KEYS[3], "state", "retry")
redis.call("HSET", KEYS[3], "retry", ARGV[3])
return redis.status_reply("OK")
`)

// Retry 重试任务，从处理队列中移除任务，推送至延时队列
func (c *Cache) Retry(msg *TaskMessage, processAt int64) error {
	keys := []string{
		c.ActiveKey(),
		c.ScheduleKey(msg.Priority),
		c.TaskKey(msg.ID),
	}
	argv := []interface{}{
		msg.ID,
		processAt,
		msg.Retry,
	}
	return retryCmd.Run(c.RedisConn, keys, argv...).Err()
}

// KEYS[1] -> AsyncQueue:{<qname>}:p:<priority>:schedule
// KEYS[2] -> AsyncQueue:{<qname>}:p:<priority>:daging
// KEYS[3] -> AsyncQueue:{<qname>}:p:<priority>:pending
// -------
// ARGV[1] -> current unix time in seconds
// ARGV[2] -> task key prefix
var forwardCmd = redis.NewScript(`
local ids = redis.call("ZRANGEBYSCORE", KEYS[1], "-inf", ARGV[1], "LIMIT", 0, 100)
for _, id in ipairs(ids) do
	local taskKey = ARGV[2] .. id
	local type = redis.call("HGET", taskKey, "type")
	local inDegree = redis.call("HGET", taskKey, "inDegree")
	if type == 'dag' and tonumber(inDegree) > 0 then
	    redis.call("ZADD", KEYS[2], inDegree, id)
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

// Forward 从延时队列中取出到达执行时间的任务，如果是DAG任务并且入度大于0则推入DAG队列，否则推入挂起队列
func (c *Cache) Forward(priority int64) (int64, error) {
	keys := []string{
		c.ScheduleKey(priority),
		c.DAGingKey(priority),
		c.PendingKey(priority),
	}
	argv := []interface{}{
		time.Now().Unix(),
		c.TaskKeyPrefix(),
	}
	n, err := forwardCmd.Run(c.RedisConn, keys, argv...).Int64()
	if err != nil {
		return 0, err
	}
	return n, nil
}

// KEYS[1] -> AsyncQueue:{<qname>}:p:<priority>:daging
// KEYS[2] -> AsyncQueue:{<qname>}:p:<priority>:pending
// -------
// ARGV[1] -> current unix time in seconds
// ARGV[2] -> task key prefix
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

// DAGForward 从DAG队列中取出入度为0的任务，推入挂起队列
func (c *Cache) DAGForward(priority int64) (int64, error) {
	keys := []string{
		c.DAGingKey(priority),
		c.PendingKey(priority),
	}
	argv := []interface{}{
		time.Now().Unix(),
		c.TaskKeyPrefix(),
	}
	n, err := dagforwardCmd.Run(c.RedisConn, keys, argv...).Int64()
	if err != nil {
		return 0, err
	}
	return n, nil
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

func (c *Cache) FailureKey() string {
	return fmt.Sprintf("%s:failure", c.QueueKeyPrefix())
}
