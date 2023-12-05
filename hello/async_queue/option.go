package async_queue

import "github.com/go-redis/redis"

type OptionFunc func(*AsyncQueue)

func ConfigName(name string) OptionFunc {
	return func(c *AsyncQueue) {
		c.QueueName = name
	}
}

func ConfigRedisConn(redisConn *redis.Client) OptionFunc {
	return func(c *AsyncQueue) {
		c.RedisConn = redisConn
	}
}

func ConfigHandler(f Handler) OptionFunc {
	return func(c *AsyncQueue) {
		if f != nil {
			c.Handlers = append(c.Handlers, f)
		}
	}
}

func ConfigConcurrency(num int64) OptionFunc {
	return func(c *AsyncQueue) {
		c.Concurrency = num
	}
}

func ConfigRetryMax(num int64) OptionFunc {
	return func(c *AsyncQueue) {
		c.RetryMax = num
	}
}

func ConfigTimeout(second int64) OptionFunc {
	return func(c *AsyncQueue) {
		c.Timeout = second
	}
}
