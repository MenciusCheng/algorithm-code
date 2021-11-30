package main

import (
	"fmt"
	"math"
)

/*
https://leetcode-cn.com/problems/guess-number-higher-or-lower-ii/

难度：中等

375. 猜数字大小 II
我们正在玩一个猜数游戏，游戏规则如下：

我从 1 到 n 之间选择一个数字。
你来猜我选了哪个数字。
如果你猜到正确的数字，就会 赢得游戏 。
如果你猜错了，那么我会告诉你，我选的数字比你的 更大或者更小 ，并且你需要继续猜数。
每当你猜了数字 x 并且猜错了的时候，你需要支付金额为 x 的现金。如果你花光了钱，就会 输掉游戏 。
给你一个特定的数字 n ，返回能够 确保你获胜 的最小现金数，不管我选择那个数字 。
*/
func main() {
	//arr := buildArr(10)
	//fmt.Println(binarySearch(arr[7:], 10))
	//fmt.Println(binarySearch(arr[:6], 6))

	fmt.Println(getMoneyAmount(10) == 16)
	fmt.Println(getMoneyAmount(1) == 0)
	fmt.Println(getMoneyAmount(2) == 1)
	fmt.Println(getMoneyAmount(16) == 34)
}

// TODO 动态规划
func getMoneyAmount(n int) int {
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i := n - 1; i >= 1; i-- {
		for j := i + 1; j <= n; j++ {
			minCost := math.MaxInt32
			for k := i; k < j; k++ {
				cost := k + max(f[i][k-1], f[k+1][j])
				if cost < minCost {
					minCost = cost
				}
			}
			f[i][j] = minCost
		}
	}
	return f[1][n]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

//func getMoneyAmount(n int) int {
//	arr := buildArr(n)
//	return calCost(arr)
//}
//
//func calCost(arr []int) int {
//	switch len(arr) {
//	case 1:
//		return 0
//	case 2:
//		return arr[0]
//	case 3:
//		return arr[1]
//	default:
//		minCost := 0
//		for i := 1; i < len(arr)-1; i++ {
//			cost := arr[i]
//			leftCost := calCost(arr[:i])
//			rightCost := calCost(arr[i+1:])
//			if leftCost > rightCost {
//				cost += leftCost
//			} else {
//				cost += rightCost
//			}
//			if minCost == 0 || cost < minCost {
//				minCost = cost
//			}
//		}
//		return minCost
//	}
//}
//
//func buildArr(n int) []int {
//	arr := make([]int, n)
//	for i := 0; i < len(arr); i++ {
//		arr[i] = i+1
//	}
//	return arr
//}
