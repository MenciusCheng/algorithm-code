package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/find-duplicate-subtrees/

652. 寻找重复的子树
给定一棵二叉树 root，返回所有重复的子树。
对于同一类的重复子树，你只需要返回其中任意一棵的根结点即可。
如果两棵树具有相同的结构和相同的结点值，则它们是重复的。

示例 1：

输入：root = [1,2,3,4,null,2,4,null,null,4]
输出：[[2,4],[4]]
示例 2：

输入：root = [2,1,1]
输出：[[1]]
示例 3：

输入：root = [2,2,2,3,null,3,null]
输出：[[2,3],[3]]

提示：

树中的结点数在[1,10^4]范围内。
-200 <= Node.val <= 200
*/
func main() {
	var tests = []struct {
		root *TreeNode
		want []*TreeNode
	}{
		{},
	}

	for _, item := range tests {
		if ans := findDuplicateSubtrees(item.root); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	repeat := make(map[string]*TreeNode)
	m := make(map[string]*TreeNode)
	var dfs func(node *TreeNode) string
	dfs = func(node *TreeNode) string {
		if node == nil {
			return ""
		}
		s := fmt.Sprintf("%d(%s)(%s)", node.Val, dfs(node.Left), dfs(node.Right))
		if _, ok := m[s]; !ok {
			m[s] = node
		} else {
			repeat[s] = node
		}
		return s
	}
	dfs(root)

	res := make([]*TreeNode, 0)
	for _, node := range repeat {
		res = append(res, node)
	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
