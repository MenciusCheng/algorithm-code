package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/trapping-rain-water/description/

42. 接雨水
困难
4.5K
相关企业
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

示例 1：

输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
示例 2：

输入：height = [4,2,0,3,2,5]
输出：9

提示：

n == height.length
1 <= n <= 2 * 10^4
0 <= height[i] <= 10^5
*/
func main() {
	var tests = []struct {
		height []int
		want   int
	}{
		{
			height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want:   6,
		},
		{
			height: []int{4, 2, 0, 3, 2, 5},
			want:   9,
		},
	}

	for _, item := range tests {
		if ans := trap(item.height); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func trap(height []int) int {
	res := 0
	stack := make([]int, 0)

	for i := range height {
		if len(stack) == 0 || height[i] <= height[stack[len(stack)-1]] {
			stack = append(stack, i)
			continue
		}
		level := height[stack[len(stack)-1]]
		for len(stack) > 0 {
			si := stack[len(stack)-1]
			if height[si]-level > 0 {
				res += (i - si - 1) * (min(height[si], height[i]) - level)
			}
			level = height[si]
			if height[i] <= height[stack[len(stack)-1]] {
				break
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
