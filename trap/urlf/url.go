package urlf

import (
	"fmt"
	"net/url"
)

func UrlParse(path string) error {
	u, err := url.Parse(path)
	if err != nil {
		return err
	}
	fmt.Printf("path: %s, host: %s\n", path, u.Host)
	return nil
}
