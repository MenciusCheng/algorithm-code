package sword

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中序遍历
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root != nil {
		res = append(res, inorderTraversal(root.Left)...)
		res = append(res, root.Val)
		res = append(res, inorderTraversal(root.Right)...)
	}
	return res
}

// 中序遍历，迭代实现版本
func inorderTraversalFor(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, cur.Val)
		cur = cur.Right
	}

	return res
}
