package csingleflight

import (
	"fmt"
	"golang.org/x/sync/singleflight"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func Test_getArticle(t *testing.T) {
	time.AfterFunc(1*time.Second, func() {
		atomic.AddInt32(&count, -count)
	})

	var (
		wg  sync.WaitGroup
		now = time.Now()
		n   = 1000
	)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			// res, _ := singleflightGetArticle(sg, 1)
			res, _ := getArticle(1)
			if res != "article: 1" {
				panic("err")
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("同时发起 %d 次请求，耗时: %s", n, time.Since(now))
}

func Test_singleflightGetArticle(t *testing.T) {
	time.AfterFunc(1*time.Second, func() {
		atomic.AddInt32(&count, -count)
	})

	var (
		wg  sync.WaitGroup
		now = time.Now()
		n   = 1000
		sg  = &singleflight.Group{}
	)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			res, _ := singleflightGetArticle(sg, 1)
			if res != "article: 1" {
				panic("err")
			}
			wg.Done()
		}()
		//time.Sleep(time.Duration(1) * time.Millisecond)
	}

	wg.Wait()
	fmt.Printf("同时发起 %d 次请求，耗时: %s", n, time.Since(now))
}
