package scheduler

import (
	"fmt"
	"runtime"
	"time"
)

// 这个陷阱已经在 Go 1.14 中基于信号实现了强制抢占而解决。
// https://golang.design/go-questions/sched/sched-trap/
// https://blog.csdn.net/EDDYCJY/article/details/115410510
func InfiniteLoop() {
	var x int
	threads := runtime.GOMAXPROCS(0)
	fmt.Println("threads =", threads)
	for i := 0; i < threads; i++ {
		go func() {
			for {
				x += 1
			}
		}()
	}
	time.Sleep(2 * time.Second)
	fmt.Println("x =", x)
}
