package csingleflight

import (
	"errors"
	"fmt"
	"github.com/MenciusCheng/algorithm-code/utils/log"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"sync"
	"sync/atomic"
	"time"
)

var limiter = rate.NewLimiter(200, 200)

func BatchCheck(req []string) ([]string, error) {
	if !limiter.Allow() {
		return nil, errors.New("limited")
	}

	res := make([]string, 0, len(req))
	for _, s := range req {
		res = append(res, fmt.Sprintf("result:%s", s))
	}
	log.Info("BatchCheck finish", zap.Any("req", req), zap.Any("res", res))
	return res, nil
}

func ContentCheck(text string) (string, error) {
	resArr, err := BatchCheck([]string{text})
	if err != nil {
		return "", err
	}
	return resArr[0], err
}

func ExampleCheck() {
	n := 1000
	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			res, err := ContentCheck("s")
			if err != nil {
				//panic(err)
				log.Error("ContentCheck err", zap.Error(err))
			} else if res != "result:s" {
				panic("err")
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func ContentCheck2(sg *Group, text string) (string, error) {
	v, err, i, kk := sg.Do("A", func(req []interface{}) (interface{}, error) {
		params := make([]string, 0)
		for _, item := range req {
			params = append(params, item.(string))
		}
		check, err := BatchCheck(params)
		if err != nil {
			return "", err
		}
		return check, nil
	}, text, 10)
	if err != nil {
		panic(err)
	}
	res := v.([]string)
	if int(i) >= len(res) {
		log.Error("ContentCheck2 res out of bounds", zap.Any("res", res), zap.Any("text", text), zap.Any("i", i), zap.String("key", kk))
		return res[0], err
	}
	return res[i], err
}

func ExampleCheck2() {
	sg := &Group{}

	n := 100
	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			res, err := ContentCheck2(sg, fmt.Sprintf("%d", i))
			if err != nil {
				//panic(err)
				log.Error("ContentCheck err", zap.Error(err))
			} else if res != fmt.Sprintf("result:%d", i) {
				//panic("err")
				log.Error("ContentCheck res wrong", zap.Any("want", fmt.Sprintf("result:%d", i)), zap.Any("got", res))
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
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

	len int32
	cap int32

	ch chan interface{}

	//params []interface{}
	mu sync.Mutex // protects m
}

// Group represents a class of work and forms a namespace in
// which units of work can be executed with duplicate suppression.
type Group struct {
	mu sync.Mutex       // protects m
	m  map[string]*call // lazily initialized
}

func (g *Group) Do(key string, fn func([]interface{}) (interface{}, error), param interface{}, cap int32) (v interface{}, err error, dups int32, kk string) {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}

	for i := 0; i < 1000; i++ {
		t := fmt.Sprintf("%s-%d", key, i)
		if c, ok := g.m[t]; ok && c.len < c.cap {
			//index := c.len
			//c.len++
			index := atomic.AddInt32(&c.len, 1) - 1
			c.ch <- param
			g.mu.Unlock()
			log.Info("wait", zap.Any("text", param), zap.Any("index", index), zap.Any("key", t))
			c.wg.Wait()
			return c.val, c.err, index, t
		} else if !ok {
			key = t
			break
		}
	}

	//if ok && len(oc) > 0 && oc[len(oc)-1].dups < cap {
	//	c := oc[len(oc)-1]
	//	c.dups++
	//	c.ch <- param
	//	dups = c.dups
	//	g.mu.Unlock()
	//	c.wg.Wait()
	//
	//	//if e, ok := c.err.(*panicError); ok {
	//	//	panic(e)
	//	//} else if c.err == errGoexit {
	//	//	runtime.Goexit()
	//	//}
	//	return c.val, c.err, dups
	//}
	c := new(call)
	c.wg.Add(1)
	c.ch = make(chan interface{}, cap)
	c.ch <- param
	c.cap = cap
	c.len = 1
	g.m[key] = c
	g.mu.Unlock()

	g.doCall(c, key, fn)
	log.Info("docall", zap.Any("text", param), zap.Any("index", 0), zap.Any("key", key))
	return c.val, c.err, 0, key
}

// doCall handles the single call for a key.
func (g *Group) doCall(c *call, key string, fn func([]interface{}) (interface{}, error)) {
	//normalReturn := false
	//recovered := false

	// use double-defer to distinguish panic from runtime.Goexit,
	// more details see https://golang.org/cl/134395
	defer func() {
		// the given function invoked runtime.Goexit
		//if !normalReturn && !recovered {
		//	c.err = errGoexit
		//}

		g.mu.Lock()
		defer g.mu.Unlock()
		//if g.m[key] == c {
		//	delete(g.m, key)
		//}
		c.wg.Done()

		//if e, ok := c.err.(*panicError); ok {
		//	// In order to prevent the waiting channels from being blocked forever,
		//	// needs to ensure that this panic cannot be recovered.
		//	if len(c.chans) > 0 {
		//		go panic(e)
		//		select {} // Keep this goroutine around so that it will appear in the crash dump.
		//	} else {
		//		panic(e)
		//	}
		//} else if c.err == errGoexit {
		//	// Already in the process of goexit, no need to call again
		//} else {
		//	// Normal return
		//	for _, ch := range c.chans {
		//		ch <- Result{c.val, c.err, c.dups > 0}
		//	}
		//}
	}()

	params := make([]interface{}, 0)
	exist := true
	for exist {
		time.Sleep(1 * time.Millisecond)
		exist = false
		select {
		case v := <-c.ch:
			//log.Info("read ch 1", zap.Any("value", v))
			params = append(params, v)
			exist = true
		default:
			//log.Info("default case")
		}
	}

	//g.mu.Lock()
	//for exist {
	//	//time.Sleep(1 * time.Millisecond)
	//	exist = false
	//	select {
	//	case v := <-c.ch:
	//		log.Info("read ch 2", zap.Any("value", v))
	//		params = append(params, v)
	//		exist = true
	//	default:
	//		//log.Info("default case")
	//	}
	//}
	//g.mu.Unlock()
	log.Info("params", zap.Any("key", key), zap.Any("params", params))

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
