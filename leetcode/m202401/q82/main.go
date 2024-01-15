package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/description/

82. 删除排序链表中的重复元素 II
中等

给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。

示例 1：

输入：head = [1,2,3,3,4,4,5]
输出：[1,2,5]
示例 2：

输入：head = [1,1,1,2,3]
输出：[2,3]

提示：

链表中节点数目在范围 [0, 300] 内
-100 <= Node.val <= 100
题目数据保证链表已经按升序 排列
*/
func main() {
	var tests = []struct {
		head *ListNode
		want *ListNode
	}{
		{},
	}

	for _, item := range tests {
		if ans := deleteDuplicates(item.head); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
	t := &ListNode{
		Next: head,
	}
	res := t
	for t.Next != nil && t.Next.Next != nil {
		if t.Next.Val == t.Next.Next.Val {
			for t.Next != nil && t.Next.Next != nil && t.Next.Val == t.Next.Next.Val {
				t.Next.Next = t.Next.Next.Next
			}
			t.Next = t.Next.Next
		} else {
			t = t.Next
		}
	}

	return res.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
