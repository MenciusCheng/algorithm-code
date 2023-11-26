package async_queue

import "github.com/go-redis/redis"

type OptionFunc func(*AsyncQueue)

func NewAsyncQueueEntity(options ...OptionFunc) *AsyncQueue {
	c := &AsyncQueue{}

	for _, f := range options {
		f(c)
	}

	c.initQueueConcurrency()
	return c
}

func ConfigLevelNames(names ...string) OptionFunc {
	return func(c *AsyncQueue) {
		for i := 0; i < len(names) && i < 3; i++ {
			switch i {
			case 0:
				c.OneLevelName = names[0]
			case 1:
				c.TwoLevelName = names[1]
			case 2:
				c.ThreeLevelName = names[2]
			}
		}
	}
}

func ConfigName(name string) OptionFunc {
	return func(c *AsyncQueue) {
		c.QueueName = name
	}
}

func ConfigMaxQueueConcurrencyNum(num int64) OptionFunc {
	return func(c *AsyncQueue) {
		c.MaxQueueConcurrencyNum = num
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
			c.handlers = append(c.handlers, f)
		}
	}
}

func ConfigBatchSize(batchSize int64) OptionFunc {
	return func(c *AsyncQueue) {
		c.BatchSize = batchSize
	}
}
