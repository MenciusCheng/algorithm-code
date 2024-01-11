package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/minimum-additions-to-make-valid-string/description/

2645. 构造有效字符串的最少插入数
中等

给你一个字符串 word ，你可以向其中任何位置插入 "a"、"b" 或 "c" 任意次，返回使 word 有效 需要插入的最少字母数。
如果字符串可以由 "abc" 串联多次得到，则认为该字符串 有效 。

示例 1：

输入：word = "b"
输出：2
解释：在 "b" 之前插入 "a" ，在 "b" 之后插入 "c" 可以得到有效字符串 "abc" 。
示例 2：

输入：word = "aaa"
输出：6
解释：在每个 "a" 之后依次插入 "b" 和 "c" 可以得到有效字符串 "abcabcabc" 。
示例 3：

输入：word = "abc"
输出：0
解释：word 已经是有效字符串，不需要进行修改。

提示：

1 <= word.length <= 50
word 仅由字母 "a"、"b" 和 "c" 组成。
*/
func main() {
	var tests = []struct {
		word string
		want int
	}{
		{
			word: "b",
			want: 2,
		},
		{
			word: "aaa",
			want: 6,
		},
		{
			word: "abc",
			want: 0,
		},
	}

	for _, item := range tests {
		if ans := addMinimum(item.word); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func addMinimum(word string) int {
	res := 0
	for i := 0; i < len(word); i++ {
		if word[i] == 'a' {
			if i+1 < len(word) && word[i+1] == 'b' {
				if i+2 < len(word) && word[i+2] == 'c' {
					i += 2
				} else {
					res += 1
					i++
				}
			} else if i+1 < len(word) && word[i+1] == 'c' {
				res += 1
				i++
			} else {
				res += 2
			}
		} else if word[i] == 'b' {
			if i+1 < len(word) && word[i+1] == 'c' {
				res += 1
				i++
			} else {
				res += 2
			}
		} else {
			res += 2
		}
	}

	return res
}
