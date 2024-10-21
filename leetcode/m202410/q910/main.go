package main

import (
	"fmt"
	"reflect"
	"slices"
)

/*
https://leetcode.cn/problems/smallest-range-ii/?envType=daily-question&envId=2024-10-21

910. 最小差值 II
中等

给你一个整数数组 nums，和一个整数 k 。
对于每个下标 i（0 <= i < nums.length），将 nums[i] 变成 nums[i] + k 或 nums[i] - k 。
nums 的 分数 是 nums 中最大元素和最小元素的差值。
在更改每个下标对应的值之后，返回 nums 的最小 分数 。

示例 1：

输入：nums = [1], k = 0
输出：0
解释：分数 = max(nums) - min(nums) = 1 - 1 = 0 。
示例 2：

输入：nums = [0,10], k = 2
输出：6
解释：将数组变为 [2, 8] 。分数 = max(nums) - min(nums) = 8 - 2 = 6 。
示例 3：

输入：nums = [1,3,6], k = 3
输出：3
解释：将数组变为 [4, 6, 3] 。分数 = max(nums) - min(nums) = 6 - 3 = 3 。

提示：

1 <= nums.length <= 10^4
0 <= nums[i] <= 10^4
0 <= k <= 10^4
*/
func main() {
	var tests = []struct {
		nums []int
		k    int
		want int
	}{
		{
			nums: []int{1},
			k:    0,
			want: 0,
		},
		{
			nums: []int{0, 10},
			k:    2,
			want: 6,
		},
		{
			nums: []int{1, 3, 6},
			k:    3,
			want: 3,
		},
	}

	for _, item := range tests {
		if ans := smallestRangeII(item.nums, item.k); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func smallestRangeII(nums []int, k int) int {
	slices.Sort(nums)
	n := len(nums)
	ans := nums[n-1] - nums[0]
	for i := 1; i < n; i++ {
		mx := max(nums[i-1]+k, nums[n-1]-k)
		mn := min(nums[0]+k, nums[i]-k)
		ans = min(ans, mx-mn)
	}
	return ans
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
