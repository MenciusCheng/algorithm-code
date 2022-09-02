package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/longest-univalue-path/

687. 最长同值路径
给定一个二叉树的 root ，返回 最长的路径的长度 ，这个路径中的 每个节点具有相同值 。 这条路径可以经过也可以不经过根节点。
两个节点之间的路径长度 由它们之间的边数表示。

示例 1:

输入：root = [5,4,5,1,1,5]
输出：2
示例 2:

输入：root = [1,4,5,4,4,5]
输出：2

提示:

树的节点数的范围是 [0, 10^4]
-1000 <= Node.val <= 1000
树的深度将不超过 1000
*/
func main() {
	var tests = []struct {
		root *TreeNode
		want int
	}{
		{},
	}

	for _, item := range tests {
		if ans := longestUnivaluePath(item.root); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func longestUnivaluePath(root *TreeNode) int {
	var res int

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		c := 0
		l := dfs(node.Left)
		if node.Left != nil && node.Left.Val == node.Val {
			c += l + 1
		}
		r := dfs(node.Right)
		if node.Right != nil && node.Right.Val == node.Val {
			c += r + 1
		}
		//fmt.Printf("v: %d, c: %d\n", node.Val, c)
		res = max(res, c)
		if node.Left != nil && node.Left.Val == node.Val && node.Right != nil && node.Right.Val == node.Val {
			return max(l+1, r+1)
		}
		return c
	}
	dfs(root)

	return res
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
