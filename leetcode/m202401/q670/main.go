package main

import (
	"fmt"
	"reflect"
	"strconv"
)

/*
https://leetcode.cn/problems/maximum-swap/description/?envType=daily-question&envId=2024-01-22

670. 最大交换
中等
给定一个非负整数，你至多可以交换一次数字中的任意两位。返回你能得到的最大值。

示例 1 :

输入: 2736
输出: 7236
解释: 交换数字2和数字7。
示例 2 :

输入: 9973
输出: 9973
解释: 不需要交换。
注意:

给定数字的范围是 [0, 10^8]
*/
func main() {
	var tests = []struct {
		num  int
		want int
	}{
		{
			num:  2736,
			want: 7236,
		},
	}

	for _, item := range tests {
		if ans := maximumSwap(item.num); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func maximumSwap(num int) int {
	s := []byte(strconv.Itoa(num))

	for i := 0; i < len(s)-1; i++ {
		cj := i
		for j := i + 1; j < len(s); j++ {
			if s[j] > s[cj] || s[j] > s[i] && s[j] == s[cj] {
				cj = j
			}
		}
		if cj != i {
			s[i], s[cj] = s[cj], s[i]
			v, _ := strconv.Atoi(string(s))
			return v
		}
	}

	return num
}
