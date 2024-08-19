package main

import (
	"fmt"
	"math"
	"reflect"
)

/*
https://leetcode.cn/problems/maximum-difference-score-in-a-grid/description/?envType=daily-question&envId=2024-08-15

3148. 矩阵中的最大得分
中等
相关标签
相关企业
提示
给你一个由 正整数 组成、大小为 m x n 的矩阵 grid。你可以从矩阵中的任一单元格移动到另一个位于正下方或正右侧的任意单元格（不必相邻）。从值为 c1 的单元格移动到值为 c2 的单元格的得分为 c2 - c1 。
你可以从 任一 单元格开始，并且必须至少移动一次。
返回你能得到的 最大 总得分。

示例 1：
输入：grid = [[9,5,7,3],[8,9,6,1],[6,7,14,3],[2,5,3,1]]

输出：9

解释：从单元格 (0, 1) 开始，并执行以下移动：
- 从单元格 (0, 1) 移动到 (2, 1)，得分为 7 - 5 = 2 。
- 从单元格 (2, 1) 移动到 (2, 2)，得分为 14 - 7 = 7 。
总得分为 2 + 7 = 9 。

示例 2：
输入：grid = [[4,3,2],[3,2,1]]
输出：-1

解释：从单元格 (0, 0) 开始，执行一次移动：从 (0, 0) 到 (0, 1) 。得分为 3 - 4 = -1 。

提示：
m == grid.length
n == grid[i].length
2 <= m, n <= 1000
4 <= m * n <= 10^5
1 <= grid[i][j] <= 10^5
*/
func main() {
	var tests = []struct {
		grid [][]int
		want int
	}{
		{
			grid: [][]int{{9, 5, 7, 3}, {8, 9, 6, 1}, {6, 7, 14, 3}, {2, 5, 3, 1}},
			want: 9,
		},
		{
			grid: [][]int{{4, 3, 2}, {3, 2, 1}},
			want: -1,
		},
	}

	for _, item := range tests {
		if ans := maxScore(item.grid); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func maxScore(grid [][]int) int {
	ans := math.MinInt
	m, n := len(grid), len(grid[0])
	f := make([][]int, m+1)
	f[0] = make([]int, n+1)
	for j := range f[0] {
		f[0][j] = math.MaxInt
	}
	for i, row := range grid {
		f[i+1] = make([]int, n+1)
		f[i+1][0] = math.MaxInt
		for j, x := range row {
			mn := min(f[i+1][j], f[i][j+1])
			ans = max(ans, x-mn)
			f[i+1][j+1] = min(mn, x)
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
