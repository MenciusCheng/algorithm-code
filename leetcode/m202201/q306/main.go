package main

import (
	"fmt"
	"reflect"
	"strconv"
)

/*
https://leetcode-cn.com/problems/additive-number/

306. 累加数
累加数 是一个字符串，组成它的数字可以形成累加序列。
一个有效的 累加序列 必须 至少 包含 3 个数。除了最开始的两个数以外，字符串中的其他数都等于它之前两个数相加的和。
给你一个只包含数字 '0'-'9' 的字符串，编写一个算法来判断给定输入是否是 累加数 。如果是，返回 true ；否则，返回 false 。
说明：累加序列里的数 不会 以 0 开头，所以不会出现 1, 2, 03 或者 1, 02, 3 的情况。

示例 1：

输入："112358"
输出：true
解释：累加序列为: 1, 1, 2, 3, 5, 8 。1 + 1 = 2, 1 + 2 = 3, 2 + 3 = 5, 3 + 5 = 8
示例 2：

输入："199100199"
输出：true
解释：累加序列为: 1, 99, 100, 199。1 + 99 = 100, 99 + 100 = 199

提示：

1 <= num.length <= 35
num 仅由数字（0 - 9）组成

进阶：你计划如何处理由过大的整数输入导致的溢出?
*/
func main() {
	var tests = []struct {
		num  string
		want bool
	}{
		{
			num:  "112358",
			want: true,
		},
		{
			num:  "199100199",
			want: true,
		},
		{
			num:  "1023",
			want: false,
		},
		{
			num:  "101",
			want: true,
		},
	}

	for _, item := range tests {
		if ans := isAdditiveNumber(item.num); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func isAdditiveNumber(num string) bool {
	if len(num) < 3 {
		return false
	}
	for i := 1; i < len(num) && i <= 17; i++ {
		for j := 1; i+j < len(num) && j <= 17; j++ {
			if isAdd(num, 0, i, i, i+j) {
				return true
			}
		}
	}
	return false
}

func isAdd(num string, a0, a1 int, b0, b1 int) bool {
	aStr := num[a0:a1]
	if len(aStr) > 17 || len(aStr) > 1 && aStr[0] == '0' {
		return false
	}
	aNum, _ := strconv.Atoi(aStr)

	bStr := num[b0:b1]
	if len(bStr) > 17 || len(bStr) > 1 && bStr[0] == '0' {
		return false
	}
	bNum, _ := strconv.Atoi(bStr)

	sumNum := aNum + bNum
	sumStr := strconv.Itoa(sumNum)
	sumIndex := b1 + len(sumStr)

	if sumIndex > len(num) {
		return false
	}
	cStr := num[b1:sumIndex]
	if sumIndex == len(num) {
		return cStr == sumStr
	} else {
		return cStr == sumStr && isAdd(num, b0, b1, b1, sumIndex)
	}
}
