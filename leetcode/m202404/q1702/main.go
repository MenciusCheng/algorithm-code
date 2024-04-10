package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/maximum-binary-string-after-change/description/?envType=daily-question&envId=2024-04-10

1702. 修改后的最大二进制字符串
中等
提示
给你一个二进制字符串 binary ，它仅有 0 或者 1 组成。你可以使用下面的操作任意次对它进行修改：

操作 1 ：如果二进制串包含子字符串 "00" ，你可以用 "10" 将其替换。
比方说， "00010" -> "10010"
操作 2 ：如果二进制串包含子字符串 "10" ，你可以用 "01" 将其替换。
比方说， "00010" -> "00001"
请你返回执行上述操作任意次以后能得到的 最大二进制字符串 。如果二进制字符串 x 对应的十进制数字大于二进制字符串 y 对应的十进制数字，那么我们称二进制字符串 x 大于二进制字符串 y 。

示例 1：

输入：binary = "000110"
输出："111011"
解释：一个可行的转换为：
"000110" -> "000101"
"000101" -> "100101"
"100101" -> "110101"
"110101" -> "110011"
"110011" -> "111011"
示例 2：

输入：binary = "01"
输出："01"
解释："01" 没办法进行任何转换。

提示：

1 <= binary.length <= 10^5
binary 仅包含 '0' 和 '1' 。
*/
func main() {
	var tests = []struct {
		binary string
		want   string
	}{
		{
			binary: "000110",
			want:   "111011",
		},
		{
			binary: "01",
			want:   "01",
		},
		{
			binary: "1100",
			want:   "1110",
		},
	}

	for _, item := range tests {
		if ans := maximumBinaryString(item.binary); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func maximumBinaryString(binary string) string {
	zc := 0
	for i := range binary {
		if binary[i] == '0' {
			zc++
		}
	}
	if zc <= 1 {
		return binary
	}
	oneStart := -1
	for i := range binary {
		if binary[i] == '1' {
			oneStart = i
		} else {
			break
		}
	}
	if oneStart > -1 {
		binary = binary[oneStart+1:]
	}
	res := make([]byte, 0, len(binary))
	for i := 0; i <= oneStart; i++ {
		res = append(res, '1')
	}
	for i := 0; i < zc; i++ {
		if i != zc-1 {
			res = append(res, '1')
		} else {
			res = append(res, '0')
		}
	}
	oc := len(binary) - zc
	for i := 0; i < oc; i++ {
		res = append(res, '1')
	}

	return string(res)
}
