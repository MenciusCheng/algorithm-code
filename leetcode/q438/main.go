package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/

438. 找到字符串中所有字母异位词
给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。

异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。

示例 1:

输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
 示例 2:

输入: s = "abab", p = "ab"
输出: [0,1,2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。

提示:

1 <= s.length, p.length <= 3 * 104
s 和 p 仅包含小写字母
*/
func main() {
	fmt.Println(reflect.DeepEqual(findAnagrams("cbaebabacd", "abc"), []int{0, 6}))
	fmt.Println(reflect.DeepEqual(findAnagrams("abab", "ab"), []int{0, 1, 2}))

}

func findAnagrams(s string, p string) []int {
	sLen := len(s)
	pLen := len(p)

	res := make([]int, 0)
	if sLen < pLen {
		return res
	}

	sCnt := [26]uint8{}
	pCnt := [26]uint8{}
	for i := 0; i < pLen; i++ {
		sCnt[s[i]-'a']++
		pCnt[p[i]-'a']++
	}

	if sCnt == pCnt {
		res = append(res, 0)
	}
	for i := 0; i < sLen-pLen; i++ {
		sCnt[s[i]-'a']--
		sCnt[s[i+pLen]-'a']++
		if sCnt == pCnt {
			res = append(res, i+1)
		}
	}
	return res
}
