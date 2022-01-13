package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode-cn.com/problems/largest-number-at-least-twice-of-others/

747. 至少是其他数字两倍的最大数
给你一个整数数组 nums ，其中总是存在 唯一的 一个最大整数 。
请你找出数组中的最大元素并检查它是否 至少是数组中每个其他数字的两倍 。如果是，则返回 最大元素的下标 ，否则返回 -1 。

示例 1：

输入：nums = [3,6,1,0]
输出：1
解释：6 是最大的整数，对于数组中的其他整数，6 大于数组中其他元素的两倍。6 的下标是 1 ，所以返回 1 。
示例 2：

输入：nums = [1,2,3,4]
输出：-1
解释：4 没有超过 3 的两倍大，所以返回 -1 。
示例 3：

输入：nums = [1]
输出：0
解释：因为不存在其他数字，所以认为现有数字 1 至少是其他数字的两倍。

提示：

1 <= nums.length <= 50
0 <= nums[i] <= 100
nums 中的最大元素是唯一的
*/
func main() {
	var tests = []struct {
		nums []int
		want int
	}{
		{
			nums: []int{3, 6, 1, 0},
			want: 1,
		},
		{
			nums: []int{1, 2, 3, 4},
			want: -1,
		},
		{
			nums: []int{1},
			want: 0,
		},
	}

	for _, item := range tests {
		if ans := dominantIndex(item.nums); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func dominantIndex(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	var top1, top2 int
	var top1index int
	if nums[0] > nums[1] {
		top1 = nums[0]
		top2 = nums[1]
		top1index = 0
	} else {
		top1 = nums[1]
		top2 = nums[0]
		top1index = 1
	}

	for i := 2; i < len(nums); i++ {
		if nums[i] > top1 {
			top2 = top1
			top1 = nums[i]
			top1index = i
		} else if nums[i] > top2 {
			top2 = nums[i]
		}
	}

	if top1 >= top2*2 {
		return top1index
	}
	return -1
}
