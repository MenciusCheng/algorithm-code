package leetcode

import (
	"fmt"
	"testing"
)

func TestGen(t *testing.T) {
	desc := `
796. 旋转字符串
给定两个字符串, s 和 goal。如果在若干次旋转操作之后，s 能变成 goal ，那么返回 true 。
s 的 旋转操作 就是将 s 最左边的字符移动到最右边。
例如, 若 s = 'abcde'，在旋转一次之后结果就是'bcdea' 。

示例 1:

输入: s = "abcde", goal = "cdeab"
输出: true
示例 2:

输入: s = "abcde", goal = "abced"
输出: false

提示:

1 <= s.length, goal.length <= 100
s 和 goal 由小写英文字母组成
`

	url := `
https://leetcode-cn.com/problems/rotate-string/
`

	cal := `
func rotateString(s string, goal string) bool {

}
`

	month := "m202204"

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
7
[[0,1],[1,2],[1,3],[2,4],[3,5],[4,6]]
`,
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
