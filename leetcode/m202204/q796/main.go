package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode-cn.com/problems/rotate-string/

796. 旋转字符串
给定两个字符串, s 和 goal。如果在若干次旋转操作之后，s 能变成 goal ，那么返回 true 。
s 的 旋转操作 就是将 s 最左边的字符移动到最右边。
例如, 若 s = 'abcde'，在旋转一次之后结果就是'bcdea' 。

示例 1:

输入: s = "abcde", goal = "cdeab"
输出: true
示例 2:

输入: s = "abcde", goal = "abced"
输出: false

提示:

1 <= s.length, goal.length <= 100
s 和 goal 由小写英文字母组成
*/
func main() {
	var tests = []struct {
		s    string
		goal string
		want bool
	}{
		{
			s:    "abcde",
			goal: "cdeab",
			want: true,
		},
		{
			s:    "abcde",
			goal: "abced",
			want: false,
		},
	}

	for _, item := range tests {
		if ans := rotateString(item.s, item.goal); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	for i := 0; i < len(goal); i++ {
		if s[0] == goal[i] {
			isEq := true
			for j := 0; j < len(s); j++ {
				if s[j] != goal[(j+i)%len(s)] {
					isEq = false
					break
				}
			}
			if isEq {
				return true
			}

		}
	}

	return false
}
