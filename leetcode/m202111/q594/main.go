package main

import "fmt"

/*
https://leetcode-cn.com/problems/longest-harmonious-subsequence/

594. 最长和谐子序列
和谐数组是指一个数组里元素的最大值和最小值之间的差别 正好是 1 。

现在，给你一个整数数组 nums ，请你在所有可能的子序列中找到最长的和谐子序列的长度。

数组的子序列是一个由数组派生出来的序列，它可以通过删除一些元素或不删除元素、且不改变其余元素的顺序而得到。
*/
func main() {
	fmt.Println(findLHS([]int{1, 3, 2, 2, 5, 2, 3, 7}) == 5)
	fmt.Println(findLHS([]int{1, 2, 3, 4}) == 2)
	fmt.Println(findLHS([]int{1, 1, 1, 1}) == 0)
}

func findLHS(nums []int) int {
	cnt := make(map[int]int)
	for _, num := range nums {
		cnt[num]++
	}

	var max int
	for k, v := range cnt {
		if v+cnt[k+1] > max && cnt[k+1] > 0 {
			max = v + cnt[k+1]
		}
	}
	return max
}
