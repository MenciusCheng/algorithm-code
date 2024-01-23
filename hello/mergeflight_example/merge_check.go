package mergeflight_example

import (
	"errors"
	"fmt"
	"github.com/MenciusCheng/algorithm-code/hello/mergeflight"
	"github.com/MenciusCheng/algorithm-code/utils/log"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
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
func ContentCheck(sg *mergeflight.Group, text string, batch int) (string, error) {
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
	}, text, batch, 10*time.Millisecond)
	if err != nil {
		panic(err)
	}
	res := v.([]string)
	if i >= len(res) {
		log.Error("ContentCheck res out of bounds", zap.Any("res", res), zap.Any("text", text), zap.Any("i", i))
		return res[0], nil
	}
	return res[i], nil
}
