package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode-cn.com/problems/elimination-game/

390. 消除游戏
给定一个从1 到 n 排序的整数列表。
首先，从左到右，从第一个数字开始，每隔一个数字进行删除，直到列表的末尾。
第二步，在剩下的数字中，从右到左，从倒数第一个数字开始，每隔一个数字进行删除，直到列表开头。
我们不断重复这两步，从左到右和从右到左交替进行，直到只剩下一个数字。
返回长度为 n 的列表中，最后剩下的数字。

示例：

输入:
n = 9,
1 2 3 4 5 6 7 8 9
2 4 6 8
2 6
6

输出:
6

提示：

1 <= n <= 10^9
*/
func main() {
	var tests = []struct {
		n    int
		want int
	}{
		{
			n:    9,
			want: 6,
		},
		{
			n:    1,
			want: 1,
		},
	}

	for _, item := range tests {
		if ans := lastRemaining(item.n); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func lastRemaining(n int) int {
	if n == 1 {
		return 1
	}

	first := 1
	last := n
	count := n
	gap := 1

	isLeft := true
	for count > 1 {
		gap *= 2
		if isLeft {
			first += gap / 2
			var i = first
			for i+gap <= last {
				i += gap
			}
			last = i
			isLeft = false
		} else {
			if count%2 > 0 {
				first += gap / 2
			}
			var i = first
			for i+gap <= last {
				i += gap
			}
			last = i
			isLeft = true
		}
		count /= 2
	}
	return first
}
