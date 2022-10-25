package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/shortest-bridge/

934. 最短的桥
中等

给你一个大小为 n x n 的二元矩阵 grid ，其中 1 表示陆地，0 表示水域。
岛 是由四面相连的 1 形成的一个最大组，即不会与非组内的任何其他 1 相连。grid 中 恰好存在两座岛 。
你可以将任意数量的 0 变为 1 ，以使两座岛连接起来，变成 一座岛 。
返回必须翻转的 0 的最小数目。

示例 1：

输入：grid = [[0,1],[1,0]]
输出：1
示例 2：

输入：grid = [[0,1,0],[0,0,0],[0,0,1]]
输出：2
示例 3：

输入：grid = [[1,1,1,1,1],[1,0,0,0,1],[1,0,1,0,1],[1,0,0,0,1],[1,1,1,1,1]]
输出：1

提示：

n == grid.length == grid[i].length
2 <= n <= 100
grid[i][j] 为 0 或 1
grid 中恰有两个岛
*/
func main() {
	var tests = []struct {
		grid [][]int
		want int
	}{
		{
			grid: [][]int{{0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 0, 0}, {1, 1, 0, 0, 0, 0}, {1, 1, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0}, {0, 0, 1, 1, 0, 0}},
			want: 2,
		},
		{
			grid: [][]int{{0, 1}, {1, 0}},
			want: 1,
		},
		{
			grid: [][]int{{0, 1, 0}, {0, 0, 0}, {0, 0, 1}},
			want: 2,
		},
		{
			grid: [][]int{{1, 1, 1, 1, 1}, {1, 0, 0, 0, 1}, {1, 0, 1, 0, 1}, {1, 0, 0, 0, 1}, {1, 1, 1, 1, 1}},
			want: 1,
		},
	}

	for _, item := range tests {
		if ans := shortestBridge(item.grid); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func shortestBridge(grid [][]int) int {
	aMap := make(map[[2]int]bool)
	bMap := make(map[[2]int]bool)
	points := make([][2]int, 0)
	ds := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	var sp []int
	for i, row := range grid {
		for j, cell := range row {
			if cell == 1 {
				bMap[[2]int{i, j}] = true
				if sp == nil {
					sp = []int{i, j}
				}
			}
		}
	}

	var dfs func(x, y int)
	dfs = func(x, y int) {
		delete(bMap, [2]int{x, y})
		aMap[[2]int{x, y}] = true
		points = append(points, [2]int{x, y})

		for _, d := range ds {
			i := x + d[0]
			j := y + d[1]
			if bMap[[2]int{i, j}] {
				dfs(i, j)
			}
		}
	}
	dfs(sp[0], sp[1])

	t := 0
	for len(points) > 0 {
		nextPoints := make([][2]int, 0)
		for _, p := range points {
			for _, d := range ds {
				i := p[0] + d[0]
				j := p[1] + d[1]
				if i >= 0 && i < len(grid) && j >= 0 && j < len(grid) && grid[i][j] == 0 {
					grid[i][j] = 1
					nextPoints = append(nextPoints, [2]int{i, j})
				} else if bMap[[2]int{i, j}] {
					return t
				}
			}
		}
		t++
		points = nextPoints
	}

	return 0
}
