package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/reformat-the-string/

1417. 重新格式化字符串
给你一个混合了数字和字母的字符串 s，其中的字母均为小写英文字母。
请你将该字符串重新格式化，使得任意两个相邻字符的类型都不同。也就是说，字母后面应该跟着数字，而数字后面应该跟着字母。
请你返回 重新格式化后 的字符串；如果无法按要求重新格式化，则返回一个 空字符串 。

示例 1：

输入：s = "a0b1c2"
输出："0a1b2c"
解释："0a1b2c" 中任意两个相邻字符的类型都不同。 "a0b1c2", "0a1b2c", "0c2a1b" 也是满足题目要求的答案。
示例 2：

输入：s = "leetcode"
输出：""
解释："leetcode" 中只有字母，所以无法满足重新格式化的条件。
示例 3：

输入：s = "1229857369"
输出：""
解释："1229857369" 中只有数字，所以无法满足重新格式化的条件。
示例 4：

输入：s = "covid2019"
输出："c2o0v1i9d"
示例 5：

输入：s = "ab123"
输出："1a2b3"

提示：

1 <= s.length <= 500
s 仅由小写英文字母和/或数字组成。
*/
func main() {
	var tests = []struct {
		s    string
		want string
	}{
		{
			s:    "a0b1c2",
			want: "0a1b2c",
		},
		{
			s:    "ab123",
			want: "1a2b3",
		},
	}

	for _, item := range tests {
		if ans := reformat(item.s); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func reformat(s string) string {
	letters := make([]byte, 0)
	nums := make([]byte, 0)
	for i := range s {
		if s[i] >= '0' && s[i] <= '9' {
			nums = append(nums, s[i])
		} else {
			letters = append(letters, s[i])
		}
	}
	diff := len(letters) - len(nums)
	if diff > 1 || diff < -1 {
		return ""
	}

	if diff > 0 {
		return string(combine(letters, nums))
	}
	return string(combine(nums, letters))
}

func combine(a, b []byte) []byte {
	res := make([]byte, 0)
	for i := 0; i < len(b); i++ {
		res = append(res, a[i])
		res = append(res, b[i])
	}
	if len(a) > len(b) {
		res = append(res, a[len(a)-1])
	}
	return res
}
