package main

import (
	"fmt"
)

/*
https://leetcode-cn.com/problems/binary-tree-tilt/

563. 二叉树的坡度
给定一个二叉树，计算 整个树 的坡度 。

一个树的 节点的坡度 定义即为，该节点左子树的节点之和和右子树节点之和的 差的绝对值 。如果没有左子树的话，左子树的节点之和为 0 ；没有右子树的话也是一样。空结点的坡度是 0 。

整个树 的坡度就是其所有节点的坡度之和。
*/
func main() {
	/*
		输入：root = [1,2,3]
		输出：1

		输入：root = [4,2,9,3,5,null,7]
		输出：15
	*/
	root := makeTreeNode([]int{1, 2, 3})
	fmt.Println(findTilt(root) == 1)

	root = makeTreeNode([]int{4, 2, 9, 3, 5, 10000, 7})
	fmt.Println(findTilt(root) == 15)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	pL := findTilt(root.Left)
	pR := findTilt(root.Right)

	left := 0
	if root.Left != nil {
		left = root.Left.Val
	}
	right := 0
	if root.Right != nil {
		right = root.Right.Val
	}
	p := abs(left, right)
	root.Val += left + right

	return pL + pR + p
}

func abs(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func makeTreeNode(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}

	treeNodes := make([]*TreeNode, len(arr))
	for i := 0; i < len(arr); i++ {
		if -1000 <= arr[i] && arr[i] <= 1000 {
			treeNodes[i] = &TreeNode{Val: arr[i]}
		}
	}
	for i, node := range treeNodes {
		if node != nil {
			leftIndex := (i+1)*2 - 1
			if leftIndex >= len(treeNodes) {
				break
			}
			node.Left = treeNodes[leftIndex]

			rightIndex := leftIndex + 1
			if rightIndex >= len(treeNodes) {
				break
			}
			node.Right = treeNodes[rightIndex]
		}
	}

	return treeNodes[0]
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	printTree(root.Left)
	fmt.Printf("%d ", root.Val)
	printTree(root.Right)
}
