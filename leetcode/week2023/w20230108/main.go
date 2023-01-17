package w20230108

import (
	"container/heap"
	"sort"
)

func maxKelements(nums []int, k int) int64 {
	h := &MaxHeap{}
	for _, num := range nums {
		heap.Push(h, num)
	}
	var res int64
	for i := 0; i < k; i++ {
		v := heap.Pop(h).(int)
		res += int64(v)
		v2 := v / 3
		if v%3 > 0 {
			v2++
		}
		heap.Push(h, v2)
	}
	return res
}

type MaxHeap struct {
	sort.IntSlice
}

func (h *MaxHeap) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

func (h *MaxHeap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	arr := h.IntSlice
	x := arr[len(arr)-1]
	h.IntSlice = arr[:len(arr)-1]
	return x
}

func findCrossingTime(n int, k int, time [][]int) int {
	for i := range time {
		time[i] = append(time[i], i)
	}
	sort.Slice(time, func(i, j int) bool {
		return (time[i][0]+time[i][2]) > (time[j][0]+time[j][2]) ||
			(time[i][0]+time[i][2]) == (time[j][0]+time[j][2]) && time[i][4] > time[j][4]
	})
	for i := range time {
		time[i] = append(time[i], i, 0)
		workEf[time[i][4]] = time[i] // = []int{leftToRighti, pickOldi, rightToLefti, putNewi, wi, effective, nextT}
	}

	leftWait := &Heap{}
	for i := 0; i < k; i++ {
		heap.Push(leftWait, i)
	}
	rightWait := &Heap{}
	timeHeap := &TimeHeap{}

	workStatus := make(map[int]int)
	qiao := false

	t := 0
	res := 0
	for n > 0 || timeHeap.Len() > 0 {
		if timeHeap.Len() > 0 {
			t = workEf[timeHeap.IntSlice[0]][6]
		}
		for timeHeap.Len() > 0 && workEf[timeHeap.IntSlice[0]][6] <= t {
			w := heap.Pop(timeHeap).(int)
			switch workStatus[w] {
			case 1:
				qiao = false
				workStatus[w] = 2
				workEf[w][6] = t + workEf[w][1]
				heap.Push(timeHeap, w)
			case 2:
				workStatus[w] = 0
				heap.Push(rightWait, w)
			case 3:
				res = max(res, t)
				qiao = false
				workStatus[w] = 4
				workEf[w][6] = t + workEf[w][3]
				heap.Push(timeHeap, w)
			case 4:
				workStatus[w] = 0
				heap.Push(leftWait, w)
			}
		}

		if !qiao && rightWait.Len() > 0 {
			w := heap.Pop(rightWait).(int)
			workStatus[w] = 3
			qiao = true

			workEf[w][6] = t + workEf[w][2]
			heap.Push(timeHeap, w)
		}
		if !qiao && leftWait.Len() > 0 && n > 0 {
			n--
			w := heap.Pop(leftWait).(int)
			workStatus[w] = 1
			qiao = true

			workEf[w][6] = t + workEf[w][0]
			heap.Push(timeHeap, w)
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var workEf = make(map[int][]int)

type Heap struct {
	sort.IntSlice
}

func (h *Heap) Less(i, j int) bool {
	return workEf[h.IntSlice[i]][5] < workEf[h.IntSlice[j]][5]
}

func (h *Heap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *Heap) Pop() interface{} {
	arr := h.IntSlice
	x := arr[len(arr)-1]
	h.IntSlice = arr[:len(arr)-1]
	return x
}

type TimeHeap struct {
	sort.IntSlice
}

func (h *TimeHeap) Less(i, j int) bool {
	return workEf[h.IntSlice[i]][6] < workEf[h.IntSlice[j]][6]
}

func (h *TimeHeap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *TimeHeap) Pop() interface{} {
	arr := h.IntSlice
	x := arr[len(arr)-1]
	h.IntSlice = arr[:len(arr)-1]
	return x
}
