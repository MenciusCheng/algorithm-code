package w20230218

import (
	"math"
	"strconv"
)

func minMaxDifference(num int) int {
	s := strconv.Itoa(num)
	bsmax := make([]byte, len(s))
	for i := range s {
		if s[i] != '9' {
			for j := i; j < len(s); j++ {
				if s[j] == s[i] {
					bsmax[j] = '9'
				} else {
					bsmax[j] = s[j]
				}
			}
			break
		} else {
			bsmax[i] = s[i]
		}
	}

	ch := s[0]
	bsmin := make([]byte, len(s))
	for i := range s {
		if s[i] == ch {
			bsmin[i] = '0'
		} else {
			bsmin[i] = s[i]
		}
	}

	vmax, _ := strconv.Atoi(string(bsmax))
	vmin, _ := strconv.Atoi(string(bsmin))
	return vmax - vmin
}

func minImpossibleOR(nums []int) int {
	cnt := make(map[int]bool)
	for _, num := range nums {
		cnt[num] = true
	}
	for i := 1; i < math.MaxInt; i *= 2 {
		if !cnt[i] {
			return i
		}
	}
	return 0
}

func handleQuery(nums1 []int, nums2 []int, queries [][]int) []int64 {
	c := 0
	for _, num := range nums1 {
		if num == 1 {
			c++
		}
	}
	sum := 0
	for _, num := range nums2 {
		sum += num
	}
	res := make([]int64, 0)
	for _, q := range queries {
		if q[0] == 1 {
			for i := q[1]; i <= q[2]; i++ {
				if nums1[i] == 1 {
					nums1[i] = 0
					c--
				} else {
					nums1[i] = 1
					c++
				}
			}
		} else if q[0] == 2 {
			sum += c * q[1]
		} else {
			res = append(res, int64(sum))
		}
	}
	return res
}
