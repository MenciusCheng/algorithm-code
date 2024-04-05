package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/maximum-difference-between-node-and-ancestor/description/?envType=daily-question&envId=2024-04-05

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
*/
func main() {
	var tests = []struct {
		root *TreeNode
		want int
	}{
		{},
	}

	for _, item := range tests {
		if ans := maxAncestorDiff(item.root); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func maxAncestorDiff(root *TreeNode) int {
	res := 0
	var dfs func(node *TreeNode, minV, maxV int)
	dfs = func(node *TreeNode, minV, maxV int) {
		if node == nil {
			return
		}
		minV = min(minV, node.Val)
		maxV = max(maxV, node.Val)
		res = max(res, maxV-minV)
		dfs(node.Left, minV, maxV)
		dfs(node.Right, minV, maxV)
	}

	dfs(root, root.Val, root.Val)

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
