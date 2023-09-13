package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/check-knight-tour-configuration/description/

2596. 检查骑士巡视方案
中等

骑士在一张 n x n 的棋盘上巡视。在 有效 的巡视方案中，骑士会从棋盘的 左上角 出发，并且访问棋盘上的每个格子 恰好一次 。
给你一个 n x n 的整数矩阵 grid ，由范围 [0, n * n - 1] 内的不同整数组成，其中 grid[row][col] 表示单元格 (row, col) 是骑士访问的第 grid[row][col] 个单元格。骑士的行动是从下标 0 开始的。
如果 grid 表示了骑士的有效巡视方案，返回 true；否则返回 false。
注意，骑士行动时可以垂直移动两个格子且水平移动一个格子，或水平移动两个格子且垂直移动一个格子。下图展示了骑士从某个格子出发可能的八种行动路线。

示例 1：

输入：grid = [[0,11,16,5,20],[17,4,19,10,15],[12,1,8,21,6],[3,18,23,14,9],[24,13,2,7,22]]
输出：true
解释：grid 如上图所示，可以证明这是一个有效的巡视方案。
示例 2：

输入：grid = [[0,3,6],[5,8,1],[2,7,4]]
输出：false
解释：grid 如上图所示，考虑到骑士第 7 次行动后的位置，第 8 次行动是无效的。

提示：

n == grid.length == grid[i].length
3 <= n <= 7
0 <= grid[row][col] < n * n
grid 中的所有整数 互不相同
*/
func main() {
	var tests = []struct {
		grid [][]int
		want bool
	}{
		{
			grid: [][]int{{24, 11, 22, 17, 4}, {21, 16, 5, 12, 9}, {6, 23, 10, 3, 18}, {15, 20, 1, 8, 13}, {0, 7, 14, 19, 2}},
			want: false,
		},
		{
			grid: [][]int{{0, 11, 16, 5, 20}, {17, 4, 19, 10, 15}, {12, 1, 8, 21, 6}, {3, 18, 23, 14, 9}, {24, 13, 2, 7, 22}},
			want: true,
		},
		{
			grid: [][]int{{0, 3, 6}, {5, 8, 1}, {2, 7, 4}},
			want: false,
		},
	}

	for _, item := range tests {
		if ans := checkValidGrid(item.grid); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func checkValidGrid(grid [][]int) bool {
	n := len(grid)
	arr := make([][]int, n*n)
	for i, row := range grid {
		for j, v := range row {
			arr[v] = []int{i, j}
		}
	}

	c := arr[0]
	if c[0] != 0 || c[1] != 0 {
		return false
	}
	for i := 1; i < len(arr); i++ {
		if !canReach(c, arr[i]) {
			return false
		}
		c = arr[i]
	}

	return true
}

func canReach(c []int, t []int) bool {
	return abs(c[0], t[0]) == 2 && abs(c[1], t[1]) == 1 || abs(c[0], t[0]) == 1 && abs(c[1], t[1]) == 2
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
