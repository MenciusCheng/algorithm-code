package w20230205

import (
	"container/heap"
	"math"
	"sort"
)

func pickGifts(gifts []int, k int) int64 {
	h := &Heap{}
	for _, gift := range gifts {
		heap.Push(h, gift)
	}
	for i := 0; i < k; i++ {
		x := heap.Pop(h).(int)
		v := int(math.Sqrt(float64(x)))
		heap.Push(h, v)
	}
	var sum int64
	for _, v := range h.IntSlice {
		sum += int64(v)
	}
	return sum
}

type Heap struct {
	sort.IntSlice
}

func (h *Heap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *Heap) Pop() interface{} {
	x := h.IntSlice[h.Len()-1]
	h.IntSlice = h.IntSlice[:h.Len()-1]
	return x
}

func (h *Heap) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

func vowelStrings(words []string, queries [][]int) []int {
	sums := make([]int, 0)
	sums = append(sums, 0)
	for i := 0; i < len(words); i++ {
		sums = append(sums, sums[i]+isVowel(words[i]))
	}

	res := make([]int, 0)
	for _, q := range queries {
		res = append(res, sums[q[1]+1]-sums[q[0]])
	}
	return res
}

var vmap map[byte]bool = map[byte]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
}

func isVowel(word string) int {
	if vmap[word[0]] && vmap[word[len(word)-1]] {
		return 1
	}
	return 0
}

func minCapability(nums []int, k int) int {
	dp := make([]int, len(nums)+2)
	for i := 0; i < k; i++ {
		dp2 := make([]int, len(nums)+2)
		dp2[i*2-1+2] = math.MaxInt
		for j := i * 2; j < len(nums); j++ {
			dp2[j+2] = min(dp2[j+1], max(nums[j], dp[j]))
		}
		dp = dp2
	}
	return dp[len(dp)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
