package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode-cn.com/problems/count-numbers-with-unique-digits/

357. 统计各位数字都不同的数字个数
给你一个整数 n ，统计并返回各位数字都不同的数字 x 的个数，其中 0 <= x < 10^n 。

示例 1：

输入：n = 2
输出：91
解释：答案应为除去 11、22、33、44、55、66、77、88、99 外，在 0 ≤ x < 100 范围内的所有数字。
示例 2：

输入：n = 0
输出：1

提示：

0 <= n <= 8
*/
func main() {
	var tests = []struct {
		n    int
		want int
	}{
		{
			n:    2,
			want: 91,
		},
		{
			n:    0,
			want: 1,
		},
	}

	for _, item := range tests {
		if ans := countNumbersWithUniqueDigits(item.n); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func countNumbersWithUniqueDigits(n int) int {
	res := 0
	for i := 0; i <= n; i++ {
		res += lenCount(i)
	}
	return res
}

func lenCount(l int) int {
	switch l {
	case 0:
		return 1
	case 1:
		return 9
	default:
		count := 9
		for i := 1; i < l; i++ {
			count *= 10 - i
		}
		return count
	}
}
