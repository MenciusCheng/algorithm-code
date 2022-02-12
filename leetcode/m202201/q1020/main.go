package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode-cn.com/problems/number-of-enclaves/

1020. 飞地的数量
给你一个大小为 m x n 的二进制矩阵 grid ，其中 0 表示一个海洋单元格、1 表示一个陆地单元格。
一次 移动 是指从一个陆地单元格走到另一个相邻（上、下、左、右）的陆地单元格或跨过 grid 的边界。
返回网格中 无法 在任意次数的移动中离开网格边界的陆地单元格的数量。

示例 1：

输入：grid = [[0,0,0,0],[1,0,1,0],[0,1,1,0],[0,0,0,0]]
输出：3
解释：有三个 1 被 0 包围。一个 1 没有被包围，因为它在边界上。
示例 2：

输入：grid = [[0,1,1,0],[0,0,1,0],[0,0,1,0],[0,0,0,0]]
输出：0
解释：所有 1 都在边界上或可以到达边界。

提示：

m == grid.length
n == grid[i].length
1 <= m, n <= 500
grid[i][j] 的值为 0 或 1
*/
func main() {
	var tests = []struct {
		grid [][]int
		want int
	}{
		{
			grid: [][]int{{0, 0, 0, 0}, {1, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}},
			want: 3,
		},
		{
			grid: [][]int{{0, 1, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}},
			want: 0,
		},
		{
			grid: [][]int{{0, 0, 0, 1, 1, 1, 0, 1, 0, 0}, {1, 1, 0, 0, 0, 1, 0, 1, 1, 1}, {0, 0, 0, 1, 1, 1, 0, 1, 0, 0}, {0, 1, 1, 0, 0, 0, 1, 0, 1, 0}, {0, 1, 1, 1, 1, 1, 0, 0, 1, 0}, {0, 0, 1, 0, 1, 1, 1, 1, 0, 1}, {0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, {0, 0, 1, 0, 0, 1, 0, 1, 0, 1}, {1, 0, 1, 0, 1, 1, 0, 0, 0, 0}, {0, 0, 0, 0, 1, 1, 0, 0, 0, 1}},
			want: 3,
		},
		{
			grid: [][]int{{0}, {1}, {1}, {0}, {0}},
			want: 0,
		},
	}

	for _, item := range tests {
		if ans := numEnclaves(item.grid); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func numEnclaves(grid [][]int) int {
	for c := 0; c < len(grid[0]); c++ {
		if grid[0][c] == 1 {
			bfs([2]int{0, c}, grid)
		}
		if grid[len(grid)-1][c] == 1 {
			bfs([2]int{len(grid) - 1, c}, grid)
		}
	}
	for r := 0; r < len(grid); r++ {
		if grid[r][0] == 1 {
			bfs([2]int{r, 0}, grid)
		}
		if grid[r][len(grid[0])-1] == 1 {
			bfs([2]int{r, len(grid[0]) - 1}, grid)
		}
	}

	count := 0
	for r := 1; r < len(grid)-1; r++ {
		for c := 1; c < len(grid[0])-1; c++ {
			if grid[r][c] == 1 {
				count++
			}
		}
	}
	return count
}

var direction = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func bfs(p [2]int, grid [][]int) {
	next := make(map[[2]int]bool)
	next[p] = true
	for len(next) > 0 {
		newNext := make(map[[2]int]bool)
		for np := range next {
			grid[np[0]][np[1]] = 0
			for _, d := range direction {
				r := np[0] + d[0]
				c := np[1] + d[1]
				if r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0]) && grid[r][c] == 1 && !newNext[[2]int{r, c}] {
					newNext[[2]int{r, c}] = true
				}
			}
		}
		next = newNext
	}
}
