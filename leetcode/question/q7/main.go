package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
)

/*
https://leetcode-cn.com/problems/reverse-integer/

7. 整数反转
给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
如果反转后整数超过 32 位的有符号整数的范围 [−2^31,  2^31 − 1] ，就返回 0。
假设环境不允许存储 64 位整数（有符号或无符号）。

示例 1：
输入：x = 123
输出：321

示例 2：
输入：x = -123
输出：-321

示例 3：
输入：x = 120
输出：21

示例 4：
输入：x = 0
输出：0

提示：

-2^31 <= x <= 2^31 - 1
*/
func main() {
	var tests = []struct {
		x    int
		want int
	}{
		//{
		//	x:    123,
		//	want: 321,
		//},
		//{
		//	x:    120,
		//	want: 21,
		//},
		//{
		//	x:    0,
		//	want: 0,
		//},
		//{
		//	x:    -120,
		//	want: -21,
		//},
		{
			x:    1463847412,
			want: 2147483641,
		},
	}

	for _, item := range tests {
		if ans := reverse(item.x); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func reverse(x int) int {
	min := "-2147483648"
	max := "2147483647"

	for (x >= 10 || x <= -10) && x%10 == 0 {
		x /= 10
	}
	s := strconv.Itoa(x)
	rb := bytes.Buffer{}
	if x < 0 {
		rb.WriteByte('-')
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != '-' {
			rb.WriteByte(s[i])
		}
	}
	rs := rb.String()

	if x < 0 && len(rs) == len(min) {
		for i := 0; i < len(rs); i++ {
			if rs[i] > min[i] {
				return 0
			} else if rs[i] < min[i] {
				break
			}
		}
	} else if x >= 0 && len(rs) == len(max) {
		for i := 0; i < len(rs); i++ {
			if rs[i] > max[i] {
				return 0
			} else if rs[i] < max[i] {
				break
			}
		}
	}

	res, _ := strconv.Atoi(rs)
	return res
}
