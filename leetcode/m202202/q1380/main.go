package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode-cn.com/problems/lucky-numbers-in-a-matrix/

1380. 矩阵中的幸运数
给你一个 m * n 的矩阵，矩阵中的数字 各不相同 。请你按 任意 顺序返回矩阵中的所有幸运数。

幸运数是指矩阵中满足同时下列两个条件的元素：
在同一行的所有元素中最小
在同一列的所有元素中最大

示例 1：

输入：matrix = [[3,7,8],[9,11,13],[15,16,17]]
输出：[15]
解释：15 是唯一的幸运数，因为它是其所在行中的最小值，也是所在列中的最大值。
示例 2：

输入：matrix = [[1,10,4,2],[9,3,8,7],[15,16,17,12]]
输出：[12]
解释：12 是唯一的幸运数，因为它是其所在行中的最小值，也是所在列中的最大值。
示例 3：

输入：matrix = [[7,8],[1,2]]
输出：[7]

提示：

m == mat.length
n == mat[i].length
1 <= n, m <= 50
1 <= matrix[i][j] <= 10^5
矩阵中的所有元素都是不同的
*/
func main() {
	var tests = []struct {
		matrix [][]int
		want   []int
	}{
		{
			matrix: [][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}},
			want:   []int{15},
		},
		{
			matrix: [][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}},
			want:   []int{12},
		},
		{
			matrix: [][]int{{7, 8}, {1, 2}},
			want:   []int{7},
		},
	}

	for _, item := range tests {
		if ans := luckyNumbers(item.matrix); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func luckyNumbers(matrix [][]int) []int {
	rowMin := make([]int, len(matrix))
	colMax := make([]int, len(matrix[0]))
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[0]); c++ {
			v := matrix[r][c]
			if v < matrix[r][rowMin[r]] {
				rowMin[r] = c
			}
			if v > matrix[colMax[c]][c] {
				colMax[c] = r
			}
		}
	}

	res := make([]int, 0)
	cnt := make(map[[2]int]bool)
	for i, v := range rowMin {
		cnt[[2]int{i, v}] = true
	}
	for i, v := range colMax {
		if cnt[[2]int{v, i}] {
			res = append(res, matrix[v][i])
		}
	}

	return res
}
