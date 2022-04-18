package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode-cn.com/problems/lexicographical-numbers/

386. 字典序排数
给你一个整数 n ，按字典序返回范围 [1, n] 内所有整数。
你必须设计一个时间复杂度为 O(n) 且使用 O(1) 额外空间的算法。

示例 1：

输入：n = 13
输出：[1,10,11,12,13,2,3,4,5,6,7,8,9]
示例 2：

输入：n = 2
输出：[1,2]

提示：

1 <= n <= 5 * 104
*/
func main() {
	var tests = []struct {
		n    int
		want []int
	}{
		{
			n:    13,
			want: []int{1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			n:    2,
			want: []int{1, 2},
		},
	}

	for _, item := range tests {
		if ans := lexicalOrder(item.n); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func lexicalOrder(n int) []int {
	res := make([]int, n)
	num := 1

	for i := range res {
		res[i] = num
		if num*10 <= n {
			num *= 10
		} else {
			for num%10 == 9 || num == n {
				num /= 10
			}
			num++
		}
	}
	return res
}
