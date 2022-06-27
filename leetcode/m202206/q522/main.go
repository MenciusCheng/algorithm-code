package main

import (
	"fmt"
	"reflect"
	"sort"
)

/*
https://leetcode.cn/problems/longest-uncommon-subsequence-ii/

522. 最长特殊序列 II
给定字符串列表 strs ，返回其中 最长的特殊序列 。如果最长特殊序列不存在，返回 -1 。
特殊序列 定义如下：该序列为某字符串 独有的子序列（即不能是其他字符串的子序列）。
 s 的 子序列可以通过删去字符串 s 中的某些字符实现。
例如，"abc" 是 "aebdc" 的子序列，因为您可以删除"aebdc"中的下划线字符来得到 "abc" 。"aebdc"的子序列还包括"aebdc"、 "aeb" 和 "" (空字符串)。

示例 1：

输入: strs = ["aba","cdc","eae"]
输出: 3
示例 2:

输入: strs = ["aaa","aaa","aa"]
输出: -1

提示:

2 <= strs.length <= 50
1 <= strs[i].length <= 10
strs[i] 只包含小写英文字母
*/
func main() {
	var tests = []struct {
		strs []string
		want int
	}{
		{
			strs: []string{"aba", "cdc", "eae"},
			want: 3,
		},
		{
			strs: []string{"aaa", "aaa", "aa"},
			want: -1,
		},
		{
			strs: []string{"aabbcc", "aabbcc", "cb"},
			want: 2,
		},
	}

	for _, item := range tests {
		if ans := findLUSlength(item.strs); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func findLUSlength(strs []string) int {
	strMap := make(map[string]int)
	for _, str := range strs {
		strMap[str]++
	}

	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) > len(strs[j])
	})

	for i, str := range strs {
		if strMap[str] == 1 {
			flag := false
			for j := 0; j < i; j++ {
				if isSub(strs[j], str) {
					flag = true
					break
				}
			}
			if !flag {
				return len(str)
			}
		}
	}

	return -1
}

func isSub(str, sub string) bool {
	if len(str) == len(sub) {
		return str == sub
	}
	j := 0
	for i := 0; i < len(sub); i++ {
		for j < len(str) && sub[i] != str[j] {
			j++
		}
		if j >= len(str) {
			return false
		}
		j++
	}
	return true
}
