package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/minimum-seconds-to-equalize-a-circular-array/description/?envType=daily-question&envId=2024-01-30

2808. 使循环数组所有元素相等的最少秒数
中等

给你一个下标从 0 开始长度为 n 的数组 nums 。
每一秒，你可以对数组执行以下操作：
对于范围在 [0, n - 1] 内的每一个下标 i ，将 nums[i] 替换成 nums[i] ，nums[(i - 1 + n) % n] 或者 nums[(i + 1) % n] 三者之一。
注意，所有元素会被同时替换。
请你返回将数组 nums 中所有元素变成相等元素所需要的 最少 秒数。

示例 1：

输入：nums = [1,2,1,2]
输出：1
解释：我们可以在 1 秒内将数组变成相等元素：
- 第 1 秒，将每个位置的元素分别变为 [nums[3],nums[1],nums[3],nums[3]] 。变化后，nums = [2,2,2,2] 。
1 秒是将数组变成相等元素所需要的最少秒数。
示例 2：

输入：nums = [2,1,3,3,2]
输出：2
解释：我们可以在 2 秒内将数组变成相等元素：
- 第 1 秒，将每个位置的元素分别变为 [nums[0],nums[2],nums[2],nums[2],nums[3]] 。变化后，nums = [2,3,3,3,3] 。
- 第 2 秒，将每个位置的元素分别变为 [nums[1],nums[1],nums[2],nums[3],nums[4]] 。变化后，nums = [3,3,3,3,3] 。
2 秒是将数组变成相等元素所需要的最少秒数。
示例 3：

输入：nums = [5,5,5,5]
输出：0
解释：不需要执行任何操作，因为一开始数组中的元素已经全部相等。

提示：

1 <= n == nums.length <= 10^5
1 <= nums[i] <= 10^9
*/
func main() {
	var tests = []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 2, 1, 2},
			want: 1,
		},
		{
			nums: []int{2, 1, 3, 3, 2},
			want: 2,
		},
		{
			nums: []int{5, 5, 5, 5},
			want: 0,
		},
		{
			nums: []int{11, 4, 10},
			want: 1,
		},
	}

	for _, item := range tests {
		if ans := minimumSeconds(item.nums); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func minimumSeconds(nums []int) int {
	lastIndex := make(map[int]int)
	dsMap := make(map[int]int)

	for i := 0; i < len(nums)*2; i++ {
		v := nums[i%len(nums)]
		if la, ok := lastIndex[v]; ok {
			dsMap[v] = max(dsMap[v], i-la)
			lastIndex[v] = i
		}
		lastIndex[v] = i
	}
	ds := len(nums)
	for _, v := range dsMap {
		ds = min(ds, v)
	}

	return ds / 2
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
