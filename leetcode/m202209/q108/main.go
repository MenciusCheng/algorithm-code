package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/zero-matrix-lcci/

108
面试题 01.08. 零矩阵
编写一种算法，若M × N矩阵中某个元素为0，则将其所在的行与列清零。

示例 1：

输入：
[
  [1,1,1],
  [1,0,1],
  [1,1,1]
]
输出：
[
  [1,0,1],
  [0,0,0],
  [1,0,1]
]
示例 2：

输入：
[
  [0,1,2,0],
  [3,4,5,2],
  [1,3,1,5]
]
输出：
[
  [0,0,0,0],
  [0,4,5,0],
  [0,3,1,0]
]
*/
func main() {
	var tests = []struct {
		matrix [][]int
		want   [][]int
	}{
		{
			matrix: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			want: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
		{
			matrix: [][]int{
				{0, 1, 2, 0},
				{3, 4, 5, 2},
				{1, 3, 1, 5},
			},
			want: [][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 3, 1, 0},
			},
		},
	}

	for _, item := range tests {
		setZeroes(item.matrix)
		if reflect.DeepEqual(item.matrix, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", item.matrix, item.want)
		}
	}
}

func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	mZero := make([]bool, m)
	nZero := make([]bool, n)

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				mZero[i] = true
				nZero[j] = true
			}
		}
	}

	for i, item := range mZero {
		if item {
			for j := 0; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}
	for i, item := range nZero {
		if item {
			for j := 0; j < m; j++ {
				matrix[j][i] = 0
			}
		}
	}
}
