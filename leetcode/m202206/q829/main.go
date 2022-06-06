package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/consecutive-numbers-sum/

829. 连续整数求和
给定一个正整数 n，返回 连续正整数满足所有数字之和为 n 的组数 。

示例 1:

输入: n = 5
输出: 2
解释: 5 = 2 + 3，共有两组连续整数([5],[2,3])求和后为 5。
示例 2:

输入: n = 9
输出: 3
解释: 9 = 4 + 5 = 2 + 3 + 4
示例 3:

输入: n = 15
输出: 4
解释: 15 = 8 + 7 = 4 + 5 + 6 = 1 + 2 + 3 + 4 + 5

提示:

1 <= n <= 10^9
*/
func main() {
	var tests = []struct {
		n    int
		want int
	}{
		{
			n:    5,
			want: 2,
		},
		{
			n:    9,
			want: 3,
		},
	}

	for _, item := range tests {
		if ans := consecutiveNumbersSum(item.n); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func isKConsecutive(n, k int) bool {
	if k%2 == 1 {
		return n%k == 0
	}
	return n%k != 0 && 2*n%k == 0
}

func consecutiveNumbersSum(n int) (ans int) {
	for k := 1; k*(k+1) <= n*2; k++ {
		if isKConsecutive(n, k) {
			ans++
		}
	}
	return
}
