package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/count-different-palindromic-subsequences/

730. 统计不同回文子序列
给定一个字符串 s，返回 s 中不同的非空「回文子序列」个数 。
通过从 s 中删除 0 个或多个字符来获得子序列。
如果一个字符序列与它反转后的字符序列一致，那么它是「回文字符序列」。
如果有某个 i , 满足 ai != bi ，则两个序列 a1, a2, ... 和 b1, b2, ... 不同。

注意：

结果可能很大，你需要对 10^9 + 7 取模 。

示例 1：

输入：s = 'bccb'
输出：6
解释：6 个不同的非空回文子字符序列分别为：'b', 'c', 'bb', 'cc', 'bcb', 'bccb'。
注意：'bcb' 虽然出现两次但仅计数一次。
示例 2：

输入：s = 'abcdabcdabcdabcdabcdabcdabcdabcddcbadcbadcbadcbadcbadcbadcbadcba'
输出：104860361
解释：共有 3104860382 个不同的非空回文子序列，104860361 对 109 + 7 取模后的值。

提示：

1 <= s.length <= 1000
s[i] 仅包含 'a', 'b', 'c' 或 'd'
*/
func main() {
	var tests = []struct {
		s    string
		want int
	}{
		{},
	}

	for _, item := range tests {
		if ans := countPalindromicSubsequences(item.s); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func countPalindromicSubsequences(s string) (ans int) {
	const mod int = 1e9 + 7
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	for sz := 2; sz <= n; sz++ {
		for i, j := 0, sz-1; j < n; i++ {
			if s[i] == s[j] {
				low, high := i+1, j-1
				for low <= high && s[low] != s[i] {
					low++
				}
				for high >= low && s[high] != s[j] {
					high--
				}
				if low > high {
					dp[i][j] = (2 + dp[i+1][j-1]*2) % mod
				} else if low == high {
					dp[i][j] = (1 + dp[i+1][j-1]*2) % mod
				} else {
					dp[i][j] = (dp[i+1][j-1]*2 - dp[low+1][high-1] + mod) % mod
				}
			} else {
				dp[i][j] = (dp[i+1][j] + dp[i][j-1] - dp[i+1][j-1] + mod) % mod
			}
			j++
		}
	}

	return dp[0][n-1]
}
