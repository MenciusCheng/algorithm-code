package main

import "strings"

func replaceWords(dictionary []string, sentence string) string {
	dm := make(map[string]bool)
	for _, s := range dictionary {
		dm[s] = true
	}

	split := strings.Split(sentence, " ")
	arr := make([]string, 0, len(split))
	for _, word := range split {
		s := ""
		for i := 1; i <= len(word); i++ {
			if dm[word[:i]] {
				s = word[:i]
				break
			}
		}
		if len(s) > 0 {
			arr = append(arr, s)
		} else {
			arr = append(arr, word)
		}
	}

	return strings.Join(arr, " ")
}
