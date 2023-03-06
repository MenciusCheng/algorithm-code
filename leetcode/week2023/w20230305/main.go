package w20230305

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	arr := make([]int, 0)
	ns := []*TreeNode{root}
	for len(ns) > 0 {
		next := make([]*TreeNode, 0)
		v := 0
		for _, item := range ns {
			v += item.Val
			if item.Left != nil {
				next = append(next, item.Left)
			}
			if item.Right != nil {
				next = append(next, item.Right)
			}
		}
		arr = append(arr, v)
		ns = next
	}
	sort.Ints(arr)
	if k > len(arr) {
		return -1
	}
	return int64(arr[len(arr)-k])
}

func findValidSplit(nums []int) int {
	r := 0
	for i := 0; i < len(nums); i++ {
		for j := len(nums) - 1; j >= r+1 && j >= i+1; j-- {
			if gcd(nums[i], nums[j]) > 1 {
				r = j
				break
			}
		}
		if r >= len(nums)-1 {
			return -1
		}
		if i >= r {
			return r
		}
	}
	return -1
}

func gcd(x, y int) int {
	var tmp int
	for {
		tmp = x % y
		if tmp > 0 {
			x = y
			y = tmp
		} else {
			return y
		}
	}
}
