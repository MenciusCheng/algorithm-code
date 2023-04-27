package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/maximum-sum-of-two-non-overlapping-subarrays/

1031. 两个非重叠子数组的最大和
中等

给你一个整数数组 nums 和两个整数 firstLen 和 secondLen，请你找出并返回两个非重叠 子数组 中元素的最大和，长度分别为 firstLen 和 secondLen 。
长度为 firstLen 的子数组可以出现在长为 secondLen 的子数组之前或之后，但二者必须是不重叠的。
子数组是数组的一个 连续 部分。

示例 1：

输入：nums = [0,6,5,2,2,5,1,9,4], firstLen = 1, secondLen = 2
输出：20
解释：子数组的一种选择中，[9] 长度为 1，[6,5] 长度为 2。
示例 2：

输入：nums = [3,8,1,3,2,1,8,9,0], firstLen = 3, secondLen = 2
输出：29
解释：子数组的一种选择中，[3,8,1] 长度为 3，[8,9] 长度为 2。
示例 3：

输入：nums = [2,1,5,6,0,9,5,0,3,8], firstLen = 4, secondLen = 3
输出：31
解释：子数组的一种选择中，[5,6,0,9] 长度为 4，[0,3,8] 长度为 3。


提示：

1 <= firstLen, secondLen <= 1000
2 <= firstLen + secondLen <= 1000
firstLen + secondLen <= nums.length <= 1000
0 <= nums[i] <= 1000
*/
func main() {
	var tests = []struct {
		nums      []int
		firstLen  int
		secondLen int
		want      int
	}{
		{
			nums:      []int{0, 6, 5, 2, 2, 5, 1, 9, 4},
			firstLen:  1,
			secondLen: 2,
			want:      20,
		},
		{
			nums:      []int{3, 8, 1, 3, 2, 1, 8, 9, 0},
			firstLen:  3,
			secondLen: 2,
			want:      29,
		},
		{
			nums:      []int{2, 1, 5, 6, 0, 9, 5, 0, 3, 8},
			firstLen:  4,
			secondLen: 3,
			want:      31,
		},
		{
			nums:      []int{1, 0, 3},
			firstLen:  1,
			secondLen: 2,
			want:      4,
		},
	}

	for _, item := range tests {
		if ans := maxSumTwoNoOverlap(item.nums, item.firstLen, item.secondLen); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	sum1 := 0
	for i := 0; i < firstLen; i++ {
		sum1 += nums[i]
	}
	sum2 := findMaxSum(nums[firstLen:], secondLen)
	res := sum1 + sum2
	for i := firstLen; i < len(nums); i++ {
		sum1 += nums[i]
		sum1 -= nums[i-firstLen]
		t1 := findMaxSum(nums[:i-firstLen+1], secondLen)
		t2 := findMaxSum(nums[i+1:], secondLen)
		sum2 = max(t1, t2)
		res = max(res, sum1+sum2)
	}

	return res
}

func findMaxSum(nums []int, iLen int) int {
	if iLen > len(nums) {
		return 0
	}
	sum := 0
	for i := 0; i < iLen; i++ {
		sum += nums[i]
	}
	res := sum
	for i := iLen; i < len(nums); i++ {
		sum += nums[i]
		sum -= nums[i-iLen]
		res = max(res, sum)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
