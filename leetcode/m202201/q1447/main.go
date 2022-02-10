package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode-cn.com/problems/simplified-fractions/

1447. 最简分数
给你一个整数 n ，请你返回所有 0 到 1 之间（不包括 0 和 1）满足分母小于等于 n 的 最简 分数 。分数可以以 任意 顺序返回。

示例 1：

输入：n = 2
输出：["1/2"]
解释："1/2" 是唯一一个分母小于等于 2 的最简分数。
示例 2：

输入：n = 3
输出：["1/2","1/3","2/3"]
示例 3：

输入：n = 4
输出：["1/2","1/3","1/4","2/3","3/4"]
解释："2/4" 不是最简分数，因为它可以化简为 "1/2" 。
示例 4：

输入：n = 1
输出：[]

提示：

1 <= n <= 100
*/
func main() {
	var tests = []struct {
		n    int
		want []string
	}{
		{
			n:    2,
			want: []string{"1/2"},
		},
		{
			n:    3,
			want: []string{"1/2", "1/3", "2/3"},
		},
		{
			n:    4,
			want: []string{"1/2", "1/3", "1/4", "2/3", "3/4"},
		},
		{
			n:    1,
			want: []string{},
		},
	}

	for _, item := range tests {
		if ans := simplifiedFractions(item.n); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func simplifiedFractions(n int) []string {
	res := make([]string, 0)

	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			if gcd(i, j) == 1 {
				res = append(res, fmt.Sprintf("%d/%d", j, i))
			}
		}
	}
	return res
}

// 最大公约数
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
