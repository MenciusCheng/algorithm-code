package main

import (
	"fmt"
	"reflect"
	"sort"
)

/*
https://leetcode.cn/problems/wiggle-sort-ii/

324. 摆动排序 II
给你一个整数数组 nums，将它重新排列成 nums[0] < nums[1] > nums[2] < nums[3]... 的顺序。
你可以假设所有输入数组都可以得到满足题目要求的结果。

示例 1：

输入：nums = [1,5,1,1,6,4]
输出：[1,6,1,5,1,4]
解释：[1,4,1,5,1,6] 同样是符合题目要求的结果，可以被判题程序接受。
示例 2：

输入：nums = [1,3,2,2,3,1]
输出：[2,3,1,3,1,2]

提示：

1 <= nums.length <= 5 * 10^4
0 <= nums[i] <= 5000
题目数据保证，对于给定的输入 nums ，总能产生满足题目要求的结果
*/
func main() {
	var tests = []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{1, 5, 1, 1, 6, 4},
			want: []int{1, 6, 1, 5, 1, 4},
		},
		{
			nums: []int{1, 3, 2, 2, 3, 1},
			want: []int{2, 3, 1, 3, 1, 2},
		},
	}

	for _, item := range tests {
		if wiggleSort(item.nums); reflect.DeepEqual(item.nums, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", item.nums, item.want)
		}
	}
}

func wiggleSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	sort.Ints(nums)
	i1 := len(nums) - 1
	i2 := len(nums) - 2
	if len(nums)%2 == 0 {
		i1, i2 = i2, i1
	}

	j := 0
	arr2 := make([]int, len(nums))
	for i1 >= 0 {
		arr2[i1] = nums[j]
		i1 -= 2
		j++
	}
	for i2 >= 0 {
		arr2[i2] = nums[j]
		i2 -= 2
		j++
	}

	for i, num := range arr2 {
		nums[i] = num
	}
}
