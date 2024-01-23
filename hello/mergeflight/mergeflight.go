package mergeflight

import (
	"sync"
	"time"
)

func NewGroup() *Group {
	g := &Group{}
	return g
}

// Group 相同 key 的请求，在延迟时间内并发时，会进行参数合并
type Group struct {
	mu sync.Mutex         // 保护 m 字段
	m  map[string][]*call // lazily initialized
}

// call 合并的请求调用
type call struct {
	wg sync.WaitGroup

	// 请求参数 channel
	param chan interface{}
	// 请求参数可合并上限
	capacity int
	// 请求参数已合并个数
	length int
	// 保护请求参数字段
	mu sync.Mutex
	// 参数合并的延迟时间
	delay time.Duration

	// 返回的结果
	val interface{}
	err error
}

// Do 同步执行合并请求
// 传入合并维度key，合并请求方法，单次请求参数，请求合并上限
// 返回合并请求结果，错误值，合并的序号
func (g *Group) Do(key string, fn func([]interface{}) (interface{}, error), param interface{}, capacity int, delay time.Duration) (v interface{}, err error, index int) {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string][]*call)
	}
	if capacity <= 0 {
		// 容量必须大于等1，包含当前参数
		capacity = 1
	}

	for _, c := range g.m[key] {
		if c.length < c.capacity && len(c.param) < cap(c.param) {
			// 请求参数未满时，合并参数
			index = c.length
			c.length++
			c.param <- param
			g.mu.Unlock()
			c.wg.Wait()
			return c.val, c.err, index
		}
	}

	c := new(call)
	c.wg.Add(1)
	c.param = make(chan interface{}, capacity)
	c.param <- param
	c.capacity = capacity
	c.length = 1
	c.delay = delay
	g.m[key] = append(g.m[key], c)
	g.mu.Unlock()

	g.doCall(c, key, fn)
	return c.val, c.err, 0
}

// doCall 处理合并请求
func (g *Group) doCall(c *call, key string, fn func([]interface{}) (interface{}, error)) {
	defer func() {
		c.wg.Done()
	}()

	params := make([]interface{}, 0, c.capacity)
	func() {
		// 延时等待参数合并
		exist := true
		t := time.NewTimer(c.delay)
		defer t.Stop()
		for exist {
			select {
			case v := <-c.param:
				params = append(params, v)
			case <-t.C:
				exist = false
			}
		}
	}()

	func() {
		g.mu.Lock()
		defer g.mu.Unlock()

		// 获取锁后合并剩余参数
		exist := true
		for exist {
			select {
			case v := <-c.param:
				params = append(params, v)
			default:
				exist = false
			}
		}

		// 移除 call
		for i := 0; i < len(g.m[key]); i++ {
			if g.m[key][i] == c {
				g.m[key] = append(g.m[key][:i], g.m[key][i+1:]...)
				break
			}
		}
		if len(g.m[key]) == 0 {
			delete(g.m, key)
		}
	}()

	c.val, c.err = fn(params)
}
