package main

import (
	"fmt"
	"github.com/MenciusCheng/algorithm-code/leetcode/model"
	"reflect"
)

/*
https://leetcode.cn/problems/remove-zero-sum-consecutive-nodes-from-linked-list/description/

1171. 从链表中删去总和值为零的连续节点
中等
给你一个链表的头节点 head，请你编写代码，反复删去链表中由 总和 值为 0 的连续节点组成的序列，直到不存在这样的序列为止。

删除完毕后，请你返回最终结果链表的头节点。

你可以返回任何满足题目要求的答案。

（注意，下面示例中的所有序列，都是对 ListNode 对象序列化的表示。）

示例 1：

输入：head = [1,2,-3,3,1]
输出：[3,1]
提示：答案 [1,2,1] 也是正确的。
示例 2：

输入：head = [1,2,3,-3,4]
输出：[1,2,4]
示例 3：

输入：head = [1,2,3,-3,-2]
输出：[1]

提示：

给你的链表中可能有 1 到 1000 个节点。
对于链表中的每个节点，节点的值：-1000 <= node.val <= 1000.
*/
func main() {
	var tests = []struct {
		head []int
		want []int
	}{
		{
			head: []int{1, 2, -3, 3, 1},
			want: []int{3, 1},
		},
		{
			head: []int{1, 2, 3, -3, 4},
			want: []int{1, 2, 4},
		},
		{
			head: []int{1, 2, 3, -3, -2},
			want: []int{1},
		},
		{
			head: []int{1, -1},
			want: []int{},
		},
	}

	for _, item := range tests {
		if ans := model.ListNodeToArr(removeZeroSumSublists(model.ArrToListNode(item.head))); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

type ListNode = model.ListNode

func removeZeroSumSublists(head *ListNode) *ListNode {
	dum := &ListNode{Next: head}
	cnt := make(map[int]*ListNode)
	sum := 0
	for dum.Next != nil {
		dum = dum.Next
		sum += dum.Val
		cnt[sum] = dum
	}

	if node, ok := cnt[0]; ok {
		head = node.Next
	}
	dum = &ListNode{Next: head}
	sum = 0
	for dum.Next != nil {
		dum = dum.Next
		sum += dum.Val
		if node, ok := cnt[sum]; ok {
			dum.Next = node.Next
		}
	}
	return head
}
