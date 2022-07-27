package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/fraction-addition-and-subtraction/

592. 分数加减运算
给定一个表示分数加减运算的字符串 expression ，你需要返回一个字符串形式的计算结果。
这个结果应该是不可约分的分数，即最简分数。 如果最终结果是一个整数，例如 2，你需要将它转换成分数形式，其分母为 1。所以在上述例子中, 2 应该被转换为 2/1。

示例 1:

输入: expression = "-1/2+1/2"
输出: "0/1"
 示例 2:

输入: expression = "-1/2+1/2+1/3"
输出: "1/3"
示例 3:

输入: expression = "1/3-1/2"
输出: "-1/6"

提示:

输入和输出字符串只包含 '0' 到 '9' 的数字，以及 '/', '+' 和 '-'。
输入和输出分数格式均为 ±分子/分母。如果输入的第一个分数或者输出的分数是正数，则 '+' 会被省略掉。
输入只包含合法的最简分数，每个分数的分子与分母的范围是  [1,10]。 如果分母是1，意味着这个分数实际上是一个整数。
输入的分数个数范围是 [1,10]。
最终结果的分子与分母保证是 32 位整数范围内的有效整数。
*/
func main() {
	var tests = []struct {
		expression string
		want       string
	}{
		{
			expression: "5/3+1/3",
			want:       "2/1",
		},
		{
			expression: "-1/2+1/2",
			want:       "0/1",
		},
		{
			expression: "-1/2+1/2+1/3",
			want:       "1/3",
		},
		{
			expression: "1/3-1/2",
			want:       "-1/6",
		},
		{
			expression: "-5/2+10/3+7/9",
			want:       "29/18",
		},
	}

	for _, item := range tests {
		if ans := fractionAddition(item.expression); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func fractionAddition(expression string) string {
	isNeg := false
	arr := make([]int, 0)

	for i := 0; i < len(expression); i++ {
		switch expression[i] {
		case '-':
			isNeg = true
		case '+', '/':
			continue
		default:
			a := int(expression[i] - '0')
			if i+1 < len(expression) && expression[i+1] == '0' {
				i++
				a = a*10 + int(expression[i]-'0')
			}
			if isNeg {
				a = -a
				isNeg = false
			}
			arr = append(arr, a)
		}
	}

	a1 := arr[0]
	a2 := arr[1]
	a1, a2 = change(a1, a2)
	for i := 2; i < len(arr); i += 2 {
		item, _ := change(arr[i], arr[i+1])
		a1 += item
	}

	if a1 == 0 {
		return "0/1"
	}
	for i := 9; i > 1; i-- {
		if a1%i == 0 && a2%i == 0 {
			a1 /= i
			a2 /= i
		}
	}
	return fmt.Sprintf("%d/%d", a1, a2)
}

func change(a1, a2 int) (int, int) {
	if a1 == 0 {
		return a1, a2
	}
	d := 5040 / a2
	return a1 * d, 5040
}
