package main

import "fmt"

// https://leetcode-cn.com/problems/single-number-iii/
// 260. 只出现一次的数字 III
func main() {
	fmt.Println(singleNumber([]int{1, 2, 1, 3, 2, 5}))
}

func singleNumber(nums []int) []int {
	m := make(map[int]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			delete(m, nums[i])
		} else {
			m[nums[i]] = true
		}
	}

	res := make([]int, 0, len(m))
	for k := range m {
		res = append(res, k)
	}
	return res
}
