package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode-cn.com/problems/construct-string-from-binary-tree/

606. 根据二叉树创建字符串
你需要采用前序遍历的方式，将一个二叉树转换成一个由括号和整数组成的字符串。

空节点则用一对空括号 "()" 表示。而且你需要省略所有不影响字符串与原始二叉树之间的一对一映射关系的空括号对。

示例 1:

输入: 二叉树: [1,2,3,4]
       1
     /   \
    2     3
   /
  4

输出: "1(2(4))(3)"

解释: 原本将是“1(2(4)())(3())”，
在你省略所有不必要的空括号对之后，
它将是“1(2(4))(3)”。
示例 2:

输入: 二叉树: [1,2,3,null,4]
       1
     /   \
    2     3
     \
      4

输出: "1(2()(4))(3)"

解释: 和第一个示例相似，
除了我们不能省略第一个对括号来中断输入和输出之间的一对一映射关系。
*/
func main() {
	var tests = []struct {
		root *TreeNode
		want string
	}{
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
				Right: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
			want: "1(2(4))(3)",
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: nil,
					Right: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
			want: "1(2()(4))(3)",
		},
	}

	for _, item := range tests {
		if ans := tree2str(item.root); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func tree2str(root *TreeNode) string {
	if root != nil {
		left := tree2str(root.Left)
		right := tree2str(root.Right)
		if left == "" && right == "" {
			return fmt.Sprintf("%d", root.Val)
		} else if left == "" {
			return fmt.Sprintf("%d()(%s)", root.Val, right)
		} else if right == "" {
			return fmt.Sprintf("%d(%s)", root.Val, left)
		} else {
			return fmt.Sprintf("%d(%s)(%s)", root.Val, left, right)
		}
	}
	return ""
}

func pre(root *TreeNode) string {
	if root != nil {
		left := pre(root.Left)
		right := pre(root.Right)
		if left == "" && right == "" {
			return fmt.Sprintf("%d", root.Val)
		} else if left == "" {
			return fmt.Sprintf("%d()(%s)", root.Val, right)
		} else if right == "" {
			return fmt.Sprintf("%d(%s)", root.Val, left)
		} else {
			return fmt.Sprintf("%d(%s)(%s)", root.Val, left, right)
		}
	}
	return ""
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
