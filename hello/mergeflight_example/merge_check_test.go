package mergeflight_example

import (
	"fmt"
	"github.com/MenciusCheng/algorithm-code/hello/mergeflight"
	"github.com/MenciusCheng/algorithm-code/utils/log"
	"go.uber.org/zap"
	"sync"
	"testing"
	"time"
)

func TestBatchCheck(t *testing.T) {
	start := time.Now()
	n := 200
	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			req := fmt.Sprintf("%d", i)
			want := fmt.Sprintf("result:%d", i)
			res, err := BatchCheck([]string{req})
			if err != nil {
				t.Errorf("BatchCheck err = %v", err)
			} else if res[0] != want {
				t.Errorf("BatchCheck res wrong, got = %v, want = %v", res, want)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	log.Info("TestBatchCheck finish", zap.Any("cost", time.Since(start)))
}

func TestContentCheck(t *testing.T) {
	start := time.Now()
	sg := mergeflight.NewGroup()

	n := 2000
	batch := 10
	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			req := fmt.Sprintf("%d", i)
			want := fmt.Sprintf("result:%d", i)
			res, err := ContentCheck(sg, req, batch)
			if err != nil {
				t.Errorf("ContentCheck err = %v", err)
			} else if res != want {
				t.Errorf("ContentCheck res wrong, got = %v, want = %v", res, want)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	log.Info("TestContentCheck finish", zap.Any("cost", time.Since(start)))
}

func TestContentCheck_Multi(t *testing.T) {
	start := time.Now()
	sg := mergeflight.NewGroup()

	n := 200
	loop := 10
	batch := 10
	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			for j := 0; j < loop; j++ {
				req := fmt.Sprintf("%d", i*1000+j)
				want := fmt.Sprintf("result:%d", i*1000+j)
				res, err := ContentCheck(sg, req, batch)
				if err != nil {
					log.Error("ContentCheck err", zap.Error(err))
				} else if res != want {
					t.Errorf("ContentCheck res wrong, got = %v, want = %v", res, want)
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	log.Info("TestContentCheck finish", zap.Any("cost", time.Since(start)))
}
