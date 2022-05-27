package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/find-closest-lcci/

17 面试题  单词距离
有个内含单词的超大文本文件，给定任意两个不同的单词，找出在这个文件中这两个单词的最短距离(相隔单词数)。如果寻找过程在这个文件中会重复多次，而每次寻找的单词不同，你能对此优化吗?

示例：

输入：words = ["I","am","a","student","from","a","university","in","a","city"], word1 = "a", word2 = "student"
输出：1
提示：

words.length <= 100000
*/
func main() {
	var tests = []struct {
		words []string
		word1 string
		word2 string
		want  int
	}{
		{
			words: []string{"I", "am", "a", "student", "from", "a", "university", "in", "a", "city"},
			word1: "a",
			word2: "student",
			want:  1,
		},
	}

	for _, item := range tests {
		if ans := findClosest(item.words, item.word1, item.word2); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func findClosest(words []string, word1 string, word2 string) int {
	wMap := make(map[string][]int)
	for i, w := range words {
		wMap[w] = append(wMap[w], i)
	}

	arr1 := wMap[word1]
	arr2 := wMap[word2]
	res := 100000
	for _, a1 := range arr1 {
		for _, a2 := range arr2 {
			res = min(res, abs(a1, a2))
		}
	}
	return res
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
