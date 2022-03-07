package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode-cn.com/problems/base-7/

504. 七进制数
给定一个整数 num，将其转化为 7 进制，并以字符串形式输出。

示例 1:

输入: num = 100
输出: "202"
示例 2:

输入: num = -7
输出: "-10"

提示：

-10^7 <= num <= 10^7
*/
func main() {
	var tests = []struct {
		num  int
		want string
	}{
		{
			num:  100,
			want: "202",
		},
		{
			num:  -7,
			want: "-10",
		},
	}

	for _, item := range tests {
		if ans := convertToBase7(item.num); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	isNag := false
	if num < 0 {
		isNag = true
		num *= -1
	}

	sb := make([]byte, 0)

	for num > 0 {
		v := num % 7
		num -= v
		num /= 7
		sb = append(sb, byte('0'+v))
	}

	for i := 0; i <= (len(sb)-1)/2; i++ {
		sb[i], sb[len(sb)-1-i] = sb[len(sb)-1-i], sb[i]
	}
	if isNag {
		sb = append([]byte{'-'}, sb...)
	}
	return string(sb)
}
