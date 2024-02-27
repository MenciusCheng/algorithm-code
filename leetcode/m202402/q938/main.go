package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/range-sum-of-bst/description/

938. 二叉搜索树的范围和
简单
给定二叉搜索树的根结点 root，返回值位于范围 [low, high] 之间的所有结点的值的和。

示例 1：

输入：root = [10,5,15,3,7,null,18], low = 7, high = 15
输出：32
示例 2：

输入：root = [10,5,15,3,7,13,18,1,null,6], low = 6, high = 10
输出：23

提示：

树中节点数目在范围 [1, 2 * 10^4] 内
1 <= Node.val <= 10^5
1 <= low <= high <= 10^5
所有 Node.val 互不相同
*/
func main() {
	var tests = []struct {
		root *TreeNode
		low  int
		high int
		want int
	}{
		{},
	}

	for _, item := range tests {
		if ans := rangeSumBST(item.root, item.low, item.high); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	sum := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Val >= low && node.Val <= high {
			sum += node.Val
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)

	return sum
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
