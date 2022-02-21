package w20220219

import "sort"

func countPairs(nums []int, k int) int {
	m := make(map[int][]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = append(m[nums[i]], i)
	}

	res := 0
	for _, v := range m {
		res += countP(v, k)
	}
	return res
}

func countP(nums []int, k int) int {
	if len(nums) < 2 {
		return 0
	}
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if (nums[i]*nums[j])%k == 0 {
				count++
			}
		}
	}
	return count
}

func sumOfThree(num int64) []int64 {
	if num%3 != 0 {
		return []int64{}
	}
	i := num / 3
	return []int64{i - 1, i, i + 1}
}

func maximumEvenSplit(finalSum int64) []int64 {
	if finalSum%2 != 0 {
		return []int64{}
	}
	n := int64(1)
	for (n+1)*(n+2) <= finalSum {
		n++
	}
	res := make([]int64, n)
	for i := int64(1); i < n; i++ {
		v := 2 * i
		res[i-1] = v
		finalSum -= v
	}
	res[n-1] = finalSum
	return res
}

func goodTriplets(nums1 []int, nums2 []int) int64 {
	m2 := make(map[int]int, len(nums2))
	for i := 0; i < len(nums2); i++ {
		m2[nums2[i]] = i
	}

	lines := make([][2]int, 0, len(nums1))
	for i := 0; i < len(nums1); i++ {
		lines = append(lines, [2]int{i, m2[nums1[i]]})
	}
	sort.Slice(lines, func(i, j int) bool {
		return lines[i][0] < lines[j][0]
	})

	res := int64(0)
	for i := 0; i < len(lines)-2; i++ {
		for j := i + 1; j < len(lines)-1; j++ {
			if cross(lines[i], lines[j]) {
				continue
			}
			for k := j + 1; k < len(lines); k++ {
				if !cross(lines[j], lines[k]) {
					res++
				}
			}
		}
	}
	return res
}

func cross(a, b [2]int) bool {
	return a[0] < b[1] && a[1] > b[0]
}
