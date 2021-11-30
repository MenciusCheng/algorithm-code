package main

import "fmt"

/*
https://leetcode-cn.com/problems/perfect-rectangle/

难度：困难

391. 完美矩形
给你一个数组 rectangles ，其中 rectangles[i] = [xi, yi, ai, bi] 表示一个坐标轴平行的矩形。这个矩形的左下顶点是 (xi, yi) ，右上顶点是 (ai, bi) 。

如果所有矩形一起精确覆盖了某个矩形区域，则返回 true ；否则，返回 false 。
*/

func main() {
	fmt.Println(isRectangleCover([][]int{{1, 1, 3, 3}, {3, 1, 4, 2}, {3, 2, 4, 4}, {1, 3, 2, 4}, {2, 3, 3, 4}}) == true)
	fmt.Println(isRectangleCover([][]int{{1, 1, 2, 3}, {1, 3, 2, 4}, {3, 1, 4, 2}, {3, 2, 4, 4}}) == false)
}

func isRectangleCover(rectangles [][]int) bool {
	maxRect := rectangles[0]

	cnt := make(map[[2]int]int)
	sumArea := 0

	for _, rect := range rectangles {
		sumArea += calArea(rect)

		maxRect[0] = min(maxRect[0], rect[0])
		maxRect[1] = min(maxRect[1], rect[1])
		maxRect[2] = max(maxRect[2], rect[2])
		maxRect[3] = max(maxRect[3], rect[3])

		cnt[[2]int{rect[0], rect[1]}]++
		cnt[[2]int{rect[0], rect[3]}]++
		cnt[[2]int{rect[2], rect[3]}]++
		cnt[[2]int{rect[2], rect[1]}]++
	}

	if calArea(maxRect) != sumArea {
		return false
	}

	cnt[[2]int{maxRect[0], maxRect[1]}]++
	cnt[[2]int{maxRect[0], maxRect[3]}]++
	cnt[[2]int{maxRect[2], maxRect[3]}]++
	cnt[[2]int{maxRect[2], maxRect[1]}]++

	for _, v := range cnt {
		if v != 2 && v != 4 {
			return false
		}
	}
	return true
}

func calArea(points []int) int {
	return (points[2] - points[0]) * (points[3] - points[1])
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
