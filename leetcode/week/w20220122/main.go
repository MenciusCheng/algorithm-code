package w20220122

import (
	"sort"
)

func minimumCost(cost []int) int {
	sort.Ints(cost)

	sum := 0
	m := len(cost) % 3

	for i := 0; i < m; i++ {
		sum += cost[i]
	}
	for i := m; i < len(cost); i += 3 {
		sum += cost[i+1] + cost[i+2]
	}

	return sum
}

func numberOfArrays(differences []int, lower int, upper int) int {
	var cur, min, max int
	for i := 0; i < len(differences); i++ {
		cur += differences[i]
		if cur < min {
			min = cur
		}
		if cur > max {
			max = cur
		}
	}

	m := max + (lower - min)
	if m > upper {
		return 0
	} else {
		return upper - m + 1
	}
}

func highestRankedKItems(grid [][]int, pricing []int, start []int, k int) [][]int {
	type Product struct {
		step  int
		price int
		row   int
		col   int
	}
	products := make([]Product, 0)

	next := [][2]int{
		{start[0], start[1]},
	}

	step := 0
	if grid[start[0]][start[1]] > 1 && grid[start[0]][start[1]] >= pricing[0] && grid[start[0]][start[1]] <= pricing[1] {
		products = append(products, Product{
			step:  step,
			price: grid[start[0]][start[1]],
			row:   start[0],
			col:   start[1],
		})
	}
	grid[start[0]][start[1]] = 0
	//fmt.Printf("r: %d, w: %d\n", len(grid), len(grid[0]))
	for len(next) > 0 {
		//fmt.Printf("step: %d, next.len: %d, prd.len: %d\n", step, len(next), len(products))
		if len(products) >= k {
			break
		}
		step++

		newNext := make([][2]int, 0)
		for i := 0; i < len(next); i++ {
			row := next[i][0]
			col := next[i][1]

			upRow := row - 1
			if upRow >= 0 && grid[upRow][col] != 0 {
				newNext = append(newNext, [2]int{upRow, col})

				if grid[upRow][col] > 1 && grid[upRow][col] >= pricing[0] && grid[upRow][col] <= pricing[1] {
					products = append(products, Product{
						step:  step,
						price: grid[upRow][col],
						row:   upRow,
						col:   col,
					})
				}
				grid[upRow][col] = 0
			}
			downRow := row + 1
			if downRow < len(grid) && grid[downRow][col] != 0 {
				newNext = append(newNext, [2]int{downRow, col})

				if grid[downRow][col] > 1 && grid[downRow][col] >= pricing[0] && grid[downRow][col] <= pricing[1] {
					products = append(products, Product{
						step:  step,
						price: grid[downRow][col],
						row:   downRow,
						col:   col,
					})
				}
				grid[downRow][col] = 0
			}

			leftCol := col - 1
			if leftCol >= 0 && grid[row][leftCol] != 0 {
				newNext = append(newNext, [2]int{row, leftCol})

				if grid[row][leftCol] > 1 && grid[row][leftCol] >= pricing[0] && grid[row][leftCol] <= pricing[1] {
					products = append(products, Product{
						step:  step,
						price: grid[row][leftCol],
						row:   row,
						col:   leftCol,
					})
				}
				grid[row][leftCol] = 0
			}
			rightCol := col + 1
			if rightCol < len(grid[0]) && grid[row][rightCol] != 0 {
				newNext = append(newNext, [2]int{row, rightCol})

				if grid[row][rightCol] > 1 && grid[row][rightCol] >= pricing[0] && grid[row][rightCol] <= pricing[1] {
					products = append(products, Product{
						step:  step,
						price: grid[row][rightCol],
						row:   row,
						col:   rightCol,
					})
				}
				grid[row][rightCol] = 0
			}
		}
		next = newNext
	}

	sort.Slice(products, func(i, j int) bool {
		if products[i].step < products[j].step {
			return true
		} else if products[i].step == products[j].step {
			if products[i].price < products[j].price {
				return true
			} else if products[i].price == products[j].price {
				if products[i].row < products[j].row {
					return true
				} else if products[i].row == products[j].row {
					if products[i].col < products[j].col {
						return true
					}
				}
			}
		}
		return false
	})
	res := make([][]int, 0, k)
	for i := 0; i < len(products) && i < k; i++ {
		res = append(res, []int{products[i].row, products[i].col})
	}
	return res
}
