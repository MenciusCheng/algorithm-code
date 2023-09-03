package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/minimum-operations-to-make-a-special-number/description/

8040. 生成特殊数字的最少操作
中等

给你一个下标从 0 开始的字符串 num ，表示一个非负整数。
在一次操作中，您可以选择 num 的任意一位数字并将其删除。请注意，如果你删除 num 中的所有数字，则 num 变为 0。
返回最少需要多少次操作可以使 num 变成特殊数字。
如果整数 x 能被 25 整除，则该整数 x 被认为是特殊数字。

示例 1：

输入：num = "2245047"
输出：2
解释：删除数字 num[5] 和 num[6] ，得到数字 "22450" ，可以被 25 整除。
可以证明要使数字变成特殊数字，最少需要删除 2 位数字。
示例 2：

输入：num = "2908305"
输出：3
解释：删除 num[3]、num[4] 和 num[6] ，得到数字 "2900" ，可以被 25 整除。
可以证明要使数字变成特殊数字，最少需要删除 3 位数字。
示例 3：

输入：num = "10"
输出：1
解释：删除 num[0] ，得到数字 "0" ，可以被 25 整除。
可以证明要使数字变成特殊数字，最少需要删除 1 位数字。

提示

1 <= num.length <= 100
num 仅由数字 '0' 到 '9' 组成
num 不含任何前导零
*/
func main() {
	var tests = []struct {
		num  string
		want int
	}{
		{
			num:  "2245047",
			want: 2,
		},
		{
			num:  "2908305",
			want: 3,
		},
		{
			num:  "10",
			want: 1,
		},
	}

	for _, item := range tests {
		if ans := minimumOperations(item.num); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func minimumOperations(num string) int {
	res := len(num)
	// 结尾为 00, 25, 50, 75
	var h0, h5 bool
	for i := len(num) - 1; i > 0; i-- {
		if !h0 && num[i] == '0' {
			h0 = true
			for j := i - 1; j >= 0; j-- {
				if num[j] == '0' {
					res = min(res, len(num)-j-2)
					break
				}
			}
			for j := i - 1; j >= 0; j-- {
				if num[j] == '5' {
					res = min(res, len(num)-j-2)
					break
				}
			}
		}
		if !h5 && num[i] == '5' {
			h5 = true
			for j := i - 1; j >= 0; j-- {
				if num[j] == '2' {
					res = min(res, len(num)-j-2)
					break
				}
			}
			for j := i - 1; j >= 0; j-- {
				if num[j] == '7' {
					res = min(res, len(num)-j-2)
					break
				}
			}
		}
	}
	if h0 {
		res = min(res, len(num)-1)
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
