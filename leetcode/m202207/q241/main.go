package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
)

/*
https://leetcode.cn/problems/different-ways-to-add-parentheses/

241. 为运算表达式设计优先级
给你一个由数字和运算符组成的字符串 expression ，按不同优先级组合数字和运算符，计算并返回所有可能组合的结果。你可以 按任意顺序 返回答案。
生成的测试用例满足其对应输出值符合 32 位整数范围，不同结果的数量不超过 10^4 。

示例 1：

输入：expression = "2-1-1"
输出：[0,2]
解释：
((2-1)-1) = 0
(2-(1-1)) = 2
示例 2：

输入：expression = "2*3-4*5"
输出：[-34,-14,-10,-10,10]
解释：
(2*(3-(4*5))) = -34
((2*3)-(4*5)) = -14
((2*(3-4))*5) = -10
(2*((3-4)*5)) = -10
(((2*3)-4)*5) = 10

提示：

1 <= expression.length <= 20
expression 由数字和算符 '+'、'-' 和 '*' 组成。
输入表达式中的所有整数值在范围 [0, 99]
*/
func main() {
	var tests = []struct {
		expression string
		want       []int
	}{
		{
			expression: "2-1-1",
			want:       []int{0, 2},
		},
		{
			expression: "2*3-4*5",
			want:       []int{-34, -14, -10, -10, 10},
		},
		{
			expression: "15*1*4",
			want:       []int{60, 60},
		},
	}

	for _, item := range tests {
		if ans := diffWaysToCompute(item.expression); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func diffWaysToCompute(expression string) []int {
	return []int{}
}

func cal(a, b string) string {
	bs := bytes.Buffer{}

	var x1 int
	if a[0] == '+' || a[0] == '-' || a[0] == '*' {
		bs.WriteByte(a[0])
		x1, _ = strconv.Atoi(a[1:])
	} else {
		x1, _ = strconv.Atoi(a)
	}
	x2, _ := strconv.Atoi(b[1:])

	var r int
	switch b[0] {
	case '+':
		r = x1 + x2
	case '-':
		r = x1 - x2
	case '*':
		r = x1 * x2
	}
	bs.WriteString(strconv.Itoa(r))
	return bs.String()
}

func split(expression string) []string {
	res := make([]string, 0)

	bs := bytes.Buffer{}
	for i := 0; i < len(expression); i++ {
		if expression[i] == '+' || expression[i] == '-' || expression[i] == '*' {
			res = append(res, bs.String())
			bs = bytes.Buffer{}
			bs.WriteByte(expression[i])
			i++
			bs.WriteByte(expression[i])
		} else {
			bs.WriteByte(expression[i])
		}
	}
	if bs.Len() > 0 {
		res = append(res, bs.String())
	}
	return res
}
