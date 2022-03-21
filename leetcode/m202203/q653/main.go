package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode-cn.com/problems/two-sum-iv-input-is-a-bst/

653. 两数之和 IV - 输入 BST
给定一个二叉搜索树 root 和一个目标结果 k，如果 BST 中存在两个元素且它们的和等于给定的目标结果，则返回 true。

示例 1：

输入: root = [5,3,6,2,4,null,7], k = 9
输出: true
示例 2：

输入: root = [5,3,6,2,4,null,7], k = 28
输出: false

提示:

二叉树的节点个数的范围是  [1, 10^4].
-10^4 <= Node.val <= 10^4
root 为二叉搜索树
-10^5 <= k <= 10^5
*/
func main() {
	var tests = []struct {
		root *TreeNode
		k    int
		want bool
	}{
		{
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val:  6,
					Left: nil,
					Right: &TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
				},
			},
			k:    9,
			want: true,
		},
	}

	for _, item := range tests {
		if ans := findTarget(item.root, item.k); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func findTarget(root *TreeNode, k int) bool {
	cnt := make(map[int]int)
	dfs(root, cnt)
	return dfsTarget(root, cnt, k)
}

func dfs(root *TreeNode, cnt map[int]int) {
	if root != nil {
		cnt[root.Val]++
		dfs(root.Left, cnt)
		dfs(root.Right, cnt)
	}
}

func dfsTarget(root *TreeNode, cnt map[int]int, k int) bool {
	if root != nil {
		if cnt[k-root.Val] > 0 && (k-root.Val != root.Val) {
			return true
		}
		if dfsTarget(root.Left, cnt, k) {
			return true
		}
		if dfsTarget(root.Right, cnt, k) {
			return true
		}
	}
	return false
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
