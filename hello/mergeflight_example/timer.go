package mergeflight_example

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

// 参考：https://www.cnblogs.com/jiujuan/p/17369964.html
// 在循环体中调用 time.After，导致内存暴涨
func ChannelParamsForTrap() {
	go func() {
		// http 监听8080, 开启 pprof
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Println("listen failed")
		}
	}()

	ch := make(chan interface{}, 100)

	n := 10000
	params := make([]interface{}, 0)

	go func() {
		for i := 0; i < n; i++ {
			ch <- i
			time.Sleep(10 * time.Millisecond)
		}
	}()

	exist := true
	for exist {
		select {
		case v := <-ch:
			params = append(params, v)
			fmt.Println(v)
		case <-time.After(50 * time.Millisecond):
			exist = false
		}
	}

	fmt.Println(params)
	//time.Sleep(1 * time.Minute)
}

// 解决内存暴涨问题
func ChannelParamsForFix(n int) []interface{} {
	ch := make(chan interface{}, 10)
	params := make([]interface{}, 0)

	go func() {
		for i := 0; i < n; i++ {
			ch <- i
			time.Sleep(5 * time.Millisecond)
		}
	}()

	duration := 20 * time.Millisecond
	t := time.NewTimer(duration)
	defer t.Stop()
	exist := true
	for exist {
		if !t.Stop() { // 防止 t 已经过期再 reset
			<-t.C
		}
		t.Reset(duration)

		select {
		case v := <-ch:
			params = append(params, v)
			fmt.Println("ch param", v)
			time.Sleep(10 * time.Millisecond)
		case <-t.C:
			exist = false
			fmt.Println("timeout")
		}
		//time.Sleep(20 * time.Millisecond)
	}

	return params
}
