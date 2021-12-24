package main

import (
	"container/heap"
	"fmt"
	"reflect"
	"sort"
)

/*
https://leetcode-cn.com/problems/maximum-number-of-eaten-apples/

1705. 吃苹果的最大数目
有一棵特殊的苹果树，一连 n 天，每天都可以长出若干个苹果。在第 i 天，树上会长出 apples[i] 个苹果，这些苹果将会在 days[i] 天后（也就是说，第 i + days[i] 天时）腐烂，变得无法食用。也可能有那么几天，树上不会长出新的苹果，此时用 apples[i] == 0 且 days[i] == 0 表示。
你打算每天 最多 吃一个苹果来保证营养均衡。注意，你可以在这 n 天之后继续吃苹果。
给你两个长度为 n 的整数数组 days 和 apples ，返回你可以吃掉的苹果的最大数目。

示例 1：

输入：apples = [1,2,3,5,2], days = [3,2,1,4,2]
输出：7
解释：你可以吃掉 7 个苹果：
- 第一天，你吃掉第一天长出来的苹果。
- 第二天，你吃掉一个第二天长出来的苹果。
- 第三天，你吃掉一个第二天长出来的苹果。过了这一天，第三天长出来的苹果就已经腐烂了。
- 第四天到第七天，你吃的都是第四天长出来的苹果。
示例 2：

输入：apples = [3,0,0,0,0,2], days = [3,0,0,0,0,2]
输出：5
解释：你可以吃掉 5 个苹果：
- 第一天到第三天，你吃的都是第一天长出来的苹果。
- 第四天和第五天不吃苹果。
- 第六天和第七天，你吃的都是第六天长出来的苹果。

提示：

apples.length == n
days.length == n
1 <= n <= 2 * 10^4
0 <= apples[i], days[i] <= 2 * 10^4
只有在 apples[i] = 0 时，days[i] = 0 才成立
*/
func main() {
	var tests = []struct {
		apples []int
		days   []int
		want   int
	}{
		{
			apples: []int{1, 2, 3, 5, 2},
			days:   []int{3, 2, 1, 4, 2},
			want:   7,
		},
		{
			apples: []int{3, 0, 0, 0, 0, 2},
			days:   []int{3, 0, 0, 0, 0, 2},
			want:   5,
		},
	}

	for _, item := range tests {
		if ans := eatenApples(item.apples, item.days); reflect.DeepEqual(ans, item.want) {
			fmt.Println(true)
		} else {
			fmt.Printf("ans: %+v, want: %+v\n", ans, item.want)
		}
	}
}

func eatenApples(apples []int, days []int) int {
	countByDayMap := make(map[int]int)
	h := &Heap{}
	sum := 0
	for i := 0; i < len(apples) || h.Len() > 0; i++ {
		if i < len(apples) && apples[i] > 0 {
			d := i + days[i]
			count, ok := countByDayMap[d]
			if !ok || count == 0 {
				heap.Push(h, d)
			}
			countByDayMap[d] = count + apples[i]
		}
		for h.Len() > 0 {
			d := heap.Pop(h).(int)
			if i < d {
				if countByDayMap[d] > 0 {
					sum++
					countByDayMap[d]--
					if countByDayMap[d] > 0 {
						heap.Push(h, d)
					}
					break
				}
			}
		}

	}
	return sum
}

// Heap 优先队列
type Heap struct {
	sort.IntSlice
}

func (h *Heap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *Heap) Pop() interface{} {
	a := h.IntSlice
	x := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return x
}
