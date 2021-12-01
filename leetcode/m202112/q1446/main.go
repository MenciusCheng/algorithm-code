package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode-cn.com/problems/consecutive-characters/

1446. 连续字符
给你一个字符串 s ，字符串的「能量」定义为：只包含一种字符的最长非空子字符串的长度。

请你返回字符串的能量。

示例 1：

输入：s = "leetcode"
输出：2
解释：子字符串 "ee" 长度为 2 ，只包含字符 'e' 。
示例 2：

输入：s = "abbcccddddeeeeedcba"
输出：5
解释：子字符串 "eeeee" 长度为 5 ，只包含字符 'e' 。
示例 3：

输入：s = "triplepillooooow"
输出：5
示例 4：

输入：s = "hooraaaaaaaaaaay"
输出：11
示例 5：

输入：s = "tourist"
输出：1

提示：

1 <= s.length <= 500
s 只包含小写英文字母。

*/
func main() {
	var tests = []struct {
		s    string
		want int
	}{
		{
			s:    "leetcode",
			want: 2,
		},
		{
			s:    "abbcccddddeeeeedcba",
			want: 5,
		},
		{
			s:    "cc",
			want: 2,
		},
	}

	for _, item := range tests {
		if ans := maxPower(item.s); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func maxPower(s string) int {
	maxP := 1
	subP := 1

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			subP += 1
			if subP > maxP {
				maxP = subP
			}
		} else {
			subP = 1
		}
	}
	return maxP
}
