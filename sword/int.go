package sword

import (
	"fmt"
)

// 二进制位运算示例
func binary() {
	a := 60 // 60 = 00111100
	fmt.Printf("a = %d, %08b\n", a, a)
	b := 13 // 13 = 00001101
	fmt.Printf("b = %d, %08b\n", b, b)

	c := a & b // 12 = 00001100
	fmt.Printf("a & b = %d, %08b\n", c, c)

	c = a | b // 61 = 00111101
	fmt.Printf("a | b = %d, %08b\n", c, c)

	c = a ^ b // 49 = 00110001
	fmt.Printf("a ^ b = %d, %08b\n", c, c)

	c = a << 2 // 240 = 11110000
	fmt.Printf("a << 2 = %d, %08b\n", c, c)

	c = a >> 2 // 15 = 00001111
	fmt.Printf("a >> 2 = %d, %08b\n", c, c)
}

/*
### 面试题6：排序数组中的两个数字之和

> 题目：输入一个递增排序的数组和一个值k，请问如何在数组中找出两个和为k的数字并返回它们的下标？假设数组中存在且只存在一对符合条件的数字，同时一个数字不能使用两次。例如，输入数组[1，2，4，6，10]，k的值为8，数组中的数字2与6的和为8，它们的下标分别为1与3。
*/
func twoSum(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1

	for i < j && numbers[i]+numbers[j] != target {
		if numbers[i]+numbers[j] < target {
			i++
		} else {
			j--
		}
	}
	return []int{i, j}
}
