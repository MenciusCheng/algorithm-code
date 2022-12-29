package sword

// 中序遍历，递归版
func dfsInOrder(root *TreeNode, f func(v int)) {
	if root != nil {
		dfsInOrder(root.Left, f)
		f(root.Val)
		dfsInOrder(root.Right, f)
	}
}

// 中序遍历，循环版
func dfsInOrderFor(root *TreeNode, f func(v int)) {
	stack := make([]*TreeNode, 0)
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		f(node.Val)
		node = node.Right
	}
}

// 前序遍历，递归版
func dfsPreOrder(root *TreeNode, f func(v int)) {
	if root != nil {
		f(root.Val)
		dfsPreOrder(root.Left, f)
		dfsPreOrder(root.Right, f)
	}
}

// 前序遍历，循环版
func dfsPreOrderFor(root *TreeNode, f func(v int)) {
	stack := make([]*TreeNode, 0)
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			f(node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node = node.Right
	}
}

// 后序遍历，递归版
func dfsPostOrder(root *TreeNode, f func(v int)) {
	if root != nil {
		dfsPostOrder(root.Left, f)
		dfsPostOrder(root.Right, f)
		f(root.Val)
	}
}

// 后序遍历，循环版
func dfsPostOrderFor(root *TreeNode, f func(v int)) {
	stack := make([]*TreeNode, 0)
	node := root
	var prev *TreeNode
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		if node.Right != nil && node.Right != prev {
			node = node.Right
		} else {
			stack = stack[:len(stack)-1]
			f(node.Val)
			prev = node
			node = nil
		}
	}
}
