package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode-cn.com/problems/super-pow/

372. 超级次方
你的任务是计算 ab 对 1337 取模，a 是一个正整数，b 是一个非常大的正整数且会以数组形式给出。

示例 1：

输入：a = 2, b = [3]
输出：8
示例 2：

输入：a = 2, b = [1,0]
输出：1024
示例 3：

输入：a = 1, b = [4,3,3,8,5,2]
输出：1
示例 4：

输入：a = 2147483647, b = [2,0,0]
输出：1198

提示：

1 <= a <= 2^31 - 1
1 <= b.length <= 2000
0 <= b[i] <= 9
b 不含前导 0
*/
func main() {
	var tests = []struct {
		a    int
		b    []int
		want int
	}{
		{
			a:    2,
			b:    []int{3},
			want: 8,
		},
		{
			a:    2,
			b:    []int{1, 0},
			want: 1024,
		},
		{
			a:    1,
			b:    []int{4, 3, 3, 8, 5, 2},
			want: 1,
		},
		{
			a:    2147483647,
			b:    []int{2, 0, 0},
			want: 1198,
		},
	}

	for _, item := range tests {
		if ans := superPow(item.a, item.b); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

const M = 1337

func superPow(a int, b []int) int {
	return dfs(a, b, len(b)-1)
}

func dfs(a int, b []int, n int) int {
	if n < 0 {
		return 1
	}
	return pow(dfs(a, b, n-1), 10) * pow(a, b[n]) % M
}

func pow(a, b int) int {
	ans := 1
	a %= M
	for b > 0 {
		ans *= a
		ans %= M
		b--
	}
	return ans
}
