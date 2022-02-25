package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

/*
https://leetcode-cn.com/problems/complex-number-multiplication/

537. 复数乘法
复数 可以用字符串表示，遵循 "实部+虚部i" 的形式，并满足下述条件：

实部 是一个整数，取值范围是 [-100, 100]
虚部 也是一个整数，取值范围是 [-100, 100]
i2 == -1
给你两个字符串表示的复数 num1 和 num2 ，请你遵循复数表示形式，返回表示它们乘积的字符串。

示例 1：

输入：num1 = "1+1i", num2 = "1+1i"
输出："0+2i"
解释：(1 + i) * (1 + i) = 1 + i2 + 2 * i = 2i ，你需要将它转换为 0+2i 的形式。
示例 2：

输入：num1 = "1+-1i", num2 = "1+-1i"
输出："0+-2i"
解释：(1 - i) * (1 - i) = 1 + i2 - 2 * i = -2i ，你需要将它转换为 0+-2i 的形式。

提示：

num1 和 num2 都是有效的复数表示。
*/
func main() {
	var tests = []struct {
		num1 string
		num2 string
		want string
	}{
		{
			num1: "1+1i",
			num2: "1+1i",
			want: "0+2i",
		},
		{
			num1: "1+-1i",
			num2: "1+-1i",
			want: "0+-2i",
		},
	}

	for _, item := range tests {
		if ans := complexNumberMultiply(item.num1, item.num2); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func complexNumberMultiply(num1 string, num2 string) string {
	split1 := strings.Split(num1, "+")
	a0, _ := strconv.Atoi(split1[0])
	a1, _ := strconv.Atoi(split1[1][:len(split1[1])-1])

	split2 := strings.Split(num2, "+")
	b0, _ := strconv.Atoi(split2[0])
	b1, _ := strconv.Atoi(split2[1][:len(split2[1])-1])

	s0 := a0*b0 + a1*b1*-1
	s1 := a0*b1 + a1*b0
	return fmt.Sprintf("%d+%di", s0, s1)
}
