package main

import (
	"fmt"
	"reflect"
	"sort"
)

/*
https://leetcode-cn.com/problems/single-element-in-a-sorted-array/

540. 有序数组中的单一元素
给你一个仅由整数组成的有序数组，其中每个元素都会出现两次，唯有一个数只会出现一次。
请你找出并返回只出现一次的那个数。
你设计的解决方案必须满足 O(log n) 时间复杂度和 O(1) 空间复杂度。

示例 1:

输入: nums = [1,1,2,3,3,4,4,8,8]
输出: 2
示例 2:

输入: nums =  [3,3,7,7,10,11,11]
输出: 10

提示:

1 <= nums.length <= 10^5
0 <= nums[i] <= 10^5
*/
func main() {
	var tests = []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 1, 2, 3, 3, 4, 4, 8, 8},
			want: 2,
		},
		{
			nums: []int{3, 3, 7, 7, 10, 11, 11},
			want: 10,
		},
	}

	for _, item := range tests {
		if ans := singleNonDuplicate(item.nums); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func singleNonDuplicate(nums []int) int {
	sort.Ints(nums)
	lastNum := nums[0]
	lastCount := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == lastNum {
			lastCount++
		} else {
			if lastCount == 1 {
				return lastNum
			} else {
				lastCount = 1
				lastNum = nums[i]
			}
		}
	}

	return nums[len(nums)-1]
}
