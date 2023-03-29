package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/count-sorted-vowel-strings/

1641. 统计字典序元音字符串的数目
中等

给你一个整数 n，请返回长度为 n 、仅由元音 (a, e, i, o, u) 组成且按 字典序排列 的字符串数量。
字符串 s 按 字典序排列 需要满足：对于所有有效的 i，s[i] 在字母表中的位置总是与 s[i+1] 相同或在 s[i+1] 之前。

示例 1：

输入：n = 1
输出：5
解释：仅由元音组成的 5 个字典序字符串为 ["a","e","i","o","u"]
示例 2：

输入：n = 2
输出：15
解释：仅由元音组成的 15 个字典序字符串为
["aa","ae","ai","ao","au","ee","ei","eo","eu","ii","io","iu","oo","ou","uu"]
注意，"ea" 不是符合题意的字符串，因为 'e' 在字母表中的位置比 'a' 靠后
示例 3：

输入：n = 33
输出：66045

提示：

1 <= n <= 50
*/
func main() {
	var tests = []struct {
		n    int
		want int
	}{
		{
			n:    1,
			want: 5,
		},
		{
			n:    2,
			want: 15,
		},
		{
			n:    33,
			want: 66045,
		},
	}

	for _, item := range tests {
		if ans := countVowelStrings(item.n); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func countVowelStrings(n int) int {
	arr := []int{1, 1, 1, 1, 1}
	for i := 2; i <= n; i++ {
		arr2 := []int{1, 1, 1, 1, 1}
		arr2[1] += arr[1]
		arr2[2] += arr[1]
		arr2[3] += arr[1]
		arr2[4] += arr[1]
		arr2[2] += arr[2]
		arr2[3] += arr[2]
		arr2[4] += arr[2]
		arr2[3] += arr[3]
		arr2[4] += arr[3]
		arr2[4] += arr[4]
		arr = arr2
	}
	return arr[0] + arr[1] + arr[2] + arr[3] + arr[4]
}
