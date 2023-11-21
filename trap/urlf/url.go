package urlf

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
)

func UrlParse(path string) error {
	unescape, _ := url.QueryUnescape(path)

	u, err := url.Parse(unescape)
	if err != nil {
		return err
	}
	fmt.Printf("path: %s, host: %s, upath: %s\n", path, u.Host, u.Path)
	fmt.Printf("%s\n", u.String())
	return nil
}

func IpParse(ipStr string) error {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return errors.New("ip is nil")
	}
	return nil
}

func SplitHostPort(hostport string) error {
	ip, port, err := net.SplitHostPort(hostport)
	if err != nil {
		return err
	}
	fmt.Printf("ip:%s, port:%s\n", ip, port)
	return nil
}

func Forwarded(req *http.Request) string {
	xff := req.Header.Get("X-Forwarded-For")
	arr := strings.SplitN(xff, ",", 2)
	ip := arr[0]
	ip = strings.TrimSpace(ip)
	return ip
}
