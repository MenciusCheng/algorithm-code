package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/maximum-value-at-a-given-index-in-a-bounded-array/

1802. 有界数组中指定下标处的最大值
中等

给你三个正整数 n、index 和 maxSum 。你需要构造一个同时满足下述所有条件的数组 nums（下标 从 0 开始 计数）：

nums.length == n
nums[i] 是 正整数 ，其中 0 <= i < n
abs(nums[i] - nums[i+1]) <= 1 ，其中 0 <= i < n-1
nums 中所有元素之和不超过 maxSum
nums[index] 的值被 最大化
返回你所构造的数组中的 nums[index] 。

注意：abs(x) 等于 x 的前提是 x >= 0 ；否则，abs(x) 等于 -x 。

示例 1：

输入：n = 4, index = 2,  maxSum = 6
输出：2
解释：数组 [1,1,2,1] 和 [1,2,2,1] 满足所有条件。不存在其他在指定下标处具有更大值的有效数组。
示例 2：

输入：n = 6, index = 1,  maxSum = 10
输出：3

提示：

1 <= n <= maxSum <= 10^9
0 <= index < n
*/
func main() {
	var tests = []struct {
		n      int
		index  int
		maxSum int
		want   int
	}{
		{
			n:      4,
			index:  0,
			maxSum: 4,
			want:   1,
		},
		{
			n:      1,
			index:  0,
			maxSum: 24,
			want:   24,
		},
		{
			n:      8,
			index:  7,
			maxSum: 14,
			want:   4,
		},
		{
			n:      4,
			index:  2,
			maxSum: 6,
			want:   2,
		},
		{
			n:      6,
			index:  1,
			maxSum: 10,
			want:   3,
		},
	}

	for _, item := range tests {
		if ans := maxValue(item.n, item.index, item.maxSum); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func maxValue(n int, index int, maxSum int) int {
	k1 := 1
	k2 := maxSum
	for k1 <= k2 {
		mk := (k1 + k2) / 2
		s := sum(mk, n, index)
		if s == maxSum {
			return mk
		}

		if s > maxSum {
			k2 = mk - 1
		} else {
			k1 = mk + 1
		}
	}
	return k2
}

func sum(k, n, i int) int {
	left := 0
	if k-i > 0 {
		left = ((k - i) + (k - 1)) * i / 2
	} else {
		left = (1+k-1)*(k-1)/2 + (i - (k - 1))
	}

	right := 0
	if k-(n-1-i) > 0 {
		right = (k + (k - (n - 1 - i))) * (n - i) / 2
	} else {
		right = (k+1)*k/2 + (n - i - k)
	}

	return left + right
}
