package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/largest-1-bordered-square/

1139. 最大的以 1 为边界的正方形
中等

给你一个由若干 0 和 1 组成的二维网格 grid，请你找出边界全部由 1 组成的最大 正方形 子网格，并返回该子网格中的元素数量。如果不存在，则返回 0。

示例 1：

输入：grid = [[1,1,1],[1,0,1],[1,1,1]]
输出：9
示例 2：

输入：grid = [[1,1,0,0]]
输出：1

提示：

1 <= grid.length <= 100
1 <= grid[0].length <= 100
grid[i][j] 为 0 或 1
*/
func main() {
	var tests = []struct {
		grid [][]int
		want int
	}{
		{
			grid: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			want: 9,
		},
		{
			grid: [][]int{{1, 1, 0, 0}},
			want: 1,
		},
		{
			grid: [][]int{{1, 1, 1}, {0, 1, 0}, {0, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			want: 9,
		},
	}

	for _, item := range tests {
		if ans := largest1BorderedSquare(item.grid); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func largest1BorderedSquare(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	sums := make([][]int, m+1)
	sums[0] = make([]int, n+1)
	for i := 1; i < len(sums); i++ {
		sums[i] = make([]int, n+1)
		for j := 1; j < len(sums[i]); j++ {
			sums[i][j] = grid[i-1][j-1] + sums[i-1][j] + sums[i][j-1] - sums[i-1][j-1]
		}
	}

	res := 0
	for i := range grid {
		for j, col := range grid[i] {
			if col == 0 {
				continue
			} else {
				res = max(res, 1)
			}
			i2 := i + 1
			j2 := j + 1
			for i2 < m && j2 < n && grid[i][j2] == 1 && grid[i2][j] == 1 {
				if sums[i2+1][j2+1]-sums[i2][j2]-(sums[i][j2+1]-sums[i][j2])-(sums[i2+1][j]-sums[i2][j]) == i2-i+j2-j+1 {
					res = max(res, (i2-i+1)*(j2-j+1))
				}
				i2++
				j2++
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
