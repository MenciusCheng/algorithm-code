package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode-cn.com/problems/online-election/

911. 在线选举
给你两个整数数组 persons 和 times 。在选举中，第 i 张票是在时刻为 times[i] 时投给候选人 persons[i] 的。
对于发生在时刻 t 的每个查询，需要找出在 t 时刻在选举中领先的候选人的编号。
在 t 时刻投出的选票也将被计入我们的查询之中。在平局的情况下，最近获得投票的候选人将会获胜。
实现 TopVotedCandidate 类：

TopVotedCandidate(int[] persons, int[] times) 使用 persons 和 times 数组初始化对象。
int q(int t) 根据前面描述的规则，返回在时刻 t 在选举中领先的候选人的编号。

示例：

输入：
["TopVotedCandidate", "q", "q", "q", "q", "q", "q"]
[[[0, 1, 1, 0, 0, 1, 0], [0, 5, 10, 15, 20, 25, 30]], [3], [12], [25], [15], [24], [8]]
输出：
[null, 0, 1, 1, 0, 0, 1]

解释：
TopVotedCandidate topVotedCandidate = new TopVotedCandidate([0, 1, 1, 0, 0, 1, 0], [0, 5, 10, 15, 20, 25, 30]);
topVotedCandidate.q(3); // 返回 0 ，在时刻 3 ，票数分布为 [0] ，编号为 0 的候选人领先。
topVotedCandidate.q(12); // 返回 1 ，在时刻 12 ，票数分布为 [0,1,1] ，编号为 1 的候选人领先。
topVotedCandidate.q(25); // 返回 1 ，在时刻 25 ，票数分布为 [0,1,1,0,0,1] ，编号为 1 的候选人领先。（在平局的情况下，1 是最近获得投票的候选人）。
topVotedCandidate.q(15); // 返回 0
topVotedCandidate.q(24); // 返回 0
topVotedCandidate.q(8); // 返回 1

提示：

1 <= persons.length <= 5000
times.length == persons.length
0 <= persons[i] < persons.length
0 <= times[i] <= 10^9
times 是一个严格递增的有序数组
times[0] <= t <= 10^9
每个测试用例最多调用 10^4 次 q
*/
func main() {
	test2()
}

func test1() {
	topVotedCandidate := Constructor([]int{0, 1, 1, 0, 0, 1, 0}, []int{0, 5, 10, 15, 20, 25, 30})

	var tests = []struct {
		t    int
		want int
	}{
		{
			t:    3,
			want: 0,
		},
		{
			t:    12,
			want: 1,
		},
		{
			t:    25,
			want: 1,
		},
		{
			t:    15,
			want: 0,
		},
		{
			t:    24,
			want: 0,
		},
		{
			t:    8,
			want: 1,
		},
	}

	for _, item := range tests {
		if ans := topVotedCandidate.Q(item.t); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func test2() {
	topVotedCandidate := Constructor([]int{0, 1, 0, 1, 1}, []int{24, 29, 31, 76, 81})

	var tests = []struct {
		t    int
		want int
	}{
		{
			t:    28,
			want: 0,
		},
		{
			t:    24,
			want: 0,
		},
		{
			t:    29,
			want: 1,
		},
		{
			t:    77,
			want: 1,
		},
		{
			t:    30,
			want: 1,
		},
		{
			t:    25,
			want: 0,
		},
		{
			t:    76,
			want: 1,
		},
		{
			t:    75,
			want: 0,
		},
		{
			t:    81,
			want: 1,
		},
		{
			t:    80,
			want: 1,
		},
	}

	for _, item := range tests {
		if ans := topVotedCandidate.Q(item.t); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

type TopVotedCandidate struct {
	persons    []int
	times      []int
	topPersons []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	t := TopVotedCandidate{
		persons: persons,
		times:   times,
	}
	max := 0
	maxPerson := 0
	cnt := make(map[int]int)
	for _, person := range persons {
		cnt[person]++
		if cnt[person] >= max {
			max = cnt[person]
			maxPerson = person
		}
		t.topPersons = append(t.topPersons, maxPerson)
	}
	return t
}

func (this *TopVotedCandidate) Q(t int) int {
	var mid int
	head, tail := 0, len(this.times)-1
	for head <= tail {
		mid = (head + tail) / 2

		if this.times[mid] == t {
			break
		} else if this.times[mid] < t {
			head = mid + 1
		} else {
			tail = mid - 1
			if mid > 0 {
				mid -= 1
			}
		}
	}

	return this.topPersons[mid]
}

/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * obj := Constructor(persons, times);
 * param_1 := obj.Q(t);
 */
