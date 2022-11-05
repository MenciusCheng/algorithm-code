package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/parsing-a-boolean-expression/

1106. 解析布尔表达式
困难
124

给你一个以字符串形式表述的 布尔表达式（boolean） expression，返回该式的运算结果。

有效的表达式需遵循以下约定：

"t"，运算结果为 True
"f"，运算结果为 False
"!(expr)"，运算过程为对内部表达式 expr 进行逻辑 非的运算（NOT）
"&(expr1,expr2,...)"，运算过程为对 2 个或以上内部表达式 expr1, expr2, ... 进行逻辑 与的运算（AND）
"|(expr1,expr2,...)"，运算过程为对 2 个或以上内部表达式 expr1, expr2, ... 进行逻辑 或的运算（OR）

示例 1：

输入：expression = "!(f)"
输出：true
示例 2：

输入：expression = "|(f,t)"
输出：true
示例 3：

输入：expression = "&(t,f)"
输出：false
示例 4：

输入：expression = "|(&(t,f,t),!(t))"
输出：false

提示：

1 <= expression.length <= 20000
expression[i] 由 {'(', ')', '&', '|', '!', 't', 'f', ','} 中的字符组成。
expression 是以上述形式给出的有效表达式，表示一个布尔值。
*/
func main() {
	var tests = []struct {
		expression string
		want       bool
	}{
		{
			expression: "!(&(f,t))",
			want:       true,
		},
		{
			expression: "!(f)",
			want:       true,
		},
		{
			expression: "|(f,t)",
			want:       true,
		},
		{
			expression: "&(t,f)",
			want:       false,
		},
		{
			expression: "|(&(t,f,t),!(t))",
			want:       false,
		},
	}

	for _, item := range tests {
		if ans := parseBoolExpr(item.expression); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func parseBoolExpr(expression string) bool {
	stack := make([][]byte, 0)
	current := make([]byte, 0)

	for i := 0; i < len(expression); i++ {
		switch expression[i] {
		case '&', '|', '!':
			if len(current) > 0 {
				stack = append(stack, current)
				current = make([]byte, 0)
			}
			current = append(current, expression[i])
			i++
		case 't', 'f':
			current = append(current, expression[i])
		case ')':
			var f byte = 't'
			switch current[0] {
			case '&':
				for j := 1; j < len(current); j++ {
					if current[j] == 'f' {
						f = 'f'
						break
					}
				}
			case '|':
				f = 'f'
				for j := 1; j < len(current); j++ {
					if current[j] == 't' {
						f = 't'
						break
					}
				}
			case '!':
				if current[1] == 't' {
					f = 'f'
				}
			}
			if len(stack) > 0 {
				current = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			} else {
				current = make([]byte, 0)
			}
			current = append(current, f)
		}
	}

	return current[0] == 't'
}
