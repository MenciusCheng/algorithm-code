package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode-cn.com/problems/plates-between-candles/

2055. 蜡烛之间的盘子
给你一个长桌子，桌子上盘子和蜡烛排成一列。给你一个下标从 0 开始的字符串 s ，它只包含字符 '*' 和 '|' ，其中 '*' 表示一个 盘子 ，'|' 表示一支 蜡烛 。
同时给你一个下标从 0 开始的二维整数数组 queries ，其中 queries[i] = [lefti, righti] 表示 子字符串 s[lefti...righti] （包含左右端点的字符）。对于每个查询，你需要找到 子字符串中 在 两支蜡烛之间 的盘子的 数目 。如果一个盘子在 子字符串中 左边和右边 都 至少有一支蜡烛，那么这个盘子满足在 两支蜡烛之间 。
比方说，s = "||**||**|*" ，查询 [3, 8] ，表示的是子字符串 "*||**|" 。子字符串中在两支蜡烛之间的盘子数目为 2 ，子字符串中右边两个盘子在它们左边和右边 都 至少有一支蜡烛。
请你返回一个整数数组 answer ，其中 answer[i] 是第 i 个查询的答案。

示例 1:

ex-1

输入：s = "**|**|***|", queries = [[2,5],[5,9]]
输出：[2,3]
解释：
- queries[0] 有两个盘子在蜡烛之间。
- queries[1] 有三个盘子在蜡烛之间。
示例 2:

ex-2

输入：s = "***|**|*****|**||**|*", queries = [[1,17],[4,5],[14,17],[5,11],[15,16]]
输出：[9,0,0,0,0]
解释：
- queries[0] 有 9 个盘子在蜡烛之间。
- 另一个查询没有盘子在蜡烛之间。

提示：

3 <= s.length <= 10^5
s 只包含字符 '*' 和 '|' 。
1 <= queries.length <= 10^5
queries[i].length == 2
0 <= lefti <= righti < s.length
*/
func main() {
	var tests = []struct {
		s       string
		queries [][]int
		want    []int
	}{
		{
			s:       "**|**|***|",
			queries: [][]int{{2, 5}, {5, 9}},
			want:    []int{2, 3},
		},
		{
			s:       "***|**|*****|**||**|*",
			queries: [][]int{{1, 17}, {4, 5}, {14, 17}, {5, 11}, {15, 16}},
			want:    []int{9, 0, 0, 0, 0},
		},
		{
			s:       "||*",
			queries: [][]int{{2, 2}},
			want:    []int{0},
		},
		{
			s:       "***",
			queries: [][]int{{0, 1}},
			want:    []int{0},
		},
	}

	for _, item := range tests {
		if ans := platesBetweenCandles(item.s, item.queries); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func platesBetweenCandles(s string, queries [][]int) []int {
	lCount := make([]int, len(s))
	rCount := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		lCount[i] = -1
		rCount[i] = -1
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '|' {
			lCount[i] = i
		} else if i > 0 {
			lCount[i] = lCount[i-1]
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '|' {
			rCount[i] = i
		} else if i < len(s)-1 {
			rCount[i] = rCount[i+1]
		}
	}

	sCount := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		if i > 0 {
			sCount[i] = sCount[i-1]
		}
		if s[i] == '|' {
			sCount[i]++
		}
	}

	res := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		l, r := queries[i][0], queries[i][1]
		if l == 0 && sCount[r] >= 2 || l > 0 && sCount[r]-sCount[l-1] >= 2 {
			l2 := rCount[l]
			r2 := lCount[r]
			count := r2 - l2 - (sCount[r2] - sCount[l2])
			res[i] = count
		}
	}
	return res
}
