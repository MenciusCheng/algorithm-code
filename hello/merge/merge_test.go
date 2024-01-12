package merge

import (
	"fmt"
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
				log.Error("BatchCheck err", zap.Error(err))
			} else if res[0] != want {
				//panic("err")
				log.Error("BatchCheck res wrong", zap.Any("want", want), zap.Any("got", res))
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	log.Info("TestBatchCheck finish", zap.Any("cost", time.Since(start)))
}

func TestContentCheck(t *testing.T) {
	start := time.Now()
	sg := &Group{}

	n := 200
	batch := int32(10)
	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			req := fmt.Sprintf("%d", i)
			want := fmt.Sprintf("result:%d", i)
			res, err := ContentCheck(sg, req, batch)
			if err != nil {
				log.Error("ContentCheck err", zap.Error(err))
			} else if res != fmt.Sprintf("result:%d", i) {
				//panic("err")
				log.Error("ContentCheck res wrong", zap.Any("want", want), zap.Any("got", res))
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	log.Info("TestContentCheck finish", zap.Any("cost", time.Since(start)))
}
