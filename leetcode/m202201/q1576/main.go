package main

import (
	"fmt"
	"reflect"
)

func main() {
	var tests = []struct {
		s    string
		want string
	}{
		{
			s:    "?zs",
			want: "azs",
		},
		{
			s:    "ubv?w",
			want: "ubvaw",
		},
		{
			s:    "j?qg??b",
			want: "jaqgacb",
		},
		{
			s:    "??yw?ipkj?",
			want: "abywaipkja",
		},
	}

	for _, item := range tests {
		if ans := modifyString(item.s); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func modifyString(s string) string {
	ans := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '?' {
			eA, eB := false, false
			if i > 0 {
				if ans[i-1] == 'a' {
					eA = true
				} else if ans[i-1] == 'b' {
					eB = true
				}
			}
			if i < len(s)-1 {
				if s[i+1] == 'a' {
					eA = true
				} else if s[i+1] == 'b' {
					eB = true
				}
			}
			if eA && eB {
				ans[i] = 'c'
			} else if eA {
				ans[i] = 'b'
			} else {
				ans[i] = 'a'
			}
		} else {
			ans[i] = s[i]
		}
	}
	return string(ans)
}
