package main

import "fmt"

func main() {
	var cnt map[int]int
	fmt.Println(cnt[1]) // nil map 可以读
	//cnt[1] = 2 // nil map 写值会报错

	cnt = map[int]int{
		1: 111,
		2: 222,
		3: 333,
	}
	for k, v := range cnt {
		//if k == 1 {
		//	delete(cnt, 2)
		//	delete(cnt, 3)
		//}
		fmt.Printf("读了 k:%d, v:%d\n", k, v)
		if k == 2 {
			delete(cnt, 1)
			delete(cnt, 3)
			fmt.Printf("删除了 k:1, 3\n")
		}
		if k == 3 {
			delete(cnt, 1)
			delete(cnt, 2)
			fmt.Printf("删除了 k:1, 2\n")
		}
	}
	fmt.Printf("最后 cnt:%+v\n", cnt)

	//readNilChannel()
	//sendNilChannel()
}

func readNilChannel() {
	var ch chan int
	<-ch // fatal error: all goroutines are asleep - deadlock!
}

func sendNilChannel() {
	var ch chan int
	ch <- 6 // fatal error: all goroutines are asleep - deadlock!
}
