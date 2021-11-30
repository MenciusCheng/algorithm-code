package main

import "fmt"

/*
https://leetcode-cn.com/problems/maximum-product-of-word-lengths/

318. 最大单词长度乘积
给定一个字符串数组 words，找到 length(word[i]) * length(word[j]) 的最大值，并且这两个单词不含有公共字母。你可以认为每个单词只包含小写字母。如果不存在这样的两个单词，返回 0。
*/
func main() {
	fmt.Println(maxProduct([]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}) == 16)
	fmt.Println(maxProduct([]string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}) == 4)
	fmt.Println(maxProduct([]string{"a", "aa", "aaa", "aaaa"}) == 0)
}

// 这个方法时间复杂度高，改成用位运算方式更好
func maxProduct(words []string) int {
	var max int
	for i := 0; i < len(words)-1; i++ {
		word := words[i]
		wMap := make(map[int32]bool)
		for _, w := range word {
			wMap[w] = true
		}

		for j := i + 1; j < len(words); j++ {
			jWord := words[j]
			isSame := false
			for _, w := range jWord {
				if wMap[w] {
					isSame = true
					break
				}
			}
			if !isSame && len(word)*len(jWord) > max {
				max = len(word) * len(jWord)
			}
		}
	}
	return max
}
