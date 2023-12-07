package async_queue

type OptionFunc func(*AsyncQueue)

// ConfigPriorities 优先级配置
func ConfigPriorities(priorities []int64) OptionFunc {
	return func(c *AsyncQueue) {
		c.Priorities = priorities
	}
}

// ConfigHandler 配置任务处理方法
func ConfigHandler(f Handler) OptionFunc {
	return func(c *AsyncQueue) {
		if f != nil {
			c.Handlers = append(c.Handlers, f)
		}
	}
}

// ConfigConcurrency 配置并发执行任务数
func ConfigConcurrency(num int64) OptionFunc {
	return func(c *AsyncQueue) {
		c.Concurrency = num
	}
}

// ConfigRetryMax 配置最大重试次数
func ConfigRetryMax(num int64) OptionFunc {
	return func(c *AsyncQueue) {
		c.RetryMax = num
	}
}

// ConfigTimeout 配置任务处理超时秒数
func ConfigTimeout(second int64) OptionFunc {
	return func(c *AsyncQueue) {
		c.Timeout = second
	}
}

// ConfigRetention 配置任务完成后状态保留秒数
func ConfigRetention(second int64) OptionFunc {
	return func(c *AsyncQueue) {
		c.Retention = second
	}
}

// ConfigFailureRetention 配置任务失败后状态保留秒数
func ConfigFailureRetention(second int64) OptionFunc {
	return func(c *AsyncQueue) {
		c.FailureRetention = second
	}
}

// ConfigDeadline 配置任务最大期限秒数
func ConfigDeadline(second int64) OptionFunc {
	return func(c *AsyncQueue) {
		c.Deadline = second
	}
}
