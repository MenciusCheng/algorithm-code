package main

import "fmt"

// https://leetcode-cn.com/problems/check-if-array-is-sorted-and-rotated/
// 1752. 检查数组是否经排序和轮转得到
func main() {
	fmt.Println(check([]int{3,4,5,1,2}))
	fmt.Println(check([]int{2,1,3,4}))
	fmt.Println(check([]int{1,2,3}))
	fmt.Println(check([]int{1,1,1}))
	fmt.Println(check([]int{2,1}))
}

func check(nums []int) bool {
	x := -1
	for i := len(nums)-1; i > 0; i-- {
		if nums[i] < nums[i-1] {
			if x == -1 {
				x = len(nums) - i
			} else {
				return false
			}
		}
	}
	if x != -1 && nums[0] < nums[len(nums)-1] {
		return false
	}
	return true
}
