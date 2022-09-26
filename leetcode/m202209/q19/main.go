package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/missing-two-lcci/

19. 消失的两个数字
给定一个数组，包含从 1 到 N 所有的整数，但其中缺了两个数字。你能在 O(N) 时间内只用 O(1) 的空间找到它们吗？

以任意顺序返回这两个数字均可。

示例 1:

输入: [1]
输出: [2,3]
示例 2:

输入: [2,3]
输出: [1,4]
提示：

nums.length <= 30000
*/
func main() {
	var tests = []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{1},
			want: []int{2, 3},
		},
		{
			nums: []int{2, 3},
			want: []int{1, 4},
		},
	}

	for _, item := range tests {
		if ans := missingTwo(item.nums); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func missingTwo(nums []int) []int {
	xorSum := 0
	n := len(nums) + 2
	for _, num := range nums {
		xorSum ^= num
	}
	for i := 1; i <= n; i++ {
		xorSum ^= i
	}
	lsb := xorSum & -xorSum
	type1, type2 := 0, 0
	for _, num := range nums {
		if num&lsb > 0 {
			type1 ^= num
		} else {
			type2 ^= num
		}
	}
	for i := 1; i <= n; i++ {
		if i&lsb > 0 {
			type1 ^= i
		} else {
			type2 ^= i
		}
	}
	return []int{type1, type2}
}
