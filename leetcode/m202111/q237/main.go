package main

import "fmt"

// https://leetcode-cn.com/problems/delete-node-in-a-linked-list/
// 237. 删除链表中的节点
func main() {
	headNode, targetNode := makeListNode([]int{4, 5, 1, 9}, 5)
	showListNode(headNode)
	showListNode(targetNode)

	deleteNode(targetNode)
	showListNode(headNode)
}

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func makeListNode(arr []int, target int) (headNode *ListNode, targetNode *ListNode) {
	var nextNode *ListNode
	for i := len(arr) - 1; i >= 0; i-- {
		node := &ListNode{
			Val: arr[i],
		}
		if nextNode != nil {
			node.Next = nextNode
		}
		nextNode = node

		if arr[i] == target {
			targetNode = node
		}
		if i == 0 {
			headNode = node
		}
	}
	return
}

func showListNode(node *ListNode) {
	fmt.Printf("[")
	for node != nil {
		fmt.Printf("%d ", node.Val)
		node = node.Next
	}
	fmt.Println("]")
}

type ListNode struct {
	Val  int
	Next *ListNode
}
