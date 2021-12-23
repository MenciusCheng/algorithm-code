package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode-cn.com/problems/longest-duplicate-substring/

1044. 最长重复子串
难度：困难
给你一个字符串 s ，考虑其所有 重复子串 ：即，s 的连续子串，在 s 中出现 2 次或更多次。这些出现之间可能存在重叠。
返回 任意一个 可能具有最长长度的重复子串。如果 s 不含重复子串，那么答案为 "" 。

示例 1：

输入：s = "banana"
输出："ana"
示例 2：

输入：s = "abcd"
输出：""

提示：

2 <= s.length <= 3 * 10^4
s 由小写英文字母组成
*/
func main() {
	var tests = []struct {
		s    string
		want string
	}{
		{},
	}

	for _, item := range tests {
		if ans := longestDupSubstring(item.s); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

//func randInt(a, b int) int {
//	return a + rand.Intn(b-a)
//}

func pow(x, n, mod int) int {
	res := 1
	for ; n > 0; n >>= 1 {
		if n&1 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

func check(arr []byte, m, a1, a2, mod1, mod2 int) int {
	aL1, aL2 := pow(a1, m, mod1), pow(a2, m, mod2)
	h1, h2 := 0, 0
	for _, c := range arr[:m] {
		h1 = (h1*a1 + int(c)) % mod1
		h2 = (h2*a2 + int(c)) % mod2
	}
	// 存储一个编码组合是否出现过
	seen := map[[2]int]bool{{h1, h2}: true}
	for start := 1; start <= len(arr)-m; start++ {
		h1 = (h1*a1 - int(arr[start-1])*aL1 + int(arr[start+m-1])) % mod1
		if h1 < 0 {
			h1 += mod1
		}
		h2 = (h2*a2 - int(arr[start-1])*aL2 + int(arr[start+m-1])) % mod2
		if h2 < 0 {
			h2 += mod2
		}
		// 如果重复，则返回重复串的起点
		if seen[[2]int{h1, h2}] {
			return start
		}
		seen[[2]int{h1, h2}] = true
	}
	// 没有重复，则返回 -1
	return -1
}

func longestDupSubstring(s string) string {
	//rand.Seed(time.Now().UnixNano())
	// 生成两个进制
	//a1, a2 := randInt(26, 100), randInt(26, 100)
	a1, a2 := 31, 37
	// 生成两个模
	//mod1, mod2 := randInt(1e9+7, math.MaxInt32), randInt(1e9+7, math.MaxInt32)
	mod := int(1e9 + 7)
	// 先对所有字符进行编码
	arr := []byte(s)
	for i := range arr {
		arr[i] -= 'a'
	}

	l, r := 1, len(s)-1
	length, start := 0, -1
	for l <= r {
		m := l + (r-l+1)/2
		idx := check(arr, m, a1, a2, mod, mod)
		if idx != -1 { // 有重复子串，移动左边界
			l = m + 1
			length = m
			start = idx
		} else { // 无重复子串，移动右边界
			r = m - 1
		}
	}
	if start == -1 {
		return ""
	}
	return s[start : start+length]
}
