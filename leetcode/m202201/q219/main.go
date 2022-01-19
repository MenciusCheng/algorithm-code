package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode-cn.com/problems/contains-duplicate-ii/

219. 存在重复元素 II
给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个 不同的索引 i 和 j ，满足 nums[i] == nums[j] 且 abs(i - j) <= k 。如果存在，返回 true ；否则，返回 false 。

示例 1：

输入：nums = [1,2,3,1], k = 3
输出：true
示例 2：

输入：nums = [1,0,1,1], k = 1
输出：true
示例 3：

输入：nums = [1,2,3,1,2,3], k = 2
输出：false

提示：

1 <= nums.length <= 10^5
-10^9 <= nums[i] <= 10^9
0 <= k <= 10^5
*/
func main() {
	var tests = []struct {
		nums []int
		k    int
		want bool
	}{
		{
			nums: []int{1, 2, 3, 1},
			k:    3,
			want: true,
		},
		{
			nums: []int{1, 0, 1, 1},
			k:    1,
			want: true,
		},
		{
			nums: []int{1, 2, 3, 1, 2, 3},
			k:    2,
			want: false,
		},
	}

	for _, item := range tests {
		if ans := containsNearbyDuplicate(item.nums, item.k); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func containsNearbyDuplicate(nums []int, k int) bool {
	if k == 0 {
		return false
	}

	cnt := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if i > k {
			cnt[nums[i-k-1]]--
		}
		if cnt[nums[i]] > 0 {
			return true
		}
		cnt[nums[i]]++
	}

	return false
}
