package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/magical-string/

481. 神奇字符串
中等

神奇字符串 s 仅由 '1' 和 '2' 组成，并需要遵守下面的规则：

神奇字符串 s 的神奇之处在于，串联字符串中 '1' 和 '2' 的连续出现次数可以生成该字符串。
s 的前几个元素是 s = "1221121221221121122……" 。如果将 s 中连续的若干 1 和 2 进行分组，可以得到 "1 22 11 2 1 22 1 22 11 2 11 22 ......" 。每组中 1 或者 2 的出现次数分别是 "1 2 2 1 1 2 1 2 2 1 2 2 ......" 。上面的出现次数正是 s 自身。

给你一个整数 n ，返回在神奇字符串 s 的前 n 个数字中 1 的数目。

示例 1：

输入：n = 6
输出：3
解释：神奇字符串 s 的前 6 个元素是 “122112”，它包含三个 1，因此返回 3 。
示例 2：

输入：n = 1
输出：1

提示：

1 <= n <= 10^5
*/
func main() {
	var tests = []struct {
		n    int
		want int
	}{
		{
			n:    6,
			want: 3,
		},
		{
			n:    1,
			want: 1,
		},
	}

	for _, item := range tests {
		if ans := magicalString(item.n); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func magicalString(n int) int {
	res := 1
	if n < 3 {
		return res
	}
	arr := []int{1, 2, 2}
	for i := 2; i < n; i++ {
		a := arr[i]
		if a == 1 {
			if arr[len(arr)-1] == 1 {
				arr = append(arr, 2)
			} else {
				arr = append(arr, 1)
			}
			res++
		} else {
			if arr[len(arr)-1] == 1 {
				arr = append(arr, 2, 2)
			} else {
				arr = append(arr, 1, 1)
			}
		}
	}

	return res
}
