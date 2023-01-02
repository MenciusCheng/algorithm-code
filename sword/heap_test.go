package sword

import (
	"container/heap"
	"reflect"
	"testing"
)

func Test_Heap(t *testing.T) {
	h := &Heap{}

	for i := 0; i < 10; i++ {
		heap.Push(h, i)
	}
	for i := 0; i < 10; i++ {
		got := heap.Pop(h).(int)
		if !reflect.DeepEqual(got, i) {
			t.Errorf("MinHeap got = %v, want %v", got, i)
		}
	}
}

func Test_MinHeap(t *testing.T) {
	h := &MinHeap{}

	for i := 0; i < 10; i++ {
		heap.Push(h, i)
	}
	for i := 0; i < 10; i++ {
		got := heap.Pop(h).(int)
		if !reflect.DeepEqual(got, i) {
			t.Errorf("MinHeap got = %v, want %v", got, i)
		}
	}
}

func Test_MaxHeap(t *testing.T) {
	h := &MaxHeap{}

	for i := 0; i < 10; i++ {
		heap.Push(h, i)
	}
	for i := 9; i >= 0; i-- {
		got := heap.Pop(h).(int)
		if !reflect.DeepEqual(got, i) {
			t.Errorf("MinHeap got = %v, want %v", got, i)
		}
	}
}
