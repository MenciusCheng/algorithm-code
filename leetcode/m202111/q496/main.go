package main

import "fmt"

// https://leetcode-cn.com/problems/next-greater-element-i/
// 496. 下一个更大元素 I
func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	fmt.Println(nextGreaterElement(nums1, nums2))
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	mp := make(map[int]int, len(nums1))
	stack := make([]int, 0)

	for i := len(nums2) - 1; i >= 0; i-- {
		num := nums2[i]
		for len(stack) > 0 && num >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			mp[num] = stack[len(stack)-1]
		} else {
			mp[num] = -1
		}
		stack = append(stack, num)
	}

	res := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		res[i] = mp[nums1[i]]
	}
	return res
}
