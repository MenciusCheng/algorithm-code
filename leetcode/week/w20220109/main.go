package w20220109

/*
5976. 检查是否每一行每一列都包含全部整数
*/
func checkValid(matrix [][]int) bool {
	for i := 0; i < len(matrix); i++ {
		cnt := make(map[int]int)
		for j := 0; j < len(matrix[i]); j++ {
			if cnt[matrix[i][j]] == 1 {
				return false
			}
			cnt[matrix[i][j]] = 1
		}
		cnt = make(map[int]int)
		for j := 0; j < len(matrix[i]); j++ {
			if cnt[matrix[j][i]] == 1 {
				return false
			}
			cnt[matrix[j][i]] = 1
		}
	}
	return true
}

/*
5978. 统计追加字母可以获得的单词数
*/
func wordCount(startWords []string, targetWords []string) int {
	// 此代码超时
	cnt := make(map[int][][26]int)
	for _, word := range startWords {
		cc := [26]int{}
		for _, c := range word {
			cc[c-'a'] = 1
		}
		cnt[len(word)] = append(cnt[len(word)], cc)
	}
	res := 0
	for _, word := range targetWords {
		ccs := cnt[len(word)-1]
		for _, cc := range ccs {
			diff := 0
			for _, c := range word {
				if cc[c-'a'] == 0 {
					diff++
				}
			}
			if diff == 1 {
				res++
				break
			}
		}
	}
	return res
}
