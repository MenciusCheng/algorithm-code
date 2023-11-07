package chttp

import (
	"net/http"
	"testing"
)

func TestHttpDo_Download(t *testing.T) {
	// 参考：[收集的一些10G测速文件地址](https://zhuanlan.zhihu.com/p/395861895)
	url := "https://lg-hkg.fdcservers.net/10MBtest.zip"
	statusCode, resBody, err := HttpDo(url, http.MethodGet, nil, nil, nil)
	if err != nil {
		t.Errorf("HttpDo() error = %v", err)
		return
	}
	t.Logf("statusCode = %d", statusCode)
	t.Logf("resBody count = %d", len(resBody))

}
