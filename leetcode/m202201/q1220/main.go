package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode-cn.com/problems/count-vowels-permutation/

1220. 统计元音字母序列的数目
给你一个整数 n，请你帮忙统计一下我们可以按下述规则形成多少个长度为 n 的字符串：

字符串中的每个字符都应当是小写元音字母（'a', 'e', 'i', 'o', 'u'）
每个元音 'a' 后面都只能跟着 'e'
每个元音 'e' 后面只能跟着 'a' 或者是 'i'
每个元音 'i' 后面 不能 再跟着另一个 'i'
每个元音 'o' 后面只能跟着 'i' 或者是 'u'
每个元音 'u' 后面只能跟着 'a'
由于答案可能会很大，所以请你返回 模 10^9 + 7 之后的结果。

示例 1：

输入：n = 1
输出：5
解释：所有可能的字符串分别是："a", "e", "i" , "o" 和 "u"。
示例 2：

输入：n = 2
输出：10
解释：所有可能的字符串分别是："ae", "ea", "ei", "ia", "ie", "io", "iu", "oi", "ou" 和 "ua"。
示例 3：

输入：n = 5
输出：68

提示：

1 <= n <= 2 * 10^4
*/
func main() {
	var tests = []struct {
		n    int
		want int
	}{
		{
			n:    1,
			want: 5,
		},
		{
			n:    2,
			want: 10,
		},
		{
			n:    144,
			want: 18208803,
		},
	}

	for _, item := range tests {
		if ans := countVowelPermutation(item.n); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func countVowelPermutation(n int) int {
	cnt := map[byte]int{
		'a': 1,
		'e': 1,
		'i': 1,
		'o': 1,
		'u': 1,
	}

	for j := 1; j < n; j++ {
		cnt = map[byte]int{
			'a': add(add(cnt['e'], cnt['i']), cnt['u']),
			'e': add(cnt['a'], cnt['i']),
			'i': add(cnt['e'], cnt['o']),
			'o': cnt['i'],
			'u': add(cnt['i'], cnt['o']),
		}
	}

	sum := 0
	for _, v := range cnt {
		sum = add(sum, v)
	}
	return sum
}

func add(a, b int) int {
	return (a + b) % (1e9 + 7)
}
