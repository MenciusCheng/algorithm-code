package main

import (
	"fmt"
	"sort"
	"strconv"
)

func minimumSum(num int) int {
	arr := make([]int, 0)
	for num > 0 {
		arr = append(arr, num%10)
		num /= 10
	}
	sort.Ints(arr)

	sum := 0
	a := 1
	for len(arr) > 0 {
		sum += arr[len(arr)-1] * a
		arr = arr[:len(arr)-1]
		if len(arr) > 0 {
			sum += arr[len(arr)-1] * a
			arr = arr[:len(arr)-1]
		}
		a *= 10
	}
	return sum
}

func pivotArray(nums []int, pivot int) []int {
	arr := make([]int, 0, len(nums))
	for _, num := range nums {
		if num < pivot {
			arr = append(arr, num)
		}
	}
	for _, num := range nums {
		if num == pivot {
			arr = append(arr, num)
		}
	}
	for _, num := range nums {
		if num > pivot {
			arr = append(arr, num)
		}
	}
	return arr
}

func minCostSetTime(startAt int, moveCost int, pushCost int, targetSeconds int) int {
	m1 := targetSeconds / 60
	if m1 == 100 {
		m1--
	}
	s1 := targetSeconds - m1*60
	ts1 := fmt.Sprintf("%02d%02d", m1, s1)
	t1, _ := strconv.Atoi(ts1)
	ts1 = strconv.Itoa(t1)

	p := startAt
	cost := 0
	for i := 0; i < len(ts1); i++ {
		if int(ts1[i]-'0') == p {
			cost += pushCost
		} else {
			cost += pushCost + moveCost
			p = int(ts1[i] - '0')
		}
	}

	if m1 > 0 && s1 <= 39 {
		m2 := m1 - 1
		s2 := targetSeconds - m2*60
		ts2 := fmt.Sprintf("%02d%02d", m2, s2)
		t2, _ := strconv.Atoi(ts2)
		ts2 = strconv.Itoa(t2)

		p = startAt
		cost2 := 0
		for i := 0; i < len(ts2); i++ {
			if int(ts2[i]-'0') == p {
				cost2 += pushCost
			} else {
				cost2 += pushCost + moveCost
				p = int(ts2[i] - '0')
			}
		}

		if cost2 < cost {
			cost = cost2
		}
	}
	return cost
}

func minimumDifference(nums []int) int64 {
	n := len(nums) / 3
	left, mid, right := make([]int, 0, n), make([]int, 0, n), make([]int, 0, n)

	for i := 0; i < len(nums); i++ {
		if i < n {
			left = append(left, nums[i])
		} else if i >= 2*n {
			right = append(right, nums[i])
		} else {
			mid = append(mid, nums[i])
		}
	}
	sort.Ints(left)
	sort.Ints(mid)
	sort.Ints(right)

	l1, r1 := n-1, 0
	p1, p2 := 0, n-1
	for p1 <= p2 {
		if l1 >= 0 && r1 <= n-1 && left[l1]-mid[p1] > 0 && mid[p2]-right[r1] > 0 {
			if left[l1]-mid[p1] > mid[p2]-right[r1] {
				left[l1] = mid[p1]
				l1--
				p1++
			} else {
				right[r1] = mid[p2]
				r1++
				p2--
			}
		} else if l1 >= 0 && left[l1]-mid[p1] > 0 {
			left[l1] = mid[p1]
			l1--
			p1++
		} else if r1 <= n-1 && mid[p2]-right[r1] > 0 {
			right[r1] = mid[p2]
			r1++
			p2--
		} else {
			break
		}
	}

	leftSum := int64(0)
	for _, item := range left {
		leftSum += int64(item)
	}
	rightSum := int64(0)
	for _, item := range right {
		rightSum += int64(item)
	}
	return leftSum - rightSum
}
