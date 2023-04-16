package mockhttp

import (
	"fmt"
	"github.com/MenciusCheng/algorithm-code/hello/chttp"
	"github.com/jarcoal/httpmock"
	"io/ioutil"
	"net/http"
)

const (
	BaiduUrl = "https://www.baidu.com/"
)

func Basic() {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Exact URL match
	httpmock.RegisterResponder(
		"GET",
		BaiduUrl,
		httpmock.NewStringResponder(200, `[{"id": 1, "name": "My Great Article"}]`),
	)

	resp, _ := http.Get(BaiduUrl)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("baidu body:%s\n", body)
}

func BasicClient() {
	httpmock.ActivateNonDefault(chttp.HttpClient)
	defer httpmock.DeactivateAndReset()

	// Exact URL match
	httpmock.RegisterResponder(
		"GET",
		BaiduUrl,
		httpmock.NewStringResponder(200, `[{"id": 1, "name": "My Great Article"}]`),
	)

	headArr := map[string]string{
		"Content-Type": "application/json",
	}
	req := map[string]interface{}{}

	resStatus, resBody, err := chttp.HttpDo(BaiduUrl, http.MethodGet, headArr, req, nil)
	if err != nil {
		return
	}
	fmt.Printf("baidu resStatus:%d\nresBody:%s\nerr:%v\n", resStatus, resBody, err)
}
