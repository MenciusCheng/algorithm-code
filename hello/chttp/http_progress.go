package chttp

import (
	"bytes"
	"encoding/json"
	"github.com/MenciusCheng/algorithm-code/utils"
	"github.com/MenciusCheng/algorithm-code/utils/log"
	"go.uber.org/zap"
	"io"
	"net/http"
	"time"
)

var HttpClientProgress = &http.Client{
	Timeout:   10 * time.Minute,
	Transport: NewTransportProgress(),
}

func NewTransportProgress() *http.Transport {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.MaxIdleConns = 100
	transport.MaxConnsPerHost = 100
	transport.MaxIdleConnsPerHost = 100
	transport.ResponseHeaderTimeout = 30 * time.Second
	return transport
}

// HttpDoProgress 请求带进度打印，用于请求较大文件流接口，较长超时时间
func HttpDoProgress(url string, requestMethod string, headers map[string]string, body map[string]interface{}, params map[string]string) (int32, []byte, error) {
	start := time.Now()
	uniqueID := utils.GetUUIDV4()
	log.Info("HttpDoProgress start", zap.String("uniqueID", uniqueID), zap.Any("headers", headers), zap.Any("body", body), zap.Any("params", params), zap.Any("requestMethod", requestMethod), zap.Any("url", url))
	var bodyJson []byte
	if body != nil {
		var err error
		bodyJson, err = json.Marshal(body)
		if err != nil {
			log.Error("HttpDoProgress Marshal", zap.String("uniqueID", uniqueID), zap.Error(err))
			return 0, nil, err
		}
	}
	req, err := http.NewRequest(
		requestMethod,
		url,
		bytes.NewBuffer(bodyJson))
	if err != nil {
		log.Error("HttpDoProgress NewRequest", zap.String("uniqueID", uniqueID), zap.Error(err))
		return 0, nil, err
	}
	// add params
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	// add headers
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}

	resp, err := HttpClientProgress.Do(req)
	if err != nil {
		log.Error("HttpDoProgress client.Do", zap.String("uniqueID", uniqueID), zap.Error(err))
		return 0, nil, err
	}
	defer resp.Body.Close()
	log.Info("HttpDoProgress resp", zap.String("uniqueID", uniqueID), zap.String("url", url), zap.String("r_url", req.URL.String()), zap.Int("statusCode", resp.StatusCode),
		zap.Int64("total", resp.ContentLength), zap.Int64("costTime", time.Since(start).Milliseconds()))

	// 间隔打印下载进度
	counter := NewWriteCounter(resp.ContentLength, 10*time.Second, func(wc *WriteCounter) {
		log.Info("HttpDoProgress downloading", zap.String("uniqueID", uniqueID), zap.Int64("written", wc.Written),
			zap.Int64("total", wc.Total), zap.Int64("costTime", time.Since(start).Milliseconds()))
	})

	var bs bytes.Buffer
	if _, err = io.Copy(&bs, io.TeeReader(resp.Body, counter)); err != nil {
		log.Error("HttpDoProgress io.Copy err", zap.String("uniqueID", uniqueID), zap.Error(err))
		return 0, nil, err
	}
	log.Info("HttpDoProgress finish", zap.String("uniqueID", uniqueID), zap.Int64("written", counter.Written),
		zap.Int64("total", counter.Total), zap.Int64("costTime", time.Since(start).Milliseconds()))

	return int32(resp.StatusCode), bs.Bytes(), nil
}

func NewWriteCounter(total int64, showInterval time.Duration, showFunc func(wc *WriteCounter)) *WriteCounter {
	return &WriteCounter{
		Total:        total,
		LastShowTime: time.Now(),
		ShowInterval: showInterval,
		ShowFunc:     showFunc,
	}
}

type WriteCounter struct {
	Total        int64                  // 总大小
	Written      int64                  // 已写入大小
	LastShowTime time.Time              // 上次打印时间
	ShowInterval time.Duration          // 打印至少间隔
	ShowFunc     func(wc *WriteCounter) // 打印方法
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Written += int64(n)
	wc.ShowProgress()
	return n, nil
}

func (wc *WriteCounter) ShowProgress() {
	if wc.ShowFunc == nil {
		return
	}

	if time.Since(wc.LastShowTime) >= wc.ShowInterval {
		wc.ShowFunc(wc)
		wc.LastShowTime = time.Now()
	}
}

// Ticker 陷阱，stop的时候不会关闭channel，只是停止发送，会导致协程泄漏
func TickerCC() {
	start := time.Now()
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	stop := make(chan struct{})
	go func(ticker *time.Ticker, ch chan struct{}) {
		for {
			select {
			case <-ticker.C:
				log.Info("ping", zap.Int64("costTime", time.Since(start).Milliseconds()))
			case <-ch:
				log.Info("ping close", zap.Int64("costTime", time.Since(start).Milliseconds()))
				return
			}
		}
	}(ticker, stop)
	time.Sleep(5 * time.Second)
	close(stop)
}
