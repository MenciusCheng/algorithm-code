package w20220313

import "fmt"

func findKDistantIndices(nums []int, key int, k int) []int {
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if i == 2 {
			fmt.Printf("aa")
		}
		for j := 0; j < len(nums); j++ {
			if abs(i, j) <= k && nums[j] == key {
				res = append(res, i)
				break
			}
		}
	}
	return res
}

func abs(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func digArtifacts(n int, artifacts [][]int, dig [][]int) int {
	dMap := make(map[[2]int]bool)
	for i := 0; i < len(dig); i++ {
		dMap[[2]int{dig[i][0], dig[i][1]}] = true
	}

	res := 0
	for i := 0; i < len(artifacts); i++ {
		art := artifacts[i]
		isAllD := true
		for ri := art[0]; ri <= art[2]; ri++ {
			for ci := art[1]; ci <= art[3]; ci++ {
				if !dMap[[2]int{ri, ci}] {
					isAllD = false
				}
			}
		}
		if isAllD {
			res++
		}
	}
	return res
}

func maximumTop(nums []int, k int) int {
	if len(nums) == 1 && k%2 == 1 {
		return -1
	}

	max := 0
	for i := 0; i < k-1 && i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	if len(nums) > k && nums[k] > max {
		max = nums[k]
	}
	return max
}
