package main

import (
	"fmt"
)

/*
https://leetcode.cn/problems/prefix-and-suffix-search/

745. 前缀和后缀搜索
设计一个包含一些单词的特殊词典，并能够通过前缀和后缀来检索单词。

实现 WordFilter 类：
WordFilter(string[] words) 使用词典中的单词 words 初始化对象。
f(string pref, string suff) 返回词典中具有前缀 prefix 和后缀 suff 的单词的下标。如果存在不止一个满足要求的下标，返回其中 最大的下标 。如果不存在这样的单词，返回 -1 。

示例：

输入
["WordFilter", "f"]
[[["apple"]], ["a", "e"]]
输出
[null, 0]
解释
WordFilter wordFilter = new WordFilter(["apple"]);
wordFilter.f("a", "e"); // 返回 0 ，因为下标为 0 的单词：前缀 prefix = "a" 且 后缀 suff = "e" 。

提示：

1 <= words.length <= 10^4
1 <= words[i].length <= 7
1 <= pref.length, suff.length <= 7
words[i]、pref 和 suff 仅由小写英文字母组成
最多对函数 f 执行 10^4 次调用
*/
func main() {
	wordFilter := Constructor([]string{"apple"})
	fmt.Println(wordFilter.F("a", "e") == 0)
}

type WordFilter struct {
	PrefMap map[string][]int
	SuffMap map[string]map[int]bool
}

func Constructor(words []string) WordFilter {
	wordFilter := WordFilter{
		PrefMap: make(map[string][]int),
		SuffMap: make(map[string]map[int]bool),
	}
	for i, word := range words {
		for j := 1; j <= len(word); j++ {
			wordFilter.PrefMap[word[:j]] = append(wordFilter.PrefMap[word[:j]], i)
		}
		for j := len(word) - 1; j >= 0; j-- {
			if wordFilter.SuffMap[word[j:]] == nil {
				wordFilter.SuffMap[word[j:]] = make(map[int]bool)
			}
			wordFilter.SuffMap[word[j:]][i] = true
		}
	}

	return wordFilter
}

func (this *WordFilter) F(pref string, suff string) int {
	prefIds := this.PrefMap[pref]
	suffMap := this.SuffMap[suff]
	if len(prefIds) == 0 || len(suffMap) == 0 {
		return -1
	}

	for i := len(prefIds) - 1; i >= 0; i-- {
		if suffMap[prefIds[i]] {
			return prefIds[i]
		}
	}

	return -1
}
