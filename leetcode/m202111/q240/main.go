package main

import "fmt"

// https://leetcode-cn.com/problems/search-a-2d-matrix-ii/
// 240. 搜索二维矩阵 II
func main() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	fmt.Println(searchMatrixZ(matrix, 5))
}

// N * 二分法
func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		if binarySearch(matrix[i], target) {
			return true
		}
	}
	return false
}

func binarySearch(arr []int, target int) bool {
	head, tail := 0, len(arr)-1
	for head <= tail {
		i := (head + tail) / 2

		if arr[i] == target {
			return true
		} else if arr[i] < target {
			head = i + 1
		} else {
			tail = i - 1
		}
	}
	return false
}

func searchMatrixZ(matrix [][]int, target int) bool {
	x, y := 0, len(matrix[0])-1

	for x < len(matrix) && y > 0 {
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] < target {
			x++
		} else {
			y--
		}
	}

	return false
}
