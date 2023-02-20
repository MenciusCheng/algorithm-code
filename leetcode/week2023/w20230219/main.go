package w20230219

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	var i1, i2 int
	res := make([][]int, 0)
	for i1 < len(nums1) || i2 < len(nums2) {
		if i1 >= len(nums1) {
			res = append(res, nums2[i2])
			i2++
		} else if i2 >= len(nums2) {
			res = append(res, nums1[i1])
			i1++
		} else if nums1[i1][0] < nums2[i2][0] {
			res = append(res, nums1[i1])
			i1++
		} else if nums1[i1][0] > nums2[i2][0] {
			res = append(res, nums2[i2])
			i2++
		} else {
			res = append(res, []int{nums1[i1][0], nums1[i1][1] + nums2[i2][1]})
			i1++
			i2++
		}
	}
	return res
}

func minOperations(n int) int {
	ps := make([]int, 20)
	ps[0] = 1
	for i := 1; i < len(ps); i++ {
		ps[i] = ps[i-1] * 2
	}

	var a, b int
	isNeg := false
	res := 0
	for n != 0 {
		res++
		isNeg = n < 0
		if isNeg {
			n = -n
		}
		for i, v := range ps {
			if v == n {
				return res
			} else if v > n {
				a = ps[i-1]
				b = v
				break
			}
		}
		if n-a <= b-n {
			n -= a
		} else {
			n -= b
		}
		if isNeg {
			n = -n
		}
	}
	return res
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
