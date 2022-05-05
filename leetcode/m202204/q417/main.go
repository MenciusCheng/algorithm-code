package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode-cn.com/problems/pacific-atlantic-water-flow/

417. 太平洋大西洋水流问题
有一个 m × n 的矩形岛屿，与 太平洋 和 大西洋 相邻。 “太平洋” 处于大陆的左边界和上边界，而 “大西洋” 处于大陆的右边界和下边界。
这个岛被分割成一个由若干方形单元格组成的网格。给定一个 m x n 的整数矩阵 heights ， heights[r][c] 表示坐标 (r, c) 上单元格 高于海平面的高度 。
岛上雨水较多，如果相邻单元格的高度 小于或等于 当前单元格的高度，雨水可以直接向北、南、东、西流向相邻单元格。水可以从海洋附近的任何单元格流入海洋。
返回 网格坐标 result 的 2D列表 ，其中 result[i] = [ri, ci] 表示雨水可以从单元格 (ri, ci) 流向 太平洋和大西洋 。

示例 1：

输入: heights = [[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]
输出: [[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]]
示例 2：

输入: heights = [[2,1],[1,2]]
输出: [[0,0],[0,1],[1,0],[1,1]]

提示：

m == heights.length
n == heights[r].length
1 <= m, n <= 200
0 <= heights[r][c] <= 10^5
*/
func main() {
	var tests = []struct {
		heights [][]int
		want    [][]int
	}{
		{
			heights: [][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}},
			want:    [][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}},
		},
		{
			heights: [][]int{{2, 1}, {1, 2}},
			want:    [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
		},
	}

	for _, item := range tests {
		if ans := pacificAtlantic(item.heights); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

var m, n int
var directions = [][2]int{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}

func pacificAtlantic(heights [][]int) [][]int {
	m = len(heights)
	n = len(heights[0])

	res := make([][]int, 0)
	resMap := make(map[[2]int]bool)
	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights[0]); j++ {
			visited := make(map[[2]int]bool)
			o1, o2 := cal(i, j, heights, resMap, visited)
			if o1 && o2 {
				res = append(res, []int{i, j})
			}
		}
	}

	return res
}

func cal(i, j int, heights [][]int, resMap map[[2]int]bool, visited map[[2]int]bool) (o1, o2 bool) {
	if i == 0 || j == 0 {
		o1 = true
	}
	if i == m-1 || j == n-1 {
		o2 = true
	}
	if o1 && o2 {
		resMap[[2]int{i, j}] = true
		return
	}

	h := heights[i][j]
	for _, d := range directions {
		x := i + d[0]
		y := j + d[1]
		p := [2]int{x, y}
		if x >= 0 && x < m && y >= 0 && y < n && heights[x][y] <= h && !visited[p] {
			visited[p] = true
			po1, po2 := cal(x, y, heights, resMap, visited)
			o1 = o1 || po1
			o2 = o2 || po2
			if o1 && o2 {
				resMap[[2]int{i, j}] = true
				return
			}
		}
	}
	return
}
