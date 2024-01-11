package crate

import (
	"github.com/MenciusCheng/algorithm-code/utils/log"
	"golang.org/x/time/rate"
	"sync"
)

func Limiter(n int) {
	limiter := rate.NewLimiter(200, 200)
	wg := sync.WaitGroup{}

	for i := 0; i < n; i++ {
		go func() {
			wg.Add(1)
			allow := limiter.Allow()
			if !allow {
				log.Info("not allow")
			} else {
				log.Info("allowed")
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
