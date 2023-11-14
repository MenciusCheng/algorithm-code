package chttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var HttpProxyClient = &http.Client{
	Transport: NewTransportProxy(),
}

func NewTransportProxy() *http.Transport {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.MaxIdleConns = 100
	transport.MaxConnsPerHost = 100
	transport.MaxIdleConnsPerHost = 100
	transport.Proxy = func(request *http.Request) (*url.URL, error) {
		return url.Parse("http://192.168.0.170:7890")
	}
	return transport
}

func HttpDoProxy(url string, requestMethod string, headers map[string]string, body map[string]interface{}, params map[string]string) (int32, []byte, error) {
	var bodyJson []byte
	if body != nil {
		var err error
		bodyJson, err = json.Marshal(body)
		if err != nil {
			return 0, nil, err
		}
	}
	req, err := http.NewRequest(
		requestMethod,
		url,
		bytes.NewBuffer(bodyJson))
	if err != nil {
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
	fmt.Printf("url:%s\n", url)
	resp, err := HttpProxyClient.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	return int32(resp.StatusCode), resBody, nil
}
