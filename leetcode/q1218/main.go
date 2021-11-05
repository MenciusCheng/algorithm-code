package main

import "fmt"

/*
1218. 最长定差子序列
难度：中等
给你一个整数数组 arr 和一个整数 difference，请你找出并返回 arr 中最长等差子序列的长度，该子序列中相邻元素之间的差等于 difference 。

子序列 是指在不改变其余元素顺序的情况下，通过删除一些元素或不删除任何元素而从 arr 派生出来的序列。

示例 1：

输入：arr = [1,2,3,4], difference = 1
输出：4
解释：最长的等差子序列是 [1,2,3,4]。
示例 2：

输入：arr = [1,3,5,7], difference = 1
输出：1
解释：最长的等差子序列是任意单个元素。
示例 3：

输入：arr = [1,5,7,8,5,3,4,2,1], difference = -2
输出：4
解释：最长的等差子序列是 [7,5,3,1]。

提示：

1 <= arr.length <= 105
-104 <= arr[i], difference <= 104

https://leetcode-cn.com/problems/longest-arithmetic-subsequence-of-given-difference/
*/

func main() {
	fmt.Println(longestSubsequence([]int{1, 2, 3, 4}, 1))
	fmt.Println(longestSubsequence([]int{1, 3, 5, 7}, 1))
	fmt.Println(longestSubsequence([]int{1, 5, 7, 8, 5, 3, 4, 2, 1}, -2))
	fmt.Println(longestSubsequence([]int{7, 8, 9, 9, 9}, 1))
}

// 官方答案：动态规划
func longestSubsequence2(arr []int, difference int) (ans int) {
	dp := map[int]int{}
	for _, v := range arr {
		dp[v] = dp[v-difference] + 1
		if dp[v] > ans {
			ans = dp[v]
		}
	}
	return
}

func longestSubsequence(arr []int, difference int) int {
	targetMap := make(map[int][]int)

	for _, item := range arr {
		var targetArr []int
		if _, ok := targetMap[item]; ok {
			targetArr = append(targetMap[item], item)
			delete(targetMap, item)
		} else {
			targetArr = []int{item}
		}

		targetKey := item + difference
		if _, ok := targetMap[targetKey]; !ok || (ok && len(targetArr) > len(targetMap[targetKey])) {
			// 等待值不存在，或者等待值存在且新的数组比旧的更长时
			targetMap[targetKey] = targetArr
		}
	}

	var maxLen int
	for _, value := range targetMap {
		if len(value) > maxLen {
			maxLen = len(value)
		}
	}
	return maxLen
}
