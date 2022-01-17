package sword

import (
	"fmt"
	"math"
)

// 二进制
func binary() {
	a := 2
	b := a << 1
	fmt.Println(b)
	c := b & 1
	fmt.Println(c)
	c = b | 1
	fmt.Println(c)
	c = b ^ 1
	fmt.Println(c)

	fmt.Println(math.MaxInt32)
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
