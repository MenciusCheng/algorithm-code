package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode-cn.com/problems/repeated-string-match/

686. 重复叠加字符串匹配
给定两个字符串 a 和 b，寻找重复叠加字符串 a 的最小次数，使得字符串 b 成为叠加后的字符串 a 的子串，如果不存在则返回 -1。
注意：字符串 "abc" 重复叠加 0 次是 ""，重复叠加 1 次是 "abc"，重复叠加 2 次是 "abcabc"。

示例 1：
输入：a = "abcd", b = "cdabcdab"
输出：3
解释：a 重复叠加三遍后为 "abcdabcdabcd", 此时 b 是其子串。

示例 2：
输入：a = "a", b = "aa"
输出：2

示例 3：
输入：a = "a", b = "a"
输出：1

示例 4：
输入：a = "abc", b = "wxyz"
输出：-1

提示：

1 <= a.length <= 10^4
1 <= b.length <= 10^4
a 和 b 由小写英文字母组成
*/
func main() {
	var tests = []struct {
		a    string
		b    string
		want int
	}{
		{
			a:    "abcd",
			b:    "cdabcdab",
			want: 3,
		},
		{
			a:    "a",
			b:    "aa",
			want: 2,
		},
		{
			a:    "a",
			b:    "a",
			want: 1,
		},
		{
			a:    "abc",
			b:    "wxyz",
			want: -1,
		},
	}

	for _, item := range tests {
		if ans := repeatedStringMatch(item.a, item.b); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func repeatedStringMatch(a string, b string) int {
	if b == "" {
		return 0
	}
	for i := 0; i < len(a); i++ {
		isMatch := true
		for j := 0; j < len(b); j++ {
			if a[(i+j)%len(a)] != b[j] {
				isMatch = false
				break
			}
		}
		if isMatch {
			left := len(b) - (len(a) - i)
			times := left / len(a)
			count := times + 1
			if left%len(a) > 0 {
				count++
			}
			return count
		}
	}

	return -1
}
