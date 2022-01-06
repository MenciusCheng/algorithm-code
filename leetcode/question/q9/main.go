package main

import (
	"fmt"
	"reflect"
	"strconv"
)

/*
https://leetcode-cn.com/problems/palindrome-number/

9. 回文数
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。

示例 1：

输入：x = 121
输出：true
示例 2：

输入：x = -121
输出：false
解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3：

输入：x = 10
输出：false
解释：从右向左读, 为 01 。因此它不是一个回文数。
示例 4：

输入：x = -101
输出：false

提示：

-2^31 <= x <= 2^31 - 1
*/
func main() {
	var tests = []struct {
		x    int
		want bool
	}{
		{
			x:    121,
			want: true,
		},
		{
			x:    10,
			want: false,
		},
	}

	for _, item := range tests {
		if ans := isPalindrome(item.x); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	}
	s := strconv.Itoa(x)
	i := 0
	j := len(s) - 1
	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}
