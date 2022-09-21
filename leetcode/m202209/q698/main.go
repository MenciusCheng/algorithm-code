package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/partition-to-k-equal-sum-subsets/

698. 划分为k个相等的子集
给定一个整数数组  nums 和一个正整数 k，找出是否有可能把这个数组分成 k 个非空子集，其总和都相等。

示例 1：

输入： nums = [4, 3, 2, 3, 5, 2, 1], k = 4
输出： True
说明： 有可能将其分成 4 个子集（5），（1,4），（2,3），（2,3）等于总和。
示例 2:

输入: nums = [1,2,3,4], k = 3
输出: false

提示：

1 <= k <= len(nums) <= 16
0 < nums[i] < 10000
每个元素的频率在 [1,4] 范围内
*/
func main() {
	var tests = []struct {
		nums []int
		k    int
		want bool
	}{
		{
			nums: []int{4, 3, 2, 3, 5, 2, 1},
			k:    4,
			want: true,
		},
		{
			nums: []int{1, 2, 3, 4},
			k:    3,
			want: false,
		},
		{
			nums: []int{4, 4, 6, 2, 3, 8, 10, 2, 10, 7},
			k:    4,
			want: true,
		},
		{
			nums: []int{18, 20, 39, 73, 96, 99, 101, 111, 114, 190, 207, 295, 471, 649, 700, 1037},
			k:    4,
			want: true,
		},
		{
			nums: []int{85, 35, 40, 64, 86, 45, 63, 16, 5364, 110, 5653, 97, 95},
			k:    7,
			want: false,
		},
		{
			nums: []int{3, 3, 10, 2, 6, 5, 10, 6, 8, 3, 2, 1, 6, 10, 7, 2},
			k:    6,
			want: true,
		},
		{
			nums: []int{2, 2, 2, 2, 3, 4, 5},
			k:    4,
			want: false,
		},
	}

	for _, item := range tests {
		if ans := canPartitionKSubsets(item.nums, item.k); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func canPartitionKSubsets(nums []int, k int) bool {
	return true
}
