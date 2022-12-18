package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/minimum-adjacent-swaps-for-k-consecutive-ones/

1703. 得到连续 K 个 1 的最少相邻交换次数
困难

给你一个整数数组 nums 和一个整数 k 。 nums 仅包含 0 和 1 。每一次移动，你可以选择 相邻 两个数字并将它们交换。
请你返回使 nums 中包含 k 个 连续 1 的 最少 交换次数。

示例 1：

输入：nums = [1,0,0,1,0,1], k = 2
输出：1
解释：在第一次操作时，nums 可以变成 [1,0,0,0,1,1] 得到连续两个 1 。
示例 2：

输入：nums = [1,0,0,0,0,0,1,1], k = 3
输出：5
解释：通过 5 次操作，最左边的 1 可以移到右边直到 nums 变为 [0,0,0,0,0,1,1,1] 。
示例 3：

输入：nums = [1,1,0,1], k = 2
输出：0
解释：nums 已经有连续 2 个 1 了。

提示：

1 <= nums.length <= 10^5
nums[i] 要么是 0 ，要么是 1 。
1 <= k <= sum(nums)
*/
func main() {
	var tests = []struct {
		nums []int
		k    int
		want int
	}{
		{
			nums: []int{0, 1, 0, 0, 1, 1, 0, 1, 0, 1, 1, 1, 1, 0, 1, 0},
			k:    7,
			want: 4,
		},
	}

	for _, item := range tests {
		if ans := minMoves(item.nums, item.k); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func minMoves(nums []int, k int) int {
	arr := make([]int, 0)
	for i, item := range nums {
		if item == 1 {
			ds := i - len(arr)
			arr = append(arr, ds)
		}
	}
	var left, right int
	mk := k / 2
	rc := k - mk
	for i := 0; i < mk; i++ {
		left += arr[i]
	}
	for i := mk; i < k; i++ {
		right += arr[i]
	}
	res := arr[mk]*mk - left + right - arr[mk]*rc
	for i := mk; i < len(arr)-rc; i++ {
		left += arr[i] - arr[i-mk]
		right += arr[i+rc] - arr[i]
		res = min(res, arr[i+1]*mk-left+right-arr[i+1]*rc)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
