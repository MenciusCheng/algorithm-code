package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode-cn.com/problems/reverse-only-letters/

917. 仅仅反转字母
给你一个字符串 s ，根据下述规则反转字符串：

所有非英文字母保留在原有位置。
所有英文字母（小写或大写）位置反转。
返回反转后的 s 。

示例 1：

输入：s = "ab-cd"
输出："dc-ba"
示例 2：

输入：s = "a-bC-dEf-ghIj"
输出："j-Ih-gfE-dCba"
示例 3：

输入：s = "Test1ng-Leet=code-Q!"
输出："Qedo1ct-eeLg=ntse-T!"

提示

1 <= s.length <= 100
s 仅由 ASCII 值在范围 [33, 122] 的字符组成
s 不含 '\"' 或 '\\'
*/
func main() {
	var tests = []struct {
		s    string
		want string
	}{
		{
			s:    "ab-cd",
			want: "dc-ba",
		},
		{
			s:    "a-bC-dEf-ghIj",
			want: "j-Ih-gfE-dCba",
		},
		{
			s:    "Test1ng-Leet=code-Q!",
			want: "Qedo1ct-eeLg=ntse-T!",
		},
	}

	for _, item := range tests {
		if ans := reverseOnlyLetters(item.s); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func reverseOnlyLetters(s string) string {
	sb := []byte(s)
	i := 0
	j := len(s) - 1
	for i < j {
		for i < j && !(sb[i] >= 'a' && sb[i] <= 'z' || sb[i] >= 'A' && sb[i] <= 'Z') {
			i++
		}
		for i < j && !(sb[j] >= 'a' && sb[j] <= 'z' || sb[j] >= 'A' && sb[j] <= 'Z') {
			j--
		}
		if i < j {
			sb[i], sb[j] = sb[j], sb[i]
			i++
			j--
		}
	}

	return string(sb)
}
