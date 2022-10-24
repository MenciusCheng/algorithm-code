package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/partition-array-into-disjoint-intervals/

915. 分割数组
中等

给定一个数组 nums ，将其划分为两个连续子数组 left 和 right， 使得：

left 中的每个元素都小于或等于 right 中的每个元素。
left 和 right 都是非空的。
left 的长度要尽可能小。
在完成这样的分组后返回 left 的 长度 。

用例可以保证存在这样的划分方法。

示例 1：

输入：nums = [5,0,3,8,6]
输出：3
解释：left = [5,0,3]，right = [8,6]
示例 2：

输入：nums = [1,1,1,0,6,12]
输出：4
解释：left = [1,1,1,0]，right = [6,12]

提示：

2 <= nums.length <= 10^5
0 <= nums[i] <= 10^6
可以保证至少有一种方法能够按题目所描述的那样对 nums 进行划分。
*/
func main() {
	var tests = []struct {
		nums []int
		want int
	}{
		{
			nums: []int{5, 0, 3, 8, 6},
			want: 3,
		},
		{
			nums: []int{1, 1, 1, 0, 6, 12},
			want: 4,
		},
	}

	for _, item := range tests {
		if ans := partitionDisjoint(item.nums); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func partitionDisjoint(nums []int) int {
	minArr := make([]int, len(nums))
	minArr[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 1; i-- {
		minArr[i] = min(nums[i], minArr[i+1])
	}

	leftMax := nums[0]
	if leftMax <= minArr[1] {
		return 1
	}

	for i := 1; i < len(nums)-1; i++ {
		leftMax = max(leftMax, nums[i])
		if leftMax <= minArr[i+1] {
			return i + 1
		}
	}

	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
