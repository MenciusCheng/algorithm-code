package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode-cn.com/problems/perfect-number/

507. 完美数
对于一个 正整数，如果它和除了它自身以外的所有 正因子 之和相等，我们称它为 「完美数」。
给定一个 整数 n， 如果是完美数，返回 true，否则返回 false

示例 1：

输入：num = 28
输出：true
解释：28 = 1 + 2 + 4 + 7 + 14
1, 2, 4, 7, 和 14 是 28 的所有正因子。
示例 2：

输入：num = 6
输出：true
示例 3：

输入：num = 496
输出：true
示例 4：

输入：num = 8128
输出：true
示例 5：

输入：num = 2
输出：false

提示：

1 <= num <= 10^8
*/
func main() {
	var tests = []struct {
		num  int
		want bool
	}{
		{
			num:  28,
			want: true,
		},
		{
			num:  6,
			want: true,
		},
		{
			num:  496,
			want: true,
		},
		{
			num:  8128,
			want: true,
		},
		{
			num:  2,
			want: false,
		},
	}

	for _, item := range tests {
		if ans := checkPerfectNumber(item.num); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}
	sum := 1
	k := num
	i := 2
	for i < k {
		if num%i == 0 {
			sum += i
			j := num / i
			if j != i {
				sum += j
			}
			k = j
		}
		i++
	}

	return num == sum
}
