package w20220319

import "sort"

func divideArray(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		if nums[i] != nums[i+1] {
			return false
		}
	}
	return true
}

func maximumSubsequenceCount(text string, pattern string) int64 {
	a := pattern[0]
	b := pattern[1]
	if a == b {
		var sum, aCount int64
		for i := len(text) - 1; i >= 0; i-- {
			if text[i] == a {
				sum += aCount
				aCount++
			}
		}
		sum += aCount
		return sum
	} else {
		var sum1, sum2 int64
		bCount1 := int64(1)
		bCount2 := int64(0)
		for i := len(text) - 1; i >= 0; i-- {
			if text[i] == b {
				bCount1++
				bCount2++
			} else if text[i] == a {
				sum1 += bCount1
				sum2 += bCount2
			}
		}
		sum2 += bCount2
		if sum1 > sum2 {
			return sum1
		} else {
			return sum2
		}
	}
}
