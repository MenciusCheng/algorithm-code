package w20220108

import (
	"fmt"
	"strings"
)

/*
5960. 将标题首字母大写
*/
func capitalizeTitle(title string) string {
	split := strings.Split(title, " ")
	arr := make([]string, 0, len(split))
	for _, s := range split {
		if len(s) <= 2 {
			str := strings.ToLower(s)
			arr = append(arr, str)
		} else {
			str := fmt.Sprintf("%s%s", strings.ToUpper(s[:1]), strings.ToLower(s)[1:])
			arr = append(arr, str)
		}
	}
	return strings.Join(arr, " ")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
5961. 链表最大孪生和
*/
func pairSum(head *ListNode) int {
	arr := make([]int, 0)
	p := head
	for p != nil {
		arr = append(arr, p.Val)
		p = p.Next
	}

	max := 0
	i := 0
	j := len(arr) - 1
	for i <= j {
		if arr[i]+arr[j] > max {
			max = arr[i] + arr[j]
		}
		i++
		j--
	}
	return max
}

/*
5962. 连接两字母单词得到的最长回文串
*/
func longestPalindrome(words []string) int {
	count := 0
	cnt := make(map[string]int)
	db := make(map[string]int)
	for _, word := range words {
		if word[0] == word[1] {
			db[word]++
		} else {
			rw := string([]byte{word[1], word[0]})
			if cnt[rw] > 0 {
				cnt[rw]--
				count += 2
			} else {
				cnt[word]++
			}
		}
	}
	maxSingle := 0
	for _, v := range db {
		if v%2 == 0 {
			count += v
		} else {
			count += v - 1
			if v > maxSingle {
				maxSingle = v
			}
		}
	}
	if maxSingle > 0 {
		count++
	}

	return count * 2
}

/*
5931. 用邮票贴满网格图
*/
func possibleToStamp(grid [][]int, stampHeight int, stampWidth int) bool {
	// 此代码不正确
	for i := range grid {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				if canFix(grid, i, j, i+stampHeight-1, j+stampWidth-1) {
					fix(grid, i, j, i+stampHeight-1, j+stampWidth-1)
					j += stampWidth - 1
				} else if i > 0 && j > 0 && grid[i-1][j] == 2 && grid[i][j-1] == 2 {
					fix(grid, i, j, i, j)
				} else if i > 0 && grid[i-1][j] == 2 && canFix(grid, i, j, i, j+stampWidth-1) {
					fix(grid, i, j, i, j+stampWidth-1)
					j += stampWidth - 1
				} else if j > 0 && grid[i][j-1] == 2 && canFix(grid, i, j, i+stampHeight-1, j) {
					fix(grid, i, j, i+stampHeight-1, j)
				} else {
					return false
				}
			}
		}
	}
	return true
}

func canFix(grid [][]int, i1, j1, i2, j2 int) bool {
	if i1 < 0 || i1 >= len(grid) || i2 < 0 || i2 >= len(grid) {
		return false
	}
	if j1 < 0 || j1 >= len(grid[0]) || j2 < 0 || j2 >= len(grid[0]) {
		return false
	}
	for i := i1; i <= i2; i++ {
		for j := j1; j <= j2; j++ {
			if grid[i][j] == 1 {
				return false
			}
		}
	}
	return true
}

func fix(grid [][]int, i1, j1, i2, j2 int) bool {
	for i := i1; i <= i2; i++ {
		for j := j1; j <= j2; j++ {
			grid[i][j] = 2
		}
	}
	return true
}
