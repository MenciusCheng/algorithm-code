package main

import (
	"fmt"
	"math"
	"reflect"
)

/*
https://leetcode.cn/problems/print-binary-tree/

655. 输出二叉树
给你一棵二叉树的根节点 root ，请你构造一个下标从 0 开始、大小为 m x n 的字符串矩阵 res ，用以表示树的 格式化布局 。构造此格式化布局矩阵需要遵循以下规则：

树的 高度 为 height ，矩阵的行数 m 应该等于 height + 1 。
矩阵的列数 n 应该等于 2^(height+1) - 1 。
根节点 需要放置在 顶行 的 正中间 ，对应位置为 res[0][(n-1)/2] 。
对于放置在矩阵中的每个节点，设对应位置为 res[r][c] ，将其左子节点放置在 res[r+1][c-2^(height-r-1)] ，右子节点放置在 res[r+1][c+2^(height-r-1)] 。
继续这一过程，直到树中的所有节点都妥善放置。
任意空单元格都应该包含空字符串 "" 。
返回构造得到的矩阵 res 。

示例 1：


输入：root = [1,2]
输出：
[["","1",""],
 ["2","",""]]
示例 2：


输入：root = [1,2,3,null,4]
输出：
[["","","","1","","",""],
 ["","2","","","","3",""],
 ["","","4","","","",""]]

提示：

树中节点数在范围 [1, 2^10] 内
-99 <= Node.val <= 99
树的深度在范围 [1, 10] 内
*/
func main() {
	var tests = []struct {
		root *TreeNode
		want [][]string
	}{
		{},
	}

	for _, item := range tests {
		if ans := printTree(item.root); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func printTree(root *TreeNode) [][]string {
	height := calHeight(root)
	m := height + 1
	n := int(math.Pow(2, float64(height+1))) - 1
	res := make([][]string, m)
	for i := range res {
		res[i] = make([]string, n)
	}
	fillRes(res, root, height, 0, (n-1)/2)
	return res
}

func calHeight(root *TreeNode) int {
	if root == nil {
		return -1
	}
	return max(calHeight(root.Left)+1, calHeight(root.Right)+1)
}

func fillRes(res [][]string, root *TreeNode, height, r, c int) {
	if root == nil {
		return
	}
	res[r][c] = fmt.Sprintf("%d", root.Val)
	fillRes(res, root.Left, height, r+1, c-int(math.Pow(2, float64(height-r-1))))
	fillRes(res, root.Right, height, r+1, c+int(math.Pow(2, float64(height-r-1))))
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
