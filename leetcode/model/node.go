package model

type ListNode struct {
	Val  int
	Next *ListNode
}

func ArrToListNode(arr []int) *ListNode {
	dum := &ListNode{}
	cur := dum
	for _, v := range arr {
		cur.Next = &ListNode{
			Val: v,
		}
		cur = cur.Next
	}
	return dum.Next
}

func ListNodeToArr(head *ListNode) []int {
	cur := head
	res := make([]int, 0)
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}
	return res
}
