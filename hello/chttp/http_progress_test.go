package chttp

import (
	"net/http"
	"testing"
)

func TestHttpDoProgress(t *testing.T) {
	// 参考：http://speedtest.tele2.net/
	// 1MB   10MB   100MB   1GB   10GB   50GB   100GB   1000GB
	url := "http://speedtest.tele2.net/10MB.zip"
	statusCode, resBody, err := HttpDoProgress(url, http.MethodGet, nil, nil, nil)
	if err != nil {
		t.Errorf("HttpDo() error = %v", err)
		return
	}
	t.Logf("statusCode = %d", statusCode)
	t.Logf("resBody count = %d", len(resBody))
}
