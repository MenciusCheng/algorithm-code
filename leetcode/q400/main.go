package main

import (
	"fmt"
	"math"
	"reflect"
)

/*
https://leetcode-cn.com/problems/nth-digit/

400. 第 N 位数字
给你一个整数 n ，请你在无限的整数序列 [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ...] 中找出并返回第 n 位数字。

示例 1：

输入：n = 3
输出：3
示例 2：

输入：n = 11
输出：0
解释：第 11 位数字在序列 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ... 里是 0 ，它是 10 的一部分。

提示：

1 <= n <= 2^31 - 1
*/
func main() {
	var tests = []struct {
		n    int
		want int
	}{
		{
			n:    3,
			want: 3,
		},
		{
			n:    11,
			want: 0,
		},
	}

	for _, item := range tests {
		if ans := findNthDigit(item.n); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func findNthDigit(n int) int {
	count := 9
	length := 1

	for {
		sum := count * length
		if n <= sum {
			break
		}
		n -= sum
		count *= 10
		length += 1
	}

	nums := make([]int, length)
	for i := 1; i <= length; i++ {
		step := int(math.Pow10(length-i)) * length
		if i == 1 {
			var j int
			for j = 1; j <= 9; j++ {
				if n <= step {
					break
				} else {
					n -= step
				}
			}
			nums[i-1] = j
		} else {
			var j int
			for j = 0; j <= 9; j++ {
				if n <= step {
					break
				} else {
					n -= step
				}
			}
			nums[i-1] = j
		}
	}

	return nums[n-1]
}
