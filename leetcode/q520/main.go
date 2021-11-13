package main

import (
	"fmt"
)

/*
https://leetcode-cn.com/problems/detect-capital/

520. 检测大写字母
我们定义，在以下情况时，单词的大写用法是正确的：

全部字母都是大写，比如 "USA" 。
单词中所有字母都不是大写，比如 "leetcode" 。
如果单词不只含有一个字母，只有首字母大写， 比如 "Google" 。
给你一个字符串 word 。如果大写用法正确，返回 true ；否则，返回 false 。
*/
func main() {
	fmt.Println(detectCapitalUse("USA") == true)
	fmt.Println(detectCapitalUse("FlaG") == false)
}

func detectCapitalUse(word string) bool {
	if len(word) == 1 {
		return true
	}
	var h int32
	isUpper := false
	for i, c := range word {
		if i == 0 {
			h = c
		} else if i == 1 {
			if 'A' <= h && h <= 'Z' {
				isUpper = 'A' <= c && c <= 'Z'
			} else {
				if 'A' <= c && c <= 'Z' {
					return false
				}
			}
		} else {
			if isUpper && 'a' <= c && c <= 'z' || !isUpper && 'A' <= c && c <= 'Z' {
				return false
			}
		}
	}
	return true
}
