package w20211225

/*
5946. 句子中的最多单词数
*/
func mostWordsFound(sentences []string) int {
	var max int
	for _, sentence := range sentences {
		wCount := 0
		for i := 0; i < len(sentence); i++ {
			if sentence[i] == ' ' {
				wCount++
			}
		}
		if wCount > max {
			max = wCount
		}
	}
	return max
}

/*
5947. 从给定原材料中找到所有可以做出的菜
*/
func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	recipesMap := make(map[string]int, len(recipes))
	for i, item := range recipes {
		recipesMap[item] = i
	}

	suppliesMap := make(map[string]bool, len(supplies))
	for _, supply := range supplies {
		suppliesMap[supply] = true
	}

	type R struct {
		status int
		needs  []int
	}

	links := make([]R, 0, len(recipes))

	resMap := make(map[string]bool)
	for i := 0; i < len(ingredients); i++ {
		r := R{}
		for _, item := range ingredients[i] {
			if suppliesMap[item] == false && resMap[item] == false {
				if index, ok := recipesMap[item]; ok {
					r.needs = append(r.needs, index)
					r.status = 1
				} else {
					r.status = 2
					break
				}
			}
		}
		if r.status == 0 {
			resMap[recipes[i]] = true
		}
		links = append(links, r)
	}

	//for _, link := range links {
	//	//if link
	//}

	res := make([]string, 0, len(resMap))
	for key := range resMap {
		res = append(res, key)
	}
	return res
}

/*
5948. 判断一个括号字符串是否有效
*/
