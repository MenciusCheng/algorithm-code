package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
https://leetcode-cn.com/problems/reconstruct-original-digits-from-english/

423. 从英文中重建数字
给你一个字符串 s ，其中包含字母顺序打乱的用英文单词表示的若干数字（0-9）。按 升序 返回原始的数字。

示例 1：

输入：s = "owoztneoer"
输出："012"
示例 2：

输入：s = "fviefuro"
输出："45"

提示：

1 <= s.length <= 105
s[i] 为 ["e","g","f","i","h","o","n","s","r","u","t","w","v","x","z"] 这些字符之一
s 保证是一个符合题目要求的字符串
*/
func main() {
	fmt.Println(originalDigits("owoztneoer") == "012")
	fmt.Println(originalDigits("fviefuro") == "45")
	fmt.Println(originalDigits("zerozero") == "00")
	fmt.Println(originalDigits("zeroonetwothreefourfivesixseveneightnine") == "0123456789")
}

var allNums = func() []map[int32]int {
	//cs := []string{
	//	"0zero", "1one", "2two", "3three", "4four", "5five", "6six", "7seven", "8eight", "9nine",
	//}
	// 为了避免少字母的数字先把其他数字给占了，所以把有特征的数字放在前面判断
	cs := []string{
		"0zero", "2two", "4four", "6six", "8eight", "5five", "3three", "7seven", "1one", "9nine",
	}
	res := make([]map[int32]int, 0)

	for _, num := range cs {
		cnt := make(map[int32]int)
		for i, c := range num {
			if i == 0 {
				cnt['!'] = int(9 - ('9' - c))
			} else {
				cnt[c]++
			}
		}
		res = append(res, cnt)
	}
	return res
}()

func originalDigits(s string) string {
	cnt := make(map[int32]int)
	for _, c := range s {
		cnt[c]++
	}
	arr := make([]int, 0)

	for _, numMap := range allNums {
		var num int
		exist := true
		for exist {
			for k, v := range numMap {
				if k == '!' {
					num = v
				} else if cnt[k] < v {
					exist = false
					break
				}
			}
			if exist {
				for k, v := range numMap {
					cnt[k] -= v
				}
				arr = append(arr, num)
			}
		}
	}
	sort.Ints(arr)
	var b strings.Builder
	for _, i := range arr {
		b.WriteString(strconv.Itoa(i))
	}
	return b.String()
}
