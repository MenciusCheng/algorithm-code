package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode-cn.com/problems/factorial-trailing-zeroes/

172. 阶乘后的零
给定一个整数 n ，返回 n! 结果中尾随零的数量。

提示 n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1

示例 1：

输入：n = 3
输出：0
解释：3! = 6 ，不含尾随 0
示例 2：

输入：n = 5
输出：1
解释：5! = 120 ，有一个尾随 0
示例 3：

输入：n = 0
输出：0

提示：

0 <= n <= 10^4

进阶：你可以设计并实现对数时间复杂度的算法来解决此问题吗？
*/
func main() {
	var tests = []struct {
		n    int
		want int
	}{
		{
			n:    3,
			want: 0,
		},
		{
			n:    5,
			want: 1,
		},
		{
			n:    30,
			want: 7,
		},
	}

	for _, item := range tests {
		if ans := trailingZeroes(item.n); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func trailingZeroes(n int) int {
	res := 0
	m := 5
	for n >= m {
		res += n / m
		m *= 5
	}

	return res
}
