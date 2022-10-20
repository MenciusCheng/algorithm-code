package main

import (
	"fmt"
	"math"
	"reflect"
)

/*
https://leetcode.cn/problems/k-th-symbol-in-grammar/

779. 第K个语法符号
中等

我们构建了一个包含 n 行( 索引从 1  开始 )的表。首先在第一行我们写上一个 0。接下来的每一行，将前一行中的0替换为01，1替换为10。

例如，对于 n = 3 ，第 1 行是 0 ，第 2 行是 01 ，第3行是 0110 。
给定行数 n 和序数 k，返回第 n 行中第 k 个字符。（ k 从索引 1 开始）

示例 1:

输入: n = 1, k = 1
输出: 0
解释: 第一行：0
示例 2:

输入: n = 2, k = 1
输出: 0
解释:
第一行: 0
第二行: 01
示例 3:

输入: n = 2, k = 2
输出: 1
解释:
第一行: 0
第二行: 01

提示:

1 <= n <= 30
1 <= k <= 2^n - 1
*/
func main() {
	var tests = []struct {
		n    int
		k    int
		want int
	}{
		{
			n:    1,
			k:    1,
			want: 0,
		},
		{
			n:    2,
			k:    1,
			want: 0,
		},
		{
			n:    2,
			k:    2,
			want: 1,
		},
	}

	for _, item := range tests {
		if ans := kthGrammar(item.n, item.k); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func kthGrammar(n int, k int) int {
	left := 1
	right := int(math.Pow(2, float64(n-1)))

	c := 0
	for i := 1; i < n; i++ {
		mid := (left + right) / 2
		if k <= mid {
			right = mid
		} else {
			if c == 0 {
				c = 1
			} else {
				c = 0
			}
			left = mid
		}
	}

	return c
}
