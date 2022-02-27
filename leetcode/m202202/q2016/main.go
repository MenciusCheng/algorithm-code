package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode-cn.com/problems/maximum-difference-between-increasing-elements/

2016. 增量元素之间的最大差值
给你一个下标从 0 开始的整数数组 nums ，该数组的大小为 n ，请你计算 nums[j] - nums[i] 能求得的 最大差值 ，其中 0 <= i < j < n 且 nums[i] < nums[j] 。

返回 最大差值 。如果不存在满足要求的 i 和 j ，返回 -1 。

示例 1：

输入：nums = [7,1,5,4]
输出：4
解释：
最大差值出现在 i = 1 且 j = 2 时，nums[j] - nums[i] = 5 - 1 = 4 。
注意，尽管 i = 1 且 j = 0 时 ，nums[j] - nums[i] = 7 - 1 = 6 > 4 ，但 i > j 不满足题面要求，所以 6 不是有效的答案。
示例 2：

输入：nums = [9,4,3,2]
输出：-1
解释：
不存在同时满足 i < j 和 nums[i] < nums[j] 这两个条件的 i, j 组合。
示例 3：

输入：nums = [1,5,2,10]
输出：9
解释：
最大差值出现在 i = 0 且 j = 3 时，nums[j] - nums[i] = 10 - 1 = 9 。

提示：

n == nums.length
2 <= n <= 1000
1 <= nums[i] <= 10^9
*/
func main() {
	var tests = []struct {
		nums []int
		want int
	}{
		{
			nums: []int{7, 1, 5, 4},
			want: 4,
		},
		{
			nums: []int{9, 4, 3, 2},
			want: -1,
		},
		{
			nums: []int{1, 5, 2, 10},
			want: 9,
		},
	}

	for _, item := range tests {
		if ans := maximumDifference(item.nums); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func maximumDifference(nums []int) int {
	max := -1
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] && nums[j]-nums[i] > max {
				max = nums[j] - nums[i]
			}
		}
	}
	return max
}
