package w20220123

import "sort"

func countElements(nums []int) int {
	if len(nums) <= 2 {
		return 0
	}
	sort.Ints(nums)
	count := len(nums)

	min := nums[0]
	max := nums[len(nums)-1]
	if min == max {
		return 0
	}

	for i := 0; i < len(nums) && nums[i] == min; i++ {
		count--
	}
	for i := len(nums) - 1; i >= 0 && nums[i] == max; i-- {
		count--
	}
	return count
}

func rearrangeArray(nums []int) []int {
	res := make([]int, len(nums))
	p1 := 0
	p2 := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] >= 0 {
			res[p1] = nums[i]
			p1 += 2
		} else {
			res[p2] = nums[i]
			p2 += 2
		}
	}
	return res
}

func findLonely(nums []int) []int {
	sort.Ints(nums)
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if (i == 0 || nums[i-1] != nums[i] && nums[i-1]+1 != nums[i]) && (i == len(nums)-1 || nums[i] != nums[i+1] && nums[i]+1 != nums[i+1]) {
			res = append(res, nums[i])
		}
	}
	return res
}
