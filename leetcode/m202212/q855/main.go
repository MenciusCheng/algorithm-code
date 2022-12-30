package main

import "fmt"

/*
https://leetcode.cn/problems/exam-room/

855. 考场就座
中等

在考场里，一排有 N 个座位，分别编号为 0, 1, 2, ..., N-1 。

当学生进入考场后，他必须坐在能够使他与离他最近的人之间的距离达到最大化的座位上。如果有多个这样的座位，他会坐在编号最小的座位上。(另外，如果考场里没有人，那么学生就坐在 0 号座位上。)
返回 ExamRoom(int N) 类，它有两个公开的函数：其中，函数 ExamRoom.seat() 会返回一个 int （整型数据），代表学生坐的位置；函数 ExamRoom.leave(int p) 代表坐在座位 p 上的学生现在离开了考场。每次调用 ExamRoom.leave(p) 时都保证有学生坐在座位 p 上。

示例：

输入：["ExamRoom","seat","seat","seat","seat","leave","seat"], [[10],[],[],[],[],[4],[]]
输出：[null,0,9,4,2,null,5]
解释：
ExamRoom(10) -> null
seat() -> 0，没有人在考场里，那么学生坐在 0 号座位上。
seat() -> 9，学生最后坐在 9 号座位上。
seat() -> 4，学生最后坐在 4 号座位上。
seat() -> 2，学生最后坐在 2 号座位上。
leave(4) -> null
seat() -> 5，学生最后坐在 5 号座位上。

提示：

1 <= N <= 10^9
在所有的测试样例中 ExamRoom.seat() 和 ExamRoom.leave() 最多被调用 10^4 次。
保证在调用 ExamRoom.leave(p) 时有学生正坐在座位 p 上。
*/
func main() {
	//examRoom := Constructor(10)
	//fmt.Printf("seat: %d\n", examRoom.Seat())
	//fmt.Printf("seat: %d\n", examRoom.Seat())
	//fmt.Printf("seat: %d\n", examRoom.Seat())
	//fmt.Printf("seat: %d\n", examRoom.Seat())
	//examRoom.Leave(4)
	//fmt.Printf("Leave\n")
	//fmt.Printf("seat: %d\n", examRoom.Seat())

	examRoom := Constructor(10)
	fmt.Printf("seat: %d\n", examRoom.Seat())
	fmt.Printf("seat: %d\n", examRoom.Seat())
	fmt.Printf("seat: %d\n", examRoom.Seat())
	examRoom.Leave(0)
	examRoom.Leave(4)
	fmt.Printf("seat: %d\n", examRoom.Seat())
	fmt.Printf("seat: %d\n", examRoom.Seat())
	fmt.Printf("seat: %d\n", examRoom.Seat())
	fmt.Printf("seat: %d\n", examRoom.Seat())
	fmt.Printf("seat: %d\n", examRoom.Seat())
	fmt.Printf("seat: %d\n", examRoom.Seat())
	fmt.Printf("seat: %d\n", examRoom.Seat())
	fmt.Printf("seat: %d\n", examRoom.Seat())
	fmt.Printf("seat: %d\n", examRoom.Seat())
	examRoom.Leave(0)
}

type ExamRoom struct {
	Size int
	Head *DListNode
	Cnt  map[int]*DListNode
}

type DListNode struct {
	Val  int
	Last *DListNode
	Next *DListNode
}

func Constructor(n int) ExamRoom {
	return ExamRoom{
		Size: n,
		Cnt:  make(map[int]*DListNode),
	}
}

func (this *ExamRoom) Seat() int {
	if this.Head == nil {
		node := &DListNode{
			Val: 0,
		}
		this.Head = node
		this.Cnt[node.Val] = node
		return node.Val
	}

	p := this.Head
	maxD := p.Val
	var pd *DListNode
	for p != nil {
		var d int
		if p.Next != nil {
			d = (p.Next.Val - p.Val) / 2
		} else {
			d = this.Size - 1 - p.Val
		}
		if d > maxD {
			maxD = d
			pd = p
		}
		p = p.Next
	}
	if pd == nil {
		if this.Head.Val == 0 {
			return 0
		}
		node := &DListNode{
			Val:  0,
			Next: this.Head,
		}
		this.Head.Last = node
		this.Head = node
		this.Cnt[node.Val] = node
		return node.Val
	} else {
		node := &DListNode{
			Val:  pd.Val + maxD,
			Last: pd,
			Next: pd.Next,
		}
		if pd.Next != nil {
			pd.Next.Last = node
		}
		pd.Next = node
		this.Cnt[node.Val] = node
		return node.Val
	}
}

func (this *ExamRoom) Leave(p int) {
	if this.Cnt[p] == this.Head {
		this.Head = this.Cnt[p].Next
	}
	if this.Cnt[p].Last != nil {
		this.Cnt[p].Last.Next = this.Cnt[p].Next
	}
	if this.Cnt[p].Next != nil {
		this.Cnt[p].Next.Last = this.Cnt[p].Last
	}
	delete(this.Cnt, p)
}

/**
 * Your ExamRoom object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Seat();
 * obj.Leave(p);
 */
