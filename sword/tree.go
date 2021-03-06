package sword

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中序遍历
// 先遍历二叉树的左子树，然后遍历二叉树的根节点，最后遍历二叉树的右子树。
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

// 前序遍历
// 先遍历二叉树的根节点，再遍历二叉树的左子树，最后遍历二叉树的右子树。
func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root != nil {
		res = append(res, root.Val)
		res = append(res, preorderTraversal(root.Left)...)
		res = append(res, preorderTraversal(root.Right)...)
	}
	return res
}

// 前序遍历，迭代实现
func preorderTraversalFor(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			res = append(res, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return res
}

// 后序遍历
// 先遍历左子树，再遍历右子树，最后遍历根节点。
func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root != nil {
		res = append(res, postorderTraversal(root.Left)...)
		res = append(res, postorderTraversal(root.Right)...)
		res = append(res, root.Val)
	}
	return res
}

// 后序遍历，迭代实现
func postorderTraversalFor(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var prev *TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right == nil || root.Right == prev {
			res = append(res, root.Val)
			prev = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}
	return res
}
