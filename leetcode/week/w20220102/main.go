package main

import "sort"

/*
5967. 检查是否所有 A 都在 B 之前
*/
func checkString(s string) bool {
	last := ' '
	for _, c := range s {
		if last != c {
			if last == ' ' || last == 'a' && c == 'b' {
				last = c
			} else {
				return false
			}
		}
	}
	return true
}

/*
5968. 银行中的激光束数量
*/
func numberOfBeams(bank []string) int {
	sum := 0
	last := 0
	for _, line := range bank {
		count := 0
		for _, c := range line {
			if c == '1' {
				count++
			}
		}
		if count > 0 {
			sum += last * count
			last = count
		}
	}

	return sum
}

/*
5969. 摧毁小行星
*/
func asteroidsDestroyed(mass int, asteroids []int) bool {
	sort.Ints(asteroids)
	for _, asteroid := range asteroids {
		if mass >= asteroid {
			mass += asteroid
		} else {
			return false
		}
	}
	return true
}

/*
5970. 参加会议的最多员工数
*/
func maximumInvitations(favorite []int) int {
	notStart := make(map[int]bool)
	for i := 0; i < len(favorite); i++ {
		if i != favorite[favorite[i]] {
			notStart[favorite[i]] = true
		}
	}

	loop := make(map[[2]int]int)
	beFavorite := make([][]int, len(favorite))
	for i, i2 := range favorite {
		beFavorite[i2] = append(beFavorite[i2], i)
		if i == favorite[i2] {
			loop[[2]int{i, i2}] = 1
		}
	}

	var max int

	for i := 0; i < len(favorite); i++ {
		if notStart[i] {
			continue
		}
		p := i
		visited := make(map[int]bool)
		visited[p] = true
		count := 1
		for !visited[favorite[p]] {
			p = favorite[p]
			visited[p] = true
			count++
		}
		if p == favorite[favorite[p]] {
			lc := count - 1
			if lc > loop[[2]int{favorite[p], p}] {
				loop[[2]int{favorite[p], p}] = lc // 替换已存在的链长
			}
		} else if favorite[p] == i {
			if count > max {
				max = count
			}
		}
	}
	loopCount := 0
	for _, i := range loop {
		loopCount += i
	}
	if loopCount > max {
		max = loopCount
	}

	return max
}
