package leetcode

import (
	"fmt"
	"testing"
)

func TestGen(t *testing.T) {
	desc := `
2645. 构造有效字符串的最少插入数
中等

给你一个字符串 word ，你可以向其中任何位置插入 "a"、"b" 或 "c" 任意次，返回使 word 有效 需要插入的最少字母数。
如果字符串可以由 "abc" 串联多次得到，则认为该字符串 有效 。

示例 1：

输入：word = "b"
输出：2
解释：在 "b" 之前插入 "a" ，在 "b" 之后插入 "c" 可以得到有效字符串 "abc" 。
示例 2：

输入：word = "aaa"
输出：6
解释：在每个 "a" 之后依次插入 "b" 和 "c" 可以得到有效字符串 "abcabcabc" 。
示例 3：

输入：word = "abc"
输出：0
解释：word 已经是有效字符串，不需要进行修改。 

提示：

1 <= word.length <= 50
word 仅由字母 "a"、"b" 和 "c" 组成。
`

	url := `
https://leetcode.cn/problems/minimum-additions-to-make-valid-string/description/
`

	cal := `
func addMinimum(word string) int {

}
`

	month := "m202401"

	if err := Gen(desc, url, cal, month); err != nil {
		t.Errorf("Gen error: %+v", err)
	}
}

func TestArrStr(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				str: `
[[24,11,22,17,4],[21,16,5,12,9],[6,23,10,3,18],[15,20,1,8,13],[0,7,14,19,2]]`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ArrStr(tt.args.str)
			fmt.Println(got)
		})
	}
}
