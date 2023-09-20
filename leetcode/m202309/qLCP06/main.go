package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/na-ying-bi/description/

LCP 06. 拿硬币
简单

桌上有 n 堆力扣币，每堆的数量保存在数组 coins 中。我们每次可以选择任意一堆，拿走其中的一枚或者两枚，求拿完所有力扣币的最少次数。

示例 1：
输入：[4,2,1]

输出：4
解释：第一堆力扣币最少需要拿 2 次，第二堆最少需要拿 1 次，第三堆最少需要拿 1 次，总共 4 次即可拿完。

示例 2：
输入：[2,3,10]
输出：8

限制：
1 <= n <= 4
*/
func main() {
	var tests = []struct {
		coins []int
		want  int
	}{
		{},
	}

	for _, item := range tests {
		if ans := minCount(item.coins); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func minCount(coins []int) int {
	res := 0
	for _, coin := range coins {
		res += coin / 2
		if coin%2 > 0 {
			res++
		}
	}

	return res
}
