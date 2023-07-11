package main

import (
	"fmt"
	"reflect"
	"sort"
)

/*
https://leetcode.cn/problems/3sum-closest/

16. 最接近的三数之和
中等
1.5K
相关企业
给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
返回这三个数的和。
假定每组输入只存在恰好一个解。

示例 1：

输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
示例 2：

输入：nums = [0,0,0], target = 1
输出：0

提示：

3 <= nums.length <= 1000
-1000 <= nums[i] <= 1000
-10^4 <= target <= 10^4
*/
func main() {
	var tests = []struct {
		nums   []int
		target int
		want   int
	}{
		{
			nums:   []int{-1, 2, 1, -4},
			target: 1,
			want:   2,
		},
		{
			nums:   []int{0, 0, 0},
			target: 1,
			want:   0,
		},
	}

	for _, item := range tests {
		if ans := threeSumClosest(item.nums, item.target); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		t := target - nums[i]
		j1 := i + 1
		j2 := len(nums) - 1
		for j1 < j2 {
			v := nums[j1] + nums[j2]
			if v == t {
				return v + nums[i]
			}
			if abs(v, t) < abs(res, target) {
				res = v + nums[i]
			}

			if abs(nums[j1+1]+nums[j2], t) < abs(nums[j1]+nums[j2-1], t) {
				j1++
			} else {
				j2--
			}
		}
	}
	return res
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
