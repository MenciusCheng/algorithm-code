package main

import (
	"fmt"
	"reflect"
	"strconv"
)

/*
https://leetcode-cn.com/problems/self-dividing-numbers/

728. 自除数
自除数 是指可以被它包含的每一位数整除的数。

例如，128 是一个 自除数 ，因为 128 % 1 == 0，128 % 2 == 0，128 % 8 == 0。
自除数 不允许包含 0 。
给定两个整数 left 和 right ，返回一个列表，列表的元素是范围 [left, right] 内所有的 自除数 。

示例 1：

输入：left = 1, right = 22
输出：[1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22]
示例 2:

输入：left = 47, right = 85
输出：[48,55,66,77]

提示：

1 <= left <= right <= 10^4
*/
func main() {
	var tests = []struct {
		left  int
		right int
		want  []int
	}{
		{
			left:  1,
			right: 22,
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22},
		},
		{
			left:  47,
			right: 85,
			want:  []int{48, 55, 66, 77},
		},
	}

	for _, item := range tests {
		if ans := selfDividingNumbers(item.left, item.right); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func selfDividingNumbers(left int, right int) []int {
	res := make([]int, 0)
	for i := left; i <= right; i++ {
		num := i
		ns := strconv.Itoa(num)
		flag := true
		for j := 0; j < len(ns); j++ {
			if ns[j]-'0' == 0 || num%int(ns[j]-'0') != 0 {
				flag = false
				break
			}
		}
		if flag {
			res = append(res, num)
		}
	}

	return res
}
