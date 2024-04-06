package main

import (
	"fmt"
)

/*
https://leetcode.cn/problems/all-possible-full-binary-trees/?envType=daily-question&envId=2024-04-02

894. 所有可能的真二叉树
中等

给你一个整数 n ，请你找出所有可能含 n 个节点的 真二叉树 ，并以列表形式返回。答案中每棵树的每个节点都必须符合 Node.val == 0 。
答案的每个元素都是一棵真二叉树的根节点。你可以按 任意顺序 返回最终的真二叉树列表。
真二叉树 是一类二叉树，树中每个节点恰好有 0 或 2 个子节点。

示例 1：

输入：n = 7
输出：[[0,0,0,null,null,0,0,null,null,0,0],[0,0,0,null,null,0,0,0,0],[0,0,0,0,0,0,0],[0,0,0,0,0,null,null,null,null,0,0],[0,0,0,0,0,null,null,0,0]]
示例 2：

输入：n = 3
输出：[[0,0,0]]

提示：

1 <= n <= 20
*/
func main() {
	var tests = []struct {
		n    int
		want []string
	}{
		{
			n: 7,
		},
		{
			n: 3,
		},
	}

	for _, item := range tests {
		ans := allPossibleFBT(item.n)
		fmt.Printf("n: %+v, ans: %+v\n", item.n, ans)
	}
}

func allPossibleFBT(n int) []*TreeNode {
	res := make([]*TreeNode, 0)
	if n%2 == 0 {
		return res
	}
	if n == 1 {
		return []*TreeNode{{}}
	}

	for i := 1; i < n; i += 2 {
		lefts := allPossibleFBT(i)
		rights := allPossibleFBT(n - i - 1)
		for _, left := range lefts {
			for _, right := range rights {
				root := &TreeNode{
					Left:  left,
					Right: right,
				}
				res = append(res, root)
			}
		}
	}

	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
