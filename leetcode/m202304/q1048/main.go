package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/longest-string-chain/

1048. 最长字符串链
中等
给出一个单词数组 words ，其中每个单词都由小写英文字母组成。
如果我们可以 不改变其他字符的顺序 ，在 wordA 的任何地方添加 恰好一个 字母使其变成 wordB ，那么我们认为 wordA 是 wordB 的 前身 。
例如，"abc" 是 "abac" 的 前身 ，而 "cba" 不是 "bcad" 的 前身
词链是单词 [word_1, word_2, ..., word_k] 组成的序列，k >= 1，其中 word1 是 word2 的前身，word2 是 word3 的前身，依此类推。一个单词通常是 k == 1 的 单词链 。
从给定单词列表 words 中选择单词组成词链，返回 词链的 最长可能长度 。

示例 1：

输入：words = ["a","b","ba","bca","bda","bdca"]
输出：4
解释：最长单词链之一为 ["a","ba","bda","bdca"]
示例 2:

输入：words = ["xbc","pcxbcf","xb","cxbc","pcxbc"]
输出：5
解释：所有的单词都可以放入单词链 ["xb", "xbc", "cxbc", "pcxbc", "pcxbcf"].
示例 3:

输入：words = ["abcd","dbqca"]
输出：1
解释：字链["abcd"]是最长的字链之一。
["abcd"，"dbqca"]不是一个有效的单词链，因为字母的顺序被改变了。

提示：

1 <= words.length <= 1000
1 <= words[i].length <= 16
words[i] 仅由小写英文字母组成。
*/
func main() {
	var tests = []struct {
		words []string
		want  int
	}{
		{
			words: []string{"a", "b", "ba", "bca", "bda", "bdca"},
			want:  4,
		},
		{
			words: []string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"},
			want:  5,
		},
		{
			words: []string{"abcd", "dbqca"},
			want:  1,
		},
	}

	for _, item := range tests {
		if ans := longestStrChain(item.words); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func longestStrChain(words []string) int {
	cnt := make(map[int][]string)
	for _, word := range words {
		cnt[len(word)] = append(cnt[len(word)], word)
	}
	chain := make(map[string][]string)
	for _, v := range cnt {
		for _, item := range v {
			for _, s := range cnt[len(item)+1] {
				if IsChain(item, s) {
					chain[item] = append(chain[item], s)
				}
			}
		}
	}
	cm := make(map[string]int)
	var cal func(s string) int
	cal = func(s string) int {
		if cm[s] > 0 {
			return cm[s]
		}
		max := 0
		for _, item := range chain[s] {
			v := cal(item)
			if v > max {
				max = v
			}
		}
		cm[s] = max + 1
		return cm[s]
	}

	max := 0
	for _, word := range words {
		v := cal(word)
		if v > max {
			max = v
		}
	}
	return max
}

func IsChain(s1, s2 string) bool {
	for i := 0; i < len(s2); i++ {
		if s1 == fmt.Sprintf("%s%s", s2[:i], s2[i+1:]) {
			return true
		}
	}
	return false
}
