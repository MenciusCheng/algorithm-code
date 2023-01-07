package sword

import "math/rand"

func sortInts(nums []int) {
	quicksort(nums, 0, len(nums)-1)
}

func quicksort(nums []int, start, end int) {
	if start < end {
		pivot := partition(nums, start, end)
		quicksort(nums, start, pivot-1)
		quicksort(nums, pivot+1, end)
	}
}

func partition(nums []int, start, end int) int {
	r := rand.Intn(end-start+1) + start
	nums[r], nums[end] = nums[end], nums[r]

	small := start - 1
	for i := start; i < end; i++ {
		if nums[i] < nums[end] {
			small++
			nums[i], nums[small] = nums[small], nums[i]
		}
	}

	small++
	nums[small], nums[end] = nums[end], nums[small]

	return small
}

func sortInts2(nums []int) {
	dst := make([]int, len(nums))
	copy(dst, nums)
	mergeSort(dst, nums, 0, len(nums)-1)
}

func mergeSort(src, dst []int, start, end int) {
	if start >= end {
		return
	}

	mid := (start + end) >> 1
	mergeSort(dst, src, start, mid)
	mergeSort(dst, src, mid+1, end)

	i := start
	j := mid + 1
	k := start
	for i <= mid || j <= end {
		if j > end || i <= mid && src[i] < src[j] {
			dst[k] = src[i]
			i++
		} else {
			dst[k] = src[j]
			j++
		}
		k++
	}
}
