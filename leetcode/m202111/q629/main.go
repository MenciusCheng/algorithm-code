package main

import "fmt"

/*
https://leetcode-cn.com/problems/k-inverse-pairs-array/

困难

629. K个逆序对数组
给出两个整数 n 和 k，找出所有包含从 1 到 n 的数字，且恰好拥有 k 个逆序对的不同的数组的个数。

逆序对的定义如下：对于数组的第i个和第 j个元素，如果满i < j且 a[i] > a[j]，则其为一个逆序对；否则不是。

由于答案可能很大，只需要返回 答案 mod 109 + 7 的值。
*/

func main() {
	fmt.Println(kInversePairs(3, 0) == 1)
	fmt.Println(kInversePairs(3, 1) == 2)
	fmt.Println(kInversePairs(3, 2) == 2)
	fmt.Println(kInversePairs(3, 3) == 1)
}

// 动态规划，TODO
func kInversePairs(n int, k int) int {
	const mod int = 1e9 + 7
	f := [2][]int{make([]int, k+1), make([]int, k+1)}
	f[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j <= k; j++ {
			cur := i & 1
			prev := cur ^ 1
			f[cur][j] = 0
			if j > 0 {
				f[cur][j] = f[cur][j-1]
			}
			if j >= i {
				f[cur][j] -= f[prev][j-i]
			}
			f[cur][j] += f[prev][j]
			if f[cur][j] >= mod {
				f[cur][j] -= mod
			} else if f[cur][j] < 0 {
				f[cur][j] += mod
			}
		}
	}
	return f[n&1][k]
}
