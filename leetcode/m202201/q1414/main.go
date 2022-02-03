package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode-cn.com/problems/find-the-minimum-number-of-fibonacci-numbers-whose-sum-is-k/

1414. 和为 K 的最少斐波那契数字数目
给你数字 k ，请你返回和为 k 的斐波那契数字的最少数目，其中，每个斐波那契数字都可以被使用多次。

斐波那契数字定义为：

F1 = 1
F2 = 1
Fn = Fn-1 + Fn-2 ， 其中 n > 2 。
数据保证对于给定的 k ，一定能找到可行解。

示例 1：

输入：k = 7
输出：2
解释：斐波那契数字为：1，1，2，3，5，8，13，……
对于 k = 7 ，我们可以得到 2 + 5 = 7 。
示例 2：

输入：k = 10
输出：2
解释：对于 k = 10 ，我们可以得到 2 + 8 = 10 。
示例 3：

输入：k = 19
输出：3
解释：对于 k = 19 ，我们可以得到 1 + 5 + 13 = 19 。

提示：

1 <= k <= 10^9
*/
func main() {
	var tests = []struct {
		k    int
		want int
	}{
		{
			k:    7,
			want: 2,
		},
		{
			k:    10,
			want: 2,
		},
		{
			k:    19,
			want: 3,
		},
		{
			k:    427010,
			want: 10,
		},
	}

	for _, item := range tests {
		if ans := findMinFibonacciNumbers(item.k); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func findMinFibonacciNumbers(k int) int {
	fa := make([]int, 0)
	a1, a2 := 1, 1
	for a2 <= k {
		a1, a2 = a2, a1+a2
		fa = append(fa, a1)
	}

	count := 0
	for k > 0 {
		for k < fa[len(fa)-1] {
			fa = fa[:len(fa)-1]
		}
		k -= fa[len(fa)-1]
		count++
	}

	return count
}
