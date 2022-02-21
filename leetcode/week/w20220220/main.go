package w20220220

import (
	"strconv"
)

func countEven(num int) int {
	res := 0
	for i := 1; i <= num; i++ {
		s := strconv.Itoa(i)
		sum := 0
		for j := 0; j < len(s); j++ {
			sum += int(s[j] - '0')
		}
		if sum%2 == 0 {
			res++
		}
	}
	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	res := &ListNode{}
	cNode := res
	isStart := false
	for head != nil {
		if head.Val == 0 {
			isStart = true
		} else {
			if isStart {
				cNode.Next = &ListNode{}
				cNode = cNode.Next
				isStart = false
			}
			cNode.Val += head.Val
		}
		head = head.Next
	}

	return res.Next
}

func repeatLimitedString(s string, repeatLimit int) string {
	cnt := [26]int{}
	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
	}

	top1 := findNext(cnt, 25)
	top2 := findNext(cnt, top1-1)

	sb := make([]byte, 0)
	count := 0
	for top1 > -1 {
		if count < repeatLimit {
			sb = append(sb, byte(top1+'a'))
			cnt[top1]--
			if cnt[top1] == 0 {
				top1 = top2
				top2 = findNext(cnt, top2-1)
				count = 0
			} else {
				count++
			}
		} else if top2 > -1 {
			sb = append(sb, byte(top2+'a'))
			cnt[top2]--
			if cnt[top2] == 0 {
				top2 = findNext(cnt, top2-1)
			}
			count = 0
		} else {
			break
		}
	}
	return string(sb)
}

func findNext(cnt [26]int, start int) int {
	for i := start; i >= 0; i-- {
		if cnt[i] > 0 {
			return i
		}
	}
	return -1
}
