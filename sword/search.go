package sword

import "sort"

func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	// 如果找不到目标target值，执行到这里时是 right < left
	// 想要最大比target小的值，则返回 right
	// 想要最小比target大的值，则返回 left
	return right
}

func binarySearch3(ws []int, num int) int {
	return sort.SearchInts(ws, num)
}
