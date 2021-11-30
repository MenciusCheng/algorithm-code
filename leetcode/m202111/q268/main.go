package main

import "fmt"

/*
268. 丢失的数字
给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。

https://leetcode-cn.com/problems/missing-number/
*/

func main() {
	fmt.Println(missingNumber([]int{3, 0, 1}))
}

func missingNumber(nums []int) int {
	length := len(nums)
	sum := (1 + length) * length / 2
	for _, num := range nums {
		sum -= num
	}
	return sum
}
