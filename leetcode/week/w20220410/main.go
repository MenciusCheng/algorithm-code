package w20220410

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func largestInteger(num int) int {
	str := strconv.Itoa(num)
	a0 := make([]int, 0)
	a1 := make([]int, 0)
	for i := 0; i < len(str); i++ {
		v := int(str[i] - '0')
		if v%2 == 0 {
			a0 = append(a0, v)
		} else {
			a1 = append(a1, v)
		}
	}
	sort.Ints(a0)
	sort.Ints(a1)
	var i0, i1 int
	rs := make([]byte, len(str))
	for i := len(str) - 1; i >= 0; i-- {
		v := int(str[i] - '0')
		if v%2 == 0 {
			rs[i] = byte(a0[i0]) + '0'
			i0++
		} else {
			rs[i] = byte(a1[i1]) + '0'
			i1++
		}
	}
	res, _ := strconv.Atoi(string(rs))
	return res
}

func minimizeResult(expression string) string {
	pi := 0
	for i := 0; i < len(expression); i++ {
		if expression[i] == '+' {
			pi = i
			break
		}
	}

	res := ""
	min := math.MaxInt32
	for i := 0; i < pi; i++ {
		for j := pi + 2; j <= len(expression); j++ {
			s0 := expression[:i]
			s1 := expression[i:pi]
			s2 := expression[pi+1 : j]
			s3 := expression[j:]

			a0 := 1
			if len(s0) > 0 {
				a0, _ = strconv.Atoi(s0)
			}
			a1 := 0
			if len(s1) > 0 {
				a1, _ = strconv.Atoi(s1)
			}
			a2 := 0
			if len(s2) > 0 {
				a2, _ = strconv.Atoi(s2)
			}
			a3 := 1
			if len(s3) > 0 {
				a3, _ = strconv.Atoi(s3)
			}

			sum := 0
			m := a1 + a2
			if m > 0 {
				sum = a0 * m * a3
			} else {
				sum = a0 * a3
			}
			if sum < min {
				min = sum
				res = fmt.Sprintf("%s(%s+%s)%s", s0, s1, s2, s3)
			}
		}
	}
	return res
}

func maximumProduct(nums []int, k int) int {
	if len(nums) == 1 {
		return nums[0] + k
	}
	sort.Ints(nums)
	for k > 0 {
		down(nums)
		diff := nums[1] - nums[0]
		if diff == 0 {
			nums[0]++
			k--
		} else {
			if diff >= k {
				nums[0] += k
				k = 0
			} else {
				nums[0] += diff
				k -= diff
			}
		}
	}
	sum := 1
	for i := 0; i < len(nums); i++ {
		sum *= nums[i]
		sum %= 1000000007
	}
	return sum
}

func down(nums []int) {
	i := 0
	for i < len(nums)-1 && nums[i] > nums[i+1] {
		nums[i], nums[i+1] = nums[i+1], nums[i]
		i++
	}
}
