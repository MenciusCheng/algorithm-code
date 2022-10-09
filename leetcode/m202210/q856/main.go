package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/score-of-parentheses/

856. 括号的分数
给定一个平衡括号字符串 S，按下述规则计算该字符串的分数：

() 得 1 分。
AB 得 A + B 分，其中 A 和 B 是平衡括号字符串。
(A) 得 2 * A 分，其中 A 是平衡括号字符串。

示例 1：

输入： "()"
输出： 1
示例 2：

输入： "(())"
输出： 2
示例 3：

输入： "()()"
输出： 2
示例 4：

输入： "(()(()))"
输出： 6

提示：

S 是平衡括号字符串，且只含有 ( 和 ) 。
2 <= S.length <= 50
*/
func main() {
	var tests = []struct {
		s    string
		want int
	}{
		{
			s:    "()",
			want: 1,
		},
		{
			s:    "(())",
			want: 2,
		},
		{
			s:    "(()(()))",
			want: 6,
		},
	}

	for _, item := range tests {
		if ans := scoreOfParentheses(item.s); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func scoreOfParentheses(s string) int {
	start := 0
	stack := 0
	sum := 0
	for i := range s {
		if s[i] == '(' {
			stack++
		} else {
			stack--
		}
		if stack == 0 {
			if i-start+1 == 2 {
				sum += 1
			} else {
				sum += 2 * scoreOfParentheses(s[start+1:i])
			}
			start = i + 1
		}
	}
	return sum
}
