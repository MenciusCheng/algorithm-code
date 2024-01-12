package merge

import (
	"errors"
	"fmt"
	"github.com/MenciusCheng/algorithm-code/utils/log"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"sync"
	"time"
)

var limiter = rate.NewLimiter(200, 200)

// 批量审核，限流200qps，耗时100毫秒
func BatchCheck(req []string) ([]string, error) {
	if !limiter.Allow() {
		return nil, errors.New("limited")
	}

	res := make([]string, 0, len(req))
	for _, s := range req {
		res = append(res, fmt.Sprintf("result:%s", s))
	}
	log.Info("BatchCheck finish", zap.Any("req", req), zap.Any("res", res))
	time.Sleep(100 * time.Millisecond)
	return res, nil
}

// 批量聚合单文本审核
func ContentCheck(sg *Group, text string, batch int32) (string, error) {
	v, err, i := sg.Do("A", func(req []interface{}) (interface{}, error) {
		params := make([]string, 0)
		for _, item := range req {
			params = append(params, item.(string))
		}
		check, err := BatchCheck(params)
		if err != nil {
			return nil, err
		}
		return check, nil
	}, text, batch)
	if err != nil {
		panic(err)
	}
	res := v.([]string)
	if int(i) >= len(res) {
		log.Error("ContentCheck res out of bounds", zap.Any("res", res), zap.Any("text", text), zap.Any("i", i))
		return res[0], nil
	}
	return res[i], nil
}

// call is an in-flight or completed singleflight.Do call
type call struct {
	wg sync.WaitGroup

	// These fields are written once before the WaitGroup is done
	// and are only read after the WaitGroup is done.
	val interface{}
	err error

	// These fields are read and written with the singleflight
	// mutex held before the WaitGroup is done, and are read but
	// not written after the WaitGroup is done.
	dups int

	length   int32
	capacity int32

	ch chan interface{}

	//params []interface{}
	mu sync.Mutex // protects m

	//chans []chan<- Result
}

// Group represents a class of work and forms a namespace in
// which units of work can be executed with duplicate suppression.
type Group struct {
	mu sync.Mutex         // protects m
	m  map[string][]*call // lazily initialized
}

// Result holds the results of Do, so they can be passed
// on a channel.
type Result struct {
	Val    interface{}
	Err    error
	Shared bool
}

func (g *Group) Do(key string, fn func([]interface{}) (interface{}, error), param interface{}, capacity int32) (v interface{}, err error, dups int32) {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string][]*call)
	}

	for _, c := range g.m[key] {
		if c.length < c.capacity && len(c.ch) < cap(c.ch) {
			index := c.length
			c.length++
			c.ch <- param
			g.mu.Unlock()
			c.wg.Wait()
			return c.val, c.err, index
		}
	}

	c := new(call)
	c.wg.Add(1)
	c.ch = make(chan interface{}, capacity)
	c.ch <- param
	c.capacity = capacity
	c.length = 1
	g.m[key] = append(g.m[key], c)
	g.mu.Unlock()

	g.doCall(c, key, fn)
	return c.val, c.err, 0
}

// doCall handles the single call for a key.
func (g *Group) doCall(c *call, key string, fn func([]interface{}) (interface{}, error)) {
	defer func() {
		// the given function invoked runtime.Goexit
		//if !normalReturn && !recovered {
		//	c.err = errGoexit
		//}
		c.wg.Done()
	}()

	params := make([]interface{}, 0)
	exist := true
	for exist {
		select {
		case v := <-c.ch:
			params = append(params, v)
		case <-time.After(1 * time.Millisecond):
			exist = false
			//default:
			//	exist = false
		}
	}

	func() {
		g.mu.Lock()
		defer g.mu.Unlock()
		exist = true
		for exist {
			select {
			case v := <-c.ch:
				//log.Info("doCall Lock and get param", zap.Any("v", v))
				params = append(params, v)
			default:
				exist = false
			}
		}

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

	func() {
		//defer func() {
		//	if !normalReturn {
		//		// Ideally, we would wait to take a stack trace until we've determined
		//		// whether this is a panic or a runtime.Goexit.
		//		//
		//		// Unfortunately, the only way we can distinguish the two is to see
		//		// whether the recover stopped the goroutine from terminating, and by
		//		// the time we know that, the part of the stack trace relevant to the
		//		// panic has been discarded.
		//		if r := recover(); r != nil {
		//			c.err = newPanicError(r)
		//		}
		//	}
		//}()
		c.val, c.err = fn(params)
		//normalReturn = true
	}()

	//if !normalReturn {
	//	recovered = true
	//}
}
