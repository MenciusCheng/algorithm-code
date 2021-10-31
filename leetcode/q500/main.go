package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/keyboard-row/
// 500. 键盘行
func main() {
	fmt.Println(findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
}

var (
	m1, m2, m3 = initMap()
)

func initMap() (m1 map[uint8]bool, m2 map[uint8]bool, m3 map[uint8]bool) {
	a1, a2, a3 := "qwertyuiop", "asdfghjkl", "zxcvbnm"
	m1, m2, m3 = make(map[uint8]bool), make(map[uint8]bool), make(map[uint8]bool)
	assign := func(a string, m map[uint8]bool) {
		for i := 0; i < len(a); i++ {
			m[a[i]] = true
		}
	}
	assign(a1, m1)
	assign(a2, m2)
	assign(a3, m3)
	return
}

func findWords(words []string) []string {
	res := make([]string, 0)
	for _, word := range words {
		if isWorkOnSameLine(word) {
			res = append(res, word)
		}
	}
	return res
}

func isWorkOnSameLine(word string) bool {
	if len(word) == 0 {
		return false
	}
	var m map[uint8]bool
	c := toLower(word[0])
	if _, exist := m1[c]; exist {
		m = m1
	} else if _, exist := m2[c]; exist {
		m = m2
	} else if _, exist := m3[c]; exist {
		m = m3
	} else {
		return false
	}
	for i := 1; i < len(word); i++ {
		if _, exist := m[toLower(word[i])]; !exist {
			return false
		}
	}
	return true
}

func toLower(c uint8) uint8 {
	if 'A' <= c && c <= 'Z' {
		c += 'a' - 'A'
	}
	return c
}
