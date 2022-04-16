package main

import (
	"fmt"
	"math"
	"reflect"
)

/*
https://leetcode-cn.com/problems/largest-palindrome-product/

479. 最大回文数乘积
给定一个整数 n ，返回 可表示为两个 n 位整数乘积的 最大回文整数 。因为答案可能非常大，所以返回它对 1337 取余 。

示例 1:

输入：n = 2
输出：987
解释：99 x 91 = 9009, 9009 % 1337 = 987
示例 2:

输入： n = 1
输出： 9

提示:

1 <= n <= 8
*/
func main() {
	var tests = []struct {
		n    int
		want int
	}{
		{
			n:    2,
			want: 987,
		},
		{
			n:    1,
			want: 9,
		},
	}

	for _, item := range tests {
		if ans := largestPalindrome(item.n); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func largestPalindrome(n int) int {
	if n == 1 {
		return 9
	}

	upper := int(math.Pow10(n) - 1)
	for left := upper; ; left-- {
		p := left
		for right := left; right > 0; right /= 10 {
			p = p*10 + (right % 10)
		}

		for x := upper; x*x >= p; x-- {
			if p%x == 0 {
				return p % 1337
			}
		}
	}
}
