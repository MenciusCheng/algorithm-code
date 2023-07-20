package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/maximum-sum-circular-subarray/

918. 环形子数组的最大和
提示
中等
576
相关企业
给定一个长度为 n 的环形整数数组 nums ，返回 nums 的非空 子数组 的最大可能和 。
环形数组 意味着数组的末端将会与开头相连呈环状。形式上， nums[i] 的下一个元素是 nums[(i + 1) % n] ， nums[i] 的前一个元素是 nums[(i - 1 + n) % n] 。
子数组 最多只能包含固定缓冲区 nums 中的每个元素一次。形式上，对于子数组 nums[i], nums[i + 1], ..., nums[j] ，不存在 i <= k1, k2 <= j 其中 k1 % n == k2 % n 。

示例 1：
输入：nums = [1,-2,3,-2]
输出：3
解释：从子数组 [3] 得到最大和 3

示例 2：
输入：nums = [5,-3,5]
输出：10
解释：从子数组 [5,5] 得到最大和 5 + 5 = 10

示例 3：
输入：nums = [3,-2,2,-3]
输出：3
解释：从子数组 [3] 和 [3,-2,2] 都可以得到最大和 3

提示：

n == nums.length
1 <= n <= 3 * 10^4
-3 * 10^4 <= nums[i] <= 3 * 10^4
*/
func main() {
	var tests = []struct {
		nums []int
		want int
	}{
		{nums: []int{-2}, want: -2},
		{nums: []int{-3, -2, -3}, want: -2},
		{nums: []int{1, -2, 3, -2}, want: 3},
		{nums: []int{5, -3, 5}, want: 10},
		{nums: []int{3, -2, 2, -3}, want: 3},
	}

	for _, item := range tests {
		if ans := maxSubarraySumCircular(item.nums); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func maxSubarraySumCircular(nums []int) int {
	sMax := findSubSum(nums, max)
	if len(nums) > 2 {
		sMin := findSubSum(nums[1:len(nums)-1], min)
		var all int
		for _, num := range nums {
			all += num
		}
		sMax = max(sMax, all-sMin)
	}

	return sMax
}

func findSubSum(nums []int, f func(a, b int) int) int {
	res := nums[0]
	pre := 0
	for i := 0; i < len(nums); i++ {
		pre = f(nums[i], nums[i]+pre)
		res = f(pre, res)
	}
	return res
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
