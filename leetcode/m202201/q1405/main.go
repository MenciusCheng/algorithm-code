package main

import (
	"fmt"
	"reflect"
	"sort"
)

/*
https://leetcode-cn.com/problems/longest-happy-string/

1405. 最长快乐字符串
如果字符串中不含有任何 'aaa'，'bbb' 或 'ccc' 这样的字符串作为子串，那么该字符串就是一个「快乐字符串」。
给你三个整数 a，b ，c，请你返回 任意一个 满足下列全部条件的字符串 s：

s 是一个尽可能长的快乐字符串。
s 中 最多 有a 个字母 'a'、b 个字母 'b'、c 个字母 'c' 。
s 中只含有 'a'、'b' 、'c' 三种字母。
如果不存在这样的字符串 s ，请返回一个空字符串 ""。

示例 1：

输入：a = 1, b = 1, c = 7
输出："ccaccbcc"
解释："ccbccacc" 也是一种正确答案。
示例 2：

输入：a = 2, b = 2, c = 1
输出："aabbc"
示例 3：

输入：a = 7, b = 1, c = 0
输出："aabaa"
解释：这是该测试用例的唯一正确答案。

提示：

0 <= a, b, c <= 100
a + b + c > 0
*/
func main() {
	var tests = []struct {
		a    int
		b    int
		c    int
		want string
	}{
		{
			a:    1,
			b:    1,
			c:    7,
			want: "ccaccbcc",
		},
		{
			a:    2,
			b:    2,
			c:    1,
			want: "aabbc",
		},
		{
			a:    7,
			b:    1,
			c:    0,
			want: "aabaa",
		},
	}

	for _, item := range tests {
		if ans := longestDiverseString(item.a, item.b, item.c); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func longestDiverseString(a int, b int, c int) string {
	arr := [][2]int{{0, a}, {1, b}, {2, c}}
	sb := make([]byte, 0)

	for arr[0][1] > 0 || arr[1][1] > 0 || arr[2][1] > 0 {
		sort.Slice(arr, func(i, j int) bool {
			return arr[i][1] > arr[j][1]
		})
		if arr[0][1] == 0 {
			break
		}

		ch := byte(arr[0][0]) + 'a'
		if len(sb) < 2 {
			sb = append(sb, ch)
			arr[0][1]--
		} else {
			if sb[len(sb)-1] != ch || sb[len(sb)-2] != ch {
				sb = append(sb, ch)
				arr[0][1]--
			} else if arr[1][1] > 0 {
				ch = byte(arr[1][0]) + 'a'
				sb = append(sb, ch)
				arr[1][1]--
			} else {
				break
			}
		}
	}

	return string(sb)
}
