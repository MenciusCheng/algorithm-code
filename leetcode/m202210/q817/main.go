package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.cn/problems/linked-list-components/

817. 链表组件
中等

给定链表头结点 head，该链表上的每个结点都有一个 唯一的整型值 。同时给定列表 nums，该列表是上述链表中整型值的一个子集。

返回列表 nums 中组件的个数，这里对组件的定义为：链表中一段最长连续结点的值（该值必须在列表 nums 中）构成的集合。

示例 1：

输入: head = [0,1,2,3], nums = [0,1,3]
输出: 2
解释: 链表中,0 和 1 是相连接的，且 nums 中不包含 2，所以 [0, 1] 是 nums 的一个组件，同理 [3] 也是一个组件，故返回 2。
示例 2：

输入: head = [0,1,2,3,4], nums = [0,3,1,4]
输出: 2
解释: 链表中，0 和 1 是相连接的，3 和 4 是相连接的，所以 [0, 1] 和 [3, 4] 是两个组件，故返回 2。

提示：

链表中节点数为n
1 <= n <= 10^4
0 <= Node.val < n
Node.val 中所有值 不同
1 <= nums.length <= n
0 <= nums[i] < n
nums 中所有值 不同
*/
func main() {
	var tests = []struct {
		head *ListNode
		nums []int
		want int
	}{
		{
			head: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
			},
			nums: []int{0, 1, 3},
			want: 2,
		},
		{
			head: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val:  4,
								Next: nil,
							},
						},
					},
				},
			},
			nums: []int{0, 3, 1, 4},
			want: 2,
		},
	}

	for _, item := range tests {
		if ans := numComponents(item.head, item.nums); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func numComponents(head *ListNode, nums []int) int {
	cnt := make(map[int]bool)
	for _, num := range nums {
		cnt[num] = true
	}

	flag := false
	res := 0
	for head != nil {
		if cnt[head.Val] {
			if !flag {
				flag = true
				res++
			}
		} else {
			flag = false
		}
		head = head.Next
	}

	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}
