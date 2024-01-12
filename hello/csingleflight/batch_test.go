package csingleflight

import (
	"fmt"
	"github.com/MenciusCheng/algorithm-code/utils/log"
	"go.uber.org/zap"
	"sync"
	"testing"
	"time"
)

func TestContentCheck(t *testing.T) {
	var (
		wg  sync.WaitGroup
		now = time.Now()
		n   = 1000
	)

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
	fmt.Printf("同时发起 %d 次请求，耗时: %s", n, time.Since(now))
}

func TestExampleCheck2(t *testing.T) {
	ExampleCheck2()
}
