package sword

import "sort"

// 使用内置结构定义优先队列

type Heap struct {
	sort.IntSlice
}

func (h *Heap) Less(i, j int) bool {
	return h.IntSlice[i] < h.IntSlice[j]
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

// 自定义优先队列

type MinHeap struct {
	arr []int
}

func (h *MinHeap) Len() int {
	return len(h.arr)
}

func (h *MinHeap) Less(i, j int) bool {
	return h.arr[i] < h.arr[j]
}

func (h *MinHeap) Swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h *MinHeap) Push(x interface{}) {
	h.arr = append(h.arr, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	x := h.arr[h.Len()-1]
	h.arr = h.arr[:h.Len()-1]
	return x
}

// 继承优先队列

type MaxHeap struct {
	MinHeap
}

func (h *MaxHeap) Less(i, j int) bool {
	return h.arr[i] > h.arr[j]
}
