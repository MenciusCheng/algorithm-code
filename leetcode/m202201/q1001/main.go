package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode-cn.com/problems/grid-illumination/

1001. 网格照明
在大小为 n x n 的网格 grid 上，每个单元格都有一盏灯，最初灯都处于 关闭 状态。
给你一个由灯的位置组成的二维数组 lamps ，其中 lamps[i] = [rowi, coli] 表示 打开 位于 grid[rowi][coli] 的灯。即便同一盏灯可能在 lamps 中多次列出，不会影响这盏灯处于 打开 状态。
当一盏灯处于打开状态，它将会照亮 自身所在单元格 以及同一 行 、同一 列 和两条 对角线 上的 所有其他单元格 。
另给你一个二维数组 queries ，其中 queries[j] = [rowj, colj] 。对于第 j 个查询，如果单元格 [rowj, colj] 是被照亮的，则查询结果为 1 ，否则为 0 。在第 j 次查询之后 [按照查询的顺序] ，关闭 位于单元格 grid[rowj][colj] 上及相邻 8 个方向上（与单元格 grid[rowi][coli] 共享角或边）的任何灯。
返回一个整数数组 ans 作为答案， ans[j] 应等于第 j 次查询 queries[j] 的结果，1 表示照亮，0 表示未照亮。

示例 1：

输入：n = 5, lamps = [[0,0],[4,4]], queries = [[1,1],[1,0]]
输出：[1,0]
解释：最初所有灯都是关闭的。在执行查询之前，打开位于 [0, 0] 和 [4, 4] 的灯。第 0 次查询检查 grid[1][1] 是否被照亮（蓝色方框）。该单元格被照亮，所以 ans[0] = 1 。然后，关闭红色方框中的所有灯。

第 1 次查询检查 grid[1][0] 是否被照亮（蓝色方框）。该单元格没有被照亮，所以 ans[1] = 0 。然后，关闭红色矩形中的所有灯。

示例 2：

输入：n = 5, lamps = [[0,0],[4,4]], queries = [[1,1],[1,1]]
输出：[1,1]
示例 3：

输入：n = 5, lamps = [[0,0],[0,4]], queries = [[0,4],[0,1],[1,4]]
输出：[1,1,0]

提示：

1 <= n <= 10^9
0 <= lamps.length <= 20000
0 <= queries.length <= 20000
lamps[i].length == 2
0 <= rowi, coli < n
queries[j].length == 2
0 <= rowj, colj < n
*/
func main() {
	var tests = []struct {
		n       int
		lamps   [][]int
		queries [][]int
		want    []int
	}{
		{
			n:       5,
			lamps:   [][]int{{0, 0}, {4, 4}},
			queries: [][]int{{1, 1}, {1, 0}},
			want:    []int{1, 0},
		},
		{
			n:       5,
			lamps:   [][]int{{0, 0}, {4, 4}},
			queries: [][]int{{1, 1}, {1, 1}},
			want:    []int{1, 1},
		},
		{
			n:       5,
			lamps:   [][]int{{0, 0}, {0, 4}},
			queries: [][]int{{0, 4}, {0, 1}, {1, 4}},
			want:    []int{1, 1, 0},
		},
		{
			n:       6,
			lamps:   [][]int{{2, 5}, {4, 2}, {0, 3}, {0, 5}, {1, 4}, {4, 2}, {3, 3}, {1, 0}},
			queries: [][]int{{4, 3}, {3, 1}, {5, 3}, {0, 5}, {4, 4}, {3, 3}},
			want:    []int{1, 0, 1, 1, 0, 1},
		},
	}

	for _, item := range tests {
		if ans := gridIllumination(item.n, item.lamps, item.queries); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func gridIllumination(n int, lamps [][]int, queries [][]int) []int {
	xym := make(map[[2]int]bool)
	xm := make(map[int]int)
	ym := make(map[int]int)
	xay := make(map[int]int) // x + y
	xby := make(map[int]int) // x - y

	ds := [][2]int{{0, 0}, {-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	var x, y int

	for i := 0; i < len(lamps); i++ {
		x = lamps[i][0]
		y = lamps[i][1]
		if !xym[[2]int{x, y}] {
			xym[[2]int{x, y}] = true
			xm[x]++
			ym[y]++
			xay[x+y]++
			xby[x-y]++
		}
	}

	res := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		x = queries[i][0]
		y = queries[i][1]
		if xm[x] > 0 || ym[y] > 0 || xay[x+y] > 0 || xby[x-y] > 0 {
			res[i] = 1
		}
		for _, d := range ds {
			x1 := x + d[0]
			y1 := y + d[1]
			if x1 >= 0 && x1 < n && y1 >= 0 && y1 < n && xym[[2]int{x1, y1}] {
				xym[[2]int{x1, y1}] = false
				xm[x1]--
				ym[y1]--
				xay[x1+y1]--
				xby[x1-y1]--
			}
		}
	}

	return res
}
