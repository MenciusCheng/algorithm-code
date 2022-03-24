package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode-cn.com/problems/k-th-smallest-in-lexicographical-order/

440. 字典序的第K小数字
给定整数 n 和 k，返回  [1, n] 中字典序第 k 小的数字。

示例 1:

输入: n = 13, k = 2
输出: 10
解释: 字典序的排列是 [1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9]，所以第二小的数字是 10。
示例 2:

输入: n = 1, k = 1
输出: 1

提示:

1 <= k <= n <= 10^9
*/
func main() {
	var tests = []struct {
		n    int
		k    int
		want int
	}{
		{
			n:    13,
			k:    2,
			want: 10,
		},
		{
			n:    1,
			k:    1,
			want: 1,
		},
		{
			n:    10,
			k:    3,
			want: 2,
		},
		{
			n:    100,
			k:    10,
			want: 17,
		},
	}

	for _, item := range tests {
		if ans := findKthNumber(item.n, item.k); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func getSteps(cur, n int) (steps int) {
	first, last := cur, cur
	for first <= n {
		steps += min(last, n) - first + 1
		first *= 10
		last = last*10 + 9
	}
	return
}

func findKthNumber(n, k int) int {
	cur := 1
	k--
	for k > 0 {
		steps := getSteps(cur, n)
		if steps <= k {
			k -= steps
			cur++
		} else {
			cur *= 10
			k--
		}
	}
	return cur
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
