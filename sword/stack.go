package sword

// 栈

/**
https://leetcode-cn.com/problems/largest-rectangle-in-histogram/

面试题39：直方图最大矩形面积

题目：直方图是由排列在同一基线上的相邻柱子组成的图形。输入一个由非负数组成的数组，数组中的数字是直方图中柱子的高。求直方图中最大矩形面积。假设直方图中柱子的宽都为1。例如，输入数组[3，2，5，4，6，1，4，2]，其对应的直方图如图6.3所示，该直方图中最大矩形面积为12，如阴影部分所示。
*/
func largestRectangleArea(heights []int) int {
	stack := make([]int, 0)
	stack = append(stack, -1)
	max := 0
	for i := 0; i < len(heights); i++ {
		for stack[len(stack)-1] != -1 && heights[stack[len(stack)-1]] >= heights[i] {
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			max = calMaxArea(heights, mid, stack[len(stack)-1], i, max)
		}
		stack = append(stack, i)
	}

	for stack[len(stack)-1] != -1 {
		mid := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		max = calMaxArea(heights, mid, stack[len(stack)-1], len(heights), max)
	}
	return max
}

func calMaxArea(heights []int, mid, left, right int, max int) int {
	width := (mid - left - 1) + 1 + (right - mid - 1)
	area := width * heights[mid]
	if area > max {
		return area
	}
	return max
}
