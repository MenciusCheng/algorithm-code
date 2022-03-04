package leetcode

import "testing"

func TestGen(t *testing.T) {
	desc := `
258. 各位相加
给定一个非负整数 num，反复将各个位上的数字相加，直到结果为一位数。返回这个结果。

示例 1:

输入: num = 38
输出: 2 
解释: 各位相加的过程为：
38 --> 3 + 8 --> 11
11 --> 1 + 1 --> 2
由于 2 是一位数，所以返回 2。
示例 1:

输入: num = 0
输出: 0 

提示：

0 <= num <= 2^31 - 1
`

	url := `
https://leetcode-cn.com/problems/add-digits/
`

	cal := `
func addDigits(num int) int {

}
`

	month := "m202203"

	if err := Gen(desc, url, cal, month); err != nil {
		t.Errorf("Gen error: %+v", err)
	}
}
