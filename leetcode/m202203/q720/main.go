package main

import (
	"fmt"
	"reflect"
	"sort"
)

/*
https://leetcode-cn.com/problems/longest-word-in-dictionary/

720. 词典中最长的单词
给出一个字符串数组 words 组成的一本英语词典。返回 words 中最长的一个单词，该单词是由 words 词典中其他单词逐步添加一个字母组成。

若其中有多个可行的答案，则返回答案中字典序最小的单词。若无答案，则返回空字符串。

示例 1：

输入：words = ["w","wo","wor","worl", "world"]
输出："world"
解释： 单词"world"可由"w", "wo", "wor", 和 "worl"逐步添加一个字母组成。
示例 2：

输入：words = ["a", "banana", "app", "appl", "ap", "apply", "apple"]
输出："apple"
解释："apply" 和 "apple" 都能由词典中的单词组成。但是 "apple" 的字典序小于 "apply"

提示：

1 <= words.length <= 1000
1 <= words[i].length <= 30
所有输入的字符串 words[i] 都只包含小写字母。
*/
func main() {
	var tests = []struct {
		words []string
		want  string
	}{
		{
			words: []string{"w", "wo", "wor", "worl", "world"},
			want:  "world",
		},
		{
			words: []string{"a", "banana", "app", "appl", "ap", "apply", "apple"},
			want:  "apple",
		},
	}

	for _, item := range tests {
		if ans := longestWord(item.words); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func longestWord(words []string) string {
	cnt := make(map[string]bool)
	for _, word := range words {
		cnt[word] = true
	}
	sort.Slice(words, func(i, j int) bool {
		return words[i] < words[j]
	})

	res := ""
	for i := 0; i < len(words); i++ {
		word := words[i]
		if len(word) == 1 {
			if len(word) > len(res) {
				res = word
			}
		} else {
			if cnt[word[:len(word)-1]] {
				if len(word) > len(res) {
					res = word
				}
			} else {
				cnt[word] = false
			}
		}
	}

	return res
}
