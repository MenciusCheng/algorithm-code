package main

import "fmt"

// https://leetcode-cn.com/problems/construct-the-rectangle/
// 492. 构造矩形
func main() {
	fmt.Println(constructRectangle(4))
}

func constructRectangle(area int) []int {
	aw := 1
	al := area
	for w := aw + 1; w < al; w++ {
		l := area / w
		if l < w {
			break
		}
		if area%w != 0 {
			continue
		}
		if l-w < al-aw {
			aw, al = w, l
		}
	}
	return []int{al, aw}
}
