package w20220305

import "sort"

func mostFrequent(nums []int, key int) int {
	cnt := make(map[int]int)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == key {
			cnt[nums[i+1]]++
		}
	}

	res := 0
	maxCount := 0
	for k, v := range cnt {
		if v > maxCount {
			res = k
			maxCount = v
		}
	}
	return res
}

func cellsInRange(s string) []string {
	a := s[0]
	b := s[1]

	res := make([]string, 0)
	for a <= s[3] && b <= s[4] {
		res = append(res, string([]byte{a, b}))
		b++
		if b > s[4] {
			b = s[1]
			a++
		}
	}
	return res
}

func minimalKSum(nums []int, k int) int64 {
	sort.Ints(nums)
	n := k
	cnt := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if nums[i] <= n {
			if cnt[nums[i]] == 0 {
				n++
				cnt[nums[i]]++
			}
		} else {
			break
		}
	}
	sum := int64(1+n) * int64(n) / 2
	for i := range cnt {
		sum -= int64(i)
	}
	return sum
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	nodeMap := make(map[int]*TreeNode)
	havePMap := make(map[int]bool)

	for i := 0; i < len(descriptions); i++ {
		desc := descriptions[i]

		p := nodeMap[desc[0]]
		if p == nil {
			p = &TreeNode{
				Val: desc[0],
			}
			nodeMap[desc[0]] = p
		}

		c := nodeMap[desc[1]]
		if c == nil {
			c = &TreeNode{
				Val: desc[1],
			}
			nodeMap[desc[1]] = c
		}
		havePMap[desc[1]] = true

		if desc[2] == 1 {
			p.Left = c
		} else {
			p.Right = c
		}
	}

	for val, node := range nodeMap {
		if !havePMap[val] {
			return node
		}
	}
	return nil
}
