package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode-cn.com/problems/map-of-highest-peak/

1765. 地图中的最高点
给你一个大小为 m x n 的整数矩阵 isWater ，它代表了一个由 陆地 和 水域 单元格组成的地图。

如果 isWater[i][j] == 0 ，格子 (i, j) 是一个 陆地 格子。
如果 isWater[i][j] == 1 ，格子 (i, j) 是一个 水域 格子。
你需要按照如下规则给每个单元格安排高度：

每个格子的高度都必须是非负的。
如果一个格子是是 水域 ，那么它的高度必须为 0 。
任意相邻的格子高度差 至多 为 1 。当两个格子在正东、南、西、北方向上相互紧挨着，就称它们为相邻的格子。（也就是说它们有一条公共边）
找到一种安排高度的方案，使得矩阵中的最高高度值 最大 。

请你返回一个大小为 m x n 的整数矩阵 height ，其中 height[i][j] 是格子 (i, j) 的高度。如果有多种解法，请返回 任意一个 。

示例 1：

输入：isWater = [[0,1],[0,0]]
输出：[[1,0],[2,1]]
解释：上图展示了给各个格子安排的高度。
蓝色格子是水域格，绿色格子是陆地格。
示例 2：

输入：isWater = [[0,0,1],[1,0,0],[0,0,0]]
输出：[[1,1,0],[0,1,1],[1,2,2]]
解释：所有安排方案中，最高可行高度为 2 。
任意安排方案中，只要最高高度为 2 且符合上述规则的，都为可行方案。

提示：

m == isWater.length
n == isWater[i].length
1 <= m, n <= 1000
isWater[i][j] 要么是 0 ，要么是 1 。
至少有 1 个水域格子。
*/
func main() {
	var tests = []struct {
		isWater [][]int
		want    [][]int
	}{
		{
			isWater: [][]int{{0, 1}, {0, 0}},
			want:    [][]int{{1, 0}, {2, 1}},
		},
		{
			isWater: [][]int{{0, 0, 1}, {1, 0, 0}, {0, 0, 0}},
			want:    [][]int{{1, 1, 0}, {0, 1, 1}, {1, 2, 2}},
		},
	}

	for _, item := range tests {
		if ans := highestPeak(item.isWater); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func highestPeak(isWater [][]int) [][]int {
	queue := make([][2]int, 0)

	height := make([][]int, len(isWater))
	for i := range isWater {
		height[i] = make([]int, len(isWater[i]))
		for j := range isWater[i] {
			if isWater[i][j] == 1 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	h := 1
	for len(queue) > 0 {
		nq := make([][2]int, 0)
		for _, point := range queue {
			x, y := point[0]-1, point[1]
			if x >= 0 && isWater[x][y] == 0 && height[x][y] == 0 {
				height[x][y] = h
				nq = append(nq, [2]int{x, y})
			}
			x, y = point[0]+1, point[1]
			if x < len(isWater) && isWater[x][y] == 0 && height[x][y] == 0 {
				height[x][y] = h
				nq = append(nq, [2]int{x, y})
			}
			x, y = point[0], point[1]-1
			if y >= 0 && isWater[x][y] == 0 && height[x][y] == 0 {
				height[x][y] = h
				nq = append(nq, [2]int{x, y})
			}
			x, y = point[0], point[1]+1
			if y < len(isWater[0]) && isWater[x][y] == 0 && height[x][y] == 0 {
				height[x][y] = h
				nq = append(nq, [2]int{x, y})
			}
		}
		h++
		queue = nq
	}

	return height
}
