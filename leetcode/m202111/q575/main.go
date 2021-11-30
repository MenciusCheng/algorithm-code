package main

import "fmt"

// https://leetcode-cn.com/problems/distribute-candies/
// 575. 分糖果
func main() {
	fmt.Println(distributeCandies([]int{1, 1, 2, 2, 3, 3}))
}

func distributeCandies(candyType []int) int {
	count := len(candyType) / 2
	m := make(map[int]bool)
	for _, item := range candyType {
		if _, ok := m[item]; !ok {
			m[item] = ok
			count -= 1
			if count == 0 {
				break
			}
		}
	}
	return len(m)
}
