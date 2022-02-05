package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode-cn.com/problems/path-with-maximum-gold/

1219. 黄金矿工
你要开发一座金矿，地质勘测学家已经探明了这座金矿中的资源分布，并用大小为 m * n 的网格 grid 进行了标注。每个单元格中的整数就表示这一单元格中的黄金数量；如果该单元格是空的，那么就是 0。
为了使收益最大化，矿工需要按以下规则来开采黄金：

每当矿工进入一个单元，就会收集该单元格中的所有黄金。
矿工每次可以从当前位置向上下左右四个方向走。
每个单元格只能被开采（进入）一次。
不得开采（进入）黄金数目为 0 的单元格。
矿工可以从网格中 任意一个 有黄金的单元格出发或者是停止。

示例 1：

输入：grid = [[0,6,0],[5,8,7],[0,9,0]]
输出：24
解释：
[[0,6,0],
 [5,8,7],
 [0,9,0]]
一种收集最多黄金的路线是：9 -> 8 -> 7。
示例 2：

输入：grid = [[1,0,7],[2,0,6],[3,4,5],[0,3,0],[9,0,20]]
输出：28
解释：
[[1,0,7],
 [2,0,6],
 [3,4,5],
 [0,3,0],
 [9,0,20]]
一种收集最多黄金的路线是：1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7。

提示：

1 <= grid.length, grid[i].length <= 15
0 <= grid[i][j] <= 100
最多 25 个单元格中有黄金。
*/
func main() {
	var tests = []struct {
		grid [][]int
		want int
	}{
		{
			grid: [][]int{{0, 6, 0}, {5, 8, 7}, {0, 9, 0}},
			want: 24,
		},
		{
			grid: [][]int{{1, 0, 7}, {2, 0, 6}, {3, 4, 5}, {0, 3, 0}, {9, 0, 20}},
			want: 28,
		},
	}

	for _, item := range tests {
		if ans := getMaximumGold(item.grid); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func getMaximumGold(grid [][]int) int {
	arr := make([]map[[2]int]int, 0)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != 0 {
				arr = append(arr, through(grid, [2]int{i, j}))
			}
		}
	}

	sum := 0
	for _, m := range arr {
		sum += maxM(m)
	}
	return sum
}

func maxM(m map[[2]int]int) int {
	starts := make([][2]int, 0)
	for k := range m {
		count := 0
		if m[[2]int{k[0] - 1, k[1]}] > 0 {
			count++
		}
		if m[[2]int{k[0] + 1, k[1]}] > 0 {
			count++
		}
		if m[[2]int{k[0], k[1] - 1}] > 0 {
			count++
		}
		if m[[2]int{k[0], k[1] + 1}] > 0 {
			count++
		}
		if count == 1 {
			starts = append(starts, k)
		}
	}

	max := 0
	for _, start := range starts {
		visited := make(map[[2]int]bool)
		ns := maxT(m, start, visited, 0)
		if ns > max {
			max = ns
		}
	}
	return max
}

func maxT(m map[[2]int]int, p [2]int, visited map[[2]int]bool, s int) int {
	v := make(map[[2]int]bool)
	for ints := range visited {
		v[ints] = true
	}
	v[p] = true
	s += m[p]
	max := 0

	nextP := [2]int{p[0] - 1, p[1]}
	if m[nextP] > 0 && !v[nextP] {
		ns := maxT(m, nextP, v, s)
		if ns > max {
			max = ns
		}
	}
	nextP = [2]int{p[0] + 1, p[1]}
	if m[nextP] > 0 && !v[nextP] {
		ns := maxT(m, nextP, v, s)
		if ns > max {
			max = ns
		}
	}
	nextP = [2]int{p[0], p[1] - 1}
	if m[nextP] > 0 && !v[nextP] {
		ns := maxT(m, nextP, v, s)
		if ns > max {
			max = ns
		}
	}
	nextP = [2]int{p[0], p[1] + 1}
	if m[nextP] > 0 && !v[nextP] {
		ns := maxT(m, nextP, v, s)
		if ns > max {
			max = ns
		}
	}
	return max
}

func through(grid [][]int, p [2]int) map[[2]int]int {
	m := make(map[[2]int]int)

	next := [][2]int{p}
	for len(next) > 0 {
		newNext := make([][2]int, 0)
		for _, point := range next {
			m[point] = grid[point[0]][point[1]]
			grid[point[0]][point[1]] = 0

			x := point[0] - 1
			y := point[1]
			if x >= 0 && grid[x][y] > 0 {
				newNext = append(newNext, [2]int{x, y})
			}
			x = point[0] + 1
			if x < len(grid) && grid[x][y] > 0 {
				newNext = append(newNext, [2]int{x, y})
			}
			x = point[0]
			y = point[1] - 1
			if y >= 0 && grid[x][y] > 0 {
				newNext = append(newNext, [2]int{x, y})
			}
			y = point[1] + 1
			if y < len(grid[0]) && grid[x][y] > 0 {
				newNext = append(newNext, [2]int{x, y})
			}
		}
		next = newNext
	}
	return m
}
