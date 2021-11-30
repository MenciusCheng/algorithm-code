package main

import "fmt"

/*
859. 亲密字符串
给你两个字符串 s 和 goal ，只要我们可以通过交换 s 中的两个字母得到与 goal 相等的结果，就返回 true ；否则返回 false 。

交换字母的定义是：取两个下标 i 和 j （下标从 0 开始）且满足 i != j ，接着交换 s[i] 和 s[j] 处的字符。

例如，在 "abcd" 中交换下标 0 和下标 2 的元素可以生成 "cbad" 。
*/
func main() {
	fmt.Println(buddyStrings("ab", "ba") == true)
	fmt.Println(buddyStrings("ab", "ab") == false)
	fmt.Println(buddyStrings("aa", "aa") == true)
	fmt.Println(buddyStrings("aaaaaaabc", "aaaaaaacb") == true)
}

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	cnt := make(map[uint8]int, 0)

	var s1, g1 uint8
	var diffCount int
	var existSame bool
	for i := 0; i < len(s); i++ {
		if s[i] != goal[i] {
			diffCount++
			if diffCount > 2 {
				return false
			}

			if s1 == 0 {
				s1 = s[i]
				g1 = goal[i]
			} else if !(s1 == goal[i] && g1 == s[i]) {
				return false
			}
		} else {
			if cnt[s[i]] >= 1 {
				existSame = true
			}
			cnt[s[i]]++
		}
	}
	if diffCount == 1 || (diffCount == 0 && !existSame) {
		return false
	}

	return true
}
