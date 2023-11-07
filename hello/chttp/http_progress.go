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

	// 打印进度
	counter := &WriteCounter{Total: resp.ContentLength}
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	go func() {
		for range ticker.C {
			log.Info("HttpDoProgress downloading", zap.String("uniqueID", uniqueID), zap.Int64("written", counter.Written),
				zap.Int64("total", counter.Total), zap.Int64("costTime", time.Since(start).Milliseconds()))
		}
	}()

	var bs bytes.Buffer
	if _, err = io.Copy(&bs, io.TeeReader(resp.Body, counter)); err != nil {
		log.Error("HttpDoProgress io.Copy err", zap.String("uniqueID", uniqueID), zap.Error(err))
		return 0, nil, err
	}
	log.Info("HttpDoProgress finish", zap.String("uniqueID", uniqueID), zap.Int64("written", counter.Written),
		zap.Int64("total", counter.Total), zap.Int64("costTime", time.Since(start).Milliseconds()))

	return int32(resp.StatusCode), bs.Bytes(), nil
}

type WriteCounter struct {
	Total   int64 // 总大小
	Written int64 // 已写入大小
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Written += int64(n)
	return n, nil
}
