package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/matchsticks-to-square/

473. 火柴拼正方形
你将得到一个整数数组 matchsticks ，其中 matchsticks[i] 是第 i 个火柴棒的长度。你要用 所有的火柴棍 拼成一个正方形。你 不能折断 任何一根火柴棒，但你可以把它们连在一起，而且每根火柴棒必须 使用一次 。
如果你能使这个正方形，则返回 true ，否则返回 false 。

示例 1:

输入: matchsticks = [1,1,2,2,2]
输出: true
解释: 能拼成一个边长为2的正方形，每边两根火柴。
示例 2:

输入: matchsticks = [3,3,3,3,4]
输出: false
解释: 不能用所有火柴拼成一个正方形。

提示:

1 <= matchsticks.length <= 15
1 <= matchsticks[i] <= 10^8
*/
func main() {
	var tests = []struct {
		matchsticks []int
		want        bool
	}{
		{
			matchsticks: []int{5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 3, 3},
			want:        true,
		},
		{
			matchsticks: []int{1, 1, 2, 2, 2},
			want:        true,
		},
		{
			matchsticks: []int{3, 3, 3, 3, 4},
			want:        false,
		},
	}

	for _, item := range tests {
		if ans := makesquare(item.matchsticks); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func makesquare(matchsticks []int) bool {
	sum := 0
	for _, item := range matchsticks {
		sum += item
	}
	if sum%4 != 0 {
		return false
	}
	side := sum / 4
	for _, item := range matchsticks {
		if item > side {
			return false
		}
	}

	return canSquare(matchsticks, side, 0)
}

func canSquare(matchsticks []int, side int, seq int) bool {
	if seq == 3 {
		sum := 0
		for _, item := range matchsticks {
			sum += item
		}
		return sum == side
	} else {
		n := len(matchsticks)
		if n < (4 - seq) {
			return false
		}
		for i := 1; i < (1 << n); i++ {
			sum := 0
			left := make([]int, 0)
			for j := 0; j < n; j++ {
				if (1<<j)&i == (1 << j) {
					sum += matchsticks[j]
				} else {
					left = append(left, matchsticks[j])
				}
			}
			if sum == side {
				if canSquare(left, side, seq+1) {
					return true
				}
			}
		}
		return false
	}
}
