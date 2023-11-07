package chttp

import (
	"github.com/MenciusCheng/algorithm-code/utils"
	"net/http"
	"testing"
	"time"
)

/*
网速参考：http://speedtest.tele2.net/
1MB   10MB   100MB   1GB   10GB   50GB   100GB   1000GB

MD5
4c6426ac7ef186464ecbb0d81cbfcb1e  100KB.zip
2f282b84e7e608d5852449ed940bfc51  100MB.zip
f1c9645dbc14efddc7d8a322685f26eb  10MB.zip
0f343b0931126a20f133d67c2b018a3b  1KB.zip
b6d81b360a5672d80c27430f39153e2c  1MB.zip
3566de3a97906edb98d004d6b947ae9b  200MB.zip
8f4e33f3dc3e414ff94e5fb6905cba8c  20MB.zip
b2d1236c286a3c0704224fe4105eca49  2MB.zip
d1dd210d6b1312cb342b56d02bd5e651  3MB.zip
d8b61b2c0025919d5321461045c8226f  500MB.zip
25e317773f308e446cc84c503a6d1f85  50MB.zip
59071590099d21dd439896592338bf95  512KB.zip
5f363e0e58a95f06cbe9bbc662c5dfb6  5MB.zip
cd573cfaace07e7949bc0c46028904ff  1GB.zip
2dd26c4d4799ebd29fa31e48d49e8e53  10GB.zip
e7f4706922e1edfdb43cd89eb1af606d  50GB.zip
09cd755eb35bc534487a5796d781a856  100GB.zip
2c9a0f21395470f88f1ded4194979af8  1000GB.zip
*/

func TestHttpDoProgress(t *testing.T) {
	url := "http://speedtest.tele2.net/10MB.zip"
	wantMd5 := "f1c9645dbc14efddc7d8a322685f26eb"
	statusCode, resBody, err := HttpDoProgress(url, http.MethodGet, nil, nil, nil)
	if err != nil {
		t.Errorf("HttpDo() error = %v", err)
		return
	}
	getMd5 := utils.Md5Byte(resBody)
	if getMd5 != wantMd5 {
		t.Errorf("HttpDo() failed, getMd5 = %s, wantMd5 = %s", getMd5, wantMd5)
		return
	}
	t.Logf("HttpDoProgress statusCode = %d, count = %d, getMd5 = %s", statusCode, len(resBody), getMd5)
}

func TestTickerCC(t *testing.T) {
	TickerCC()
	t.Logf("TickerCC finish")
	time.Sleep(5 * time.Second)
}
