package leetcode

import (
	"fmt"
	"testing"
)

func TestGen(t *testing.T) {
	desc := `
1026. 节点与其祖先之间的最大差值
中等
给定二叉树的根节点 root，找出存在于 不同 节点 A 和 B 之间的最大值 V，其中 V = |A.val - B.val|，且 A 是 B 的祖先。
（如果 A 的任何子节点之一为 B，或者 A 的任何子节点是 B 的祖先，那么我们认为 A 是 B 的祖先）

示例 1：

输入：root = [8,3,10,1,6,null,14,null,null,4,7,13]
输出：7
解释： 
我们有大量的节点与其祖先的差值，其中一些如下：
|8 - 3| = 5
|3 - 7| = 4
|8 - 1| = 7
|10 - 13| = 3
在所有可能的差值中，最大值 7 由 |8 - 1| = 7 得出。
示例 2：

输入：root = [1,null,2,null,0,3]
输出：3

提示：

树中的节点数在 2 到 5000 之间。
0 <= Node.val <= 10^5
`

	url := `
https://leetcode.cn/problems/maximum-difference-between-node-and-ancestor/description/?envType=daily-question&envId=2024-04-05
`

	cal := `
func maxAncestorDiff(root *TreeNode) int {

}
`

	month := "m202404"

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
输入：n = 6, edgeList = [[0,3],[5,0],[2,3],[4,3],[5,3],[1,3],[2,5],[0,1],[4,5],[4,2],[4,0],[2,1],[5,1]]
输出：[[2,4,5],[0,2,4,5],[4],[0,1,2,4,5],[],[2,4]]`,
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
