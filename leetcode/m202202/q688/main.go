package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode-cn.com/problems/knight-probability-in-chessboard/

688. 骑士在棋盘上的概率
在一个 n x n 的国际象棋棋盘上，一个骑士从单元格 (row, column) 开始，并尝试进行 k 次移动。行和列是 从 0 开始 的，所以左上单元格是 (0,0) ，右下单元格是 (n - 1, n - 1) 。
象棋骑士有8种可能的走法，如下图所示。每次移动在基本方向上是两个单元格，然后在正交方向上是一个单元格。
每次骑士要移动时，它都会随机从8种可能的移动中选择一种(即使棋子会离开棋盘)，然后移动到那里。
骑士继续移动，直到它走了 k 步或离开了棋盘。
返回 骑士在棋盘停止移动后仍留在棋盘上的概率 。

示例 1：

输入: n = 3, k = 2, row = 0, column = 0
输出: 0.0625
解释: 有两步(到(1,2)，(2,1))可以让骑士留在棋盘上。
在每一个位置上，也有两种移动可以让骑士留在棋盘上。
骑士留在棋盘上的总概率是0.0625。
示例 2：

输入: n = 1, k = 0, row = 0, column = 0
输出: 1.00000

提示:

1 <= n <= 25
0 <= k <= 100
0 <= row, column <= n
*/
func main() {
	var tests = []struct {
		n      int
		k      int
		row    int
		column int
		want   float64
	}{
		{
			n:      3,
			k:      2,
			row:    0,
			column: 0,
			want:   0.0625,
		},
		//{
		//	n:      1,
		//	k:      0,
		//	row:    0,
		//	column: 0,
		//	want:   1,
		//},
		//{
		//	n:      8,
		//	k:      30,
		//	row:    6,
		//	column: 4,
		//	want:   1,
		//},
	}

	for _, item := range tests {
		if ans := knightProbability(item.n, item.k, item.row, item.column); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func knightProbability2(n int, k int, row int, column int) float64 {
	ds := [][2]int{
		{-1, -2}, {-2, -1}, {1, -2}, {2, -1}, {-1, 2}, {-2, 1}, {1, 2}, {2, 1},
	}
	next := map[[2]int]int{
		[2]int{row, column}: 1,
	}

	out := 0
	for i := 0; i < k; i++ {
		if len(next) == 0 {
			break
		}
		newNext := make(map[[2]int]int)
		for ch, v := range next {
			for _, d := range ds {
				r1 := ch[0] + d[0]
				c1 := ch[1] + d[1]
				if r1 >= 0 && r1 < n && c1 >= 0 && c1 < n {
					newNext[[2]int{r1, c1}] += v
				} else {
					out += v
				}
			}
		}
		next = newNext
	}

	if len(next) == 0 {
		return 0
	}
	sum := 0
	for _, v := range next {
		sum += v
	}
	return float64(sum) / float64(out+sum)
}

var dirs = []struct{ i, j int }{{-2, -1}, {-2, 1}, {2, -1}, {2, 1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}}

func knightProbability(n, k, row, column int) float64 {
	dp := make([][][]float64, k+1)
	for step := range dp {
		dp[step] = make([][]float64, n)
		for i := 0; i < n; i++ {
			dp[step][i] = make([]float64, n)
			for j := 0; j < n; j++ {
				if step == 0 {
					dp[step][i][j] = 1
				} else {
					for _, d := range dirs {
						if x, y := i+d.i, j+d.j; 0 <= x && x < n && 0 <= y && y < n {
							dp[step][i][j] += dp[step-1][x][y] / 8
						}
					}
				}
			}
		}
	}
	return dp[k][row][column]
}
