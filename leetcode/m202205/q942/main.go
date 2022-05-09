package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/di-string-match/

942. 增减字符串匹配
由范围 [0,n] 内所有整数组成的 n + 1 个整数的排列序列可以表示为长度为 n 的字符串 s ，其中:

如果 perm[i] < perm[i + 1] ，那么 s[i] == 'I'
如果 perm[i] > perm[i + 1] ，那么 s[i] == 'D'
给定一个字符串 s ，重构排列 perm 并返回它。如果有多个有效排列perm，则返回其中 任何一个 。

示例 1：

输入：s = "IDID"
输出：[0,4,1,3,2]
示例 2：

输入：s = "III"
输出：[0,1,2,3]
示例 3：

输入：s = "DDI"
输出：[3,2,0,1]

提示：

1 <= s.length <= 10^5
s 只包含字符 "I" 或 "D"
*/
func main() {
	var tests = []struct {
		s    string
		want []int
	}{
		{
			s:    "IDID",
			want: []int{0, 4, 1, 3, 2},
		},
		{
			s:    "III",
			want: []int{0, 1, 2, 3},
		},
		{
			s:    "DDI",
			want: []int{3, 2, 0, 1},
		},
	}

	for _, item := range tests {
		if ans := diStringMatch(item.s); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func diStringMatch(s string) []int {
	res := make([]int, len(s)+1)
	left := 0
	right := len(s)

	for i := 0; i < len(s); i++ {
		if s[i] == 'I' {
			res[i] = left
			left++
		} else {
			res[i] = right
			right--
		}
	}
	res[len(s)] = left
	return res
}
