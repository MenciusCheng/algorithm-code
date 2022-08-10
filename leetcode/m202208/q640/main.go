package main

import (
	"fmt"
	"reflect"
	"strings"
)

/*
https://leetcode.cn/problems/solve-the-equation/

640. 求解方程
求解一个给定的方程，将x以字符串 "x=#value" 的形式返回。该方程仅包含 '+' ， '-' 操作，变量 x 和其对应系数。
如果方程没有解，请返回 "No solution" 。如果方程有无限解，则返回 “Infinite solutions” 。
题目保证，如果方程中只有一个解，则 'x' 的值是一个整数。

示例 1：

输入: equation = "x+5-3+x=6+x-2"
输出: "x=2"
示例 2:

输入: equation = "x=x"
输出: "Infinite solutions"
示例 3:

输入: equation = "2x=x"
输出: "x=0"

提示:

3 <= equation.length <= 1000
equation 只有一个 '='.
equation 方程由整数组成，其绝对值在 [0, 100] 范围内，不含前导零和变量 'x'
*/
func main() {
	var tests = []struct {
		equation string
		want     string
	}{
		{
			equation: "2x+3x-6x=x+2",
			want:     "x=-1",
		},
		{
			equation: "x+5-3+x=6+x-2",
			want:     "x=2",
		},
		{
			equation: "x=x",
			want:     "Infinite solutions",
		},
		{
			equation: "x=x+2",
			want:     "No solution",
		},
		{
			equation: "2x=x",
			want:     "x=0",
		},
		{
			equation: "0x=0",
			want:     "Infinite solutions",
		},
	}

	for _, item := range tests {
		if ans := solveEquation(item.equation); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func solveEquation(equation string) string {
	split := strings.Split(equation, "=")

	v0, xcount0 := calEx(split[0])
	v1, xcount1 := calEx(split[1])
	v, xcount := v0-v1, xcount1-xcount0

	if xcount == 0 {
		if v == 0 {
			return "Infinite solutions"
		} else {
			return "No solution"
		}
	}
	return fmt.Sprintf("x=%d", v/xcount)
}

func calEx(exp string) (int, int) {
	v := 0
	xcount := 0
	neg := 1
	lastV := 0
	for i := 0; i < len(exp); i++ {
		item := exp[i]
		switch item {
		case 'x':
			if lastV > 0 {
				xcount += neg * lastV
				lastV = 0
			} else if i == 0 || exp[i-1] != '0' {
				xcount += neg
			}
		case '+':
			v = v + neg*lastV
			lastV = 0
			neg = 1
		case '-':
			v = v + neg*lastV
			lastV = 0
			neg = -1
		default:
			lastV = lastV*10 + int(item-'0')
			if i == len(exp)-1 {
				v = v + neg*lastV
			}
		}
	}
	return v, xcount
}
