package main

import (
	"fmt"
	"math/rand"
)

/*
https://leetcode-cn.com/problems/linked-list-random-node/

382. 链表随机节点
给你一个单链表，随机选择链表的一个节点，并返回相应的节点值。每个节点 被选中的概率一样 。
实现 Solution 类：
Solution(ListNode head) 使用整数数组初始化对象。
int getRandom() 从链表中随机选择一个节点并返回该节点的值。链表中所有节点被选中的概率相等。

示例：

输入
["Solution", "getRandom", "getRandom", "getRandom", "getRandom", "getRandom"]
[[[1, 2, 3]], [], [], [], [], []]
输出
[null, 1, 3, 2, 2, 3]

解释
Solution solution = new Solution([1, 2, 3]);
solution.getRandom(); // 返回 1
solution.getRandom(); // 返回 3
solution.getRandom(); // 返回 2
solution.getRandom(); // 返回 2
solution.getRandom(); // 返回 3
// getRandom() 方法应随机返回 1、2、3中的一个，每个元素被返回的概率相等。

提示：

链表中的节点数在范围 [1, 104] 内
-104 <= Node.val <= 104
至多调用 getRandom 方法 104 次
*/
func main() {
	ln := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	c := Constructor(ln)
	fmt.Println(c.GetRandom())
	fmt.Println(c.GetRandom())
	fmt.Println(c.GetRandom())
	fmt.Println(c.GetRandom())
	fmt.Println(c.GetRandom())
	fmt.Println(c.GetRandom())
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
	arr []int
}

func Constructor(head *ListNode) Solution {
	s := Solution{}
	for head != nil {
		s.arr = append(s.arr, head.Val)
		head = head.Next
	}
	return s
}

func (this *Solution) GetRandom() int {
	index := rand.Intn(len(this.arr))
	return this.arr[index]
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
