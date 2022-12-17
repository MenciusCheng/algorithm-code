package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/description/

28. 找出字符串中第一个匹配项的下标
中等
给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）。如果 needle 不是 haystack 的一部分，则返回  -1 。

示例 1：

输入：haystack = "sadbutsad", needle = "sad"
输出：0
解释："sad" 在下标 0 和 6 处匹配。
第一个匹配项的下标是 0 ，所以返回 0 。
示例 2：

输入：haystack = "leetcode", needle = "leeto"
输出：-1
解释："leeto" 没有在 "leetcode" 中出现，所以返回 -1 。

提示：

1 <= haystack.length, needle.length <= 10^4
haystack 和 needle 仅由小写英文字符组成
*/
func main() {
	var tests = []struct {
		haystack string
		needle   string
		want     int
	}{
		{
			haystack: "abeababeabf",
			needle:   "abeabf",
			want:     5,
		},
		{
			haystack: "sadbutsad",
			needle:   "sad",
			want:     0,
		},
		{
			haystack: "leetcode",
			needle:   "leeto",
			want:     -1,
		},
	}

	for _, item := range tests {
		if ans := strStr(item.haystack, item.needle); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func strStr(haystack string, needle string) int {
	next := MakeKMP(needle)
	for i := 0; i <= len(haystack)-len(needle); i++ {
		j := 0
		for j < len(needle) && needle[j] == haystack[i+j] {
			j++
		}
		if j == len(needle) {
			return i
		}
		if j > 0 {
			j--
			i += j - next[j]
		}
	}

	return -1
}

func MakeKMP(needle string) []int {
	next := make([]int, len(needle))
	for i, j := 1, 0; i < len(needle); i++ {
		for j > 0 && needle[j] != needle[i] {
			j = next[j-1]
		}
		if needle[j] == needle[i] {
			j++
		}
		next[i] = j
	}
	return next
}
