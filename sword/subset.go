package sword

import "fmt"

func allSubset(nums []int) {
	for i := 0; i < 1<<len(nums); i++ {
		arr := make([]int, 0)
		for j, num := range nums {
			if i>>j&1 == 1 {
				arr = append(arr, num)
			}
		}
		fmt.Printf("%+v\n", arr)
	}
}
