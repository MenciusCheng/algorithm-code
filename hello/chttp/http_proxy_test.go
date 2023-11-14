package chttp

import (
	"net/http"
	"testing"
)

func TestHttpDoProxy(t *testing.T) {
	urlPath := "https://www.google.com"
	//urlPath := "https://www.baidu.com"

	statusCode, resBody, err := HttpDoProxy(urlPath, http.MethodGet, nil, nil, nil)
	if err != nil {
		t.Errorf("HttpDo() error = %v", err)
		return
	}

	t.Logf("statusCode = %d", statusCode)
	t.Logf("resBody = %s", string(resBody))
}
