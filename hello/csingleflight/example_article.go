package csingleflight

import (
	"fmt"
	"github.com/MenciusCheng/algorithm-code/utils/log"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"sync/atomic"
	"time"
)

// 参考：https://lailin.xyz/post/go-training-week5-singleflight.html

var count int32

func getArticle(id int) (article string, err error) {
	// 假设这里会对数据库进行调用, 模拟不同并发下耗时不同
	n := atomic.AddInt32(&count, 1)
	time.Sleep(time.Duration(count) * time.Millisecond)

	log.Info("getArticle success", zap.Int("id", id), zap.Int32("count", n))
	return fmt.Sprintf("article: %d", id), nil
}

func singleflightGetArticle(sg *singleflight.Group, id int) (string, error) {
	v, err, _ := sg.Do(fmt.Sprintf("%d", id), func() (interface{}, error) {
		return getArticle(id)
	})

	return v.(string), err
}
