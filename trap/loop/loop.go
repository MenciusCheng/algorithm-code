package loop

import (
	"fmt"
	"time"
)

// 在循环中创建协程
func ForGoroutine() {
	arr := []string{"a", "b", "c"}

	// 打印错误的值
	fmt.Println("wrong for goroutine:")
	for i, s := range arr {
		go func() {
			PrintAB(i, s)
		}()
	}
	time.Sleep(1 * time.Second)

	// 修复写法1
	fmt.Println("fix 1:")
	for i, s := range arr {
		i := i
		s := s
		go func() {
			PrintAB(i, s)
		}()
	}
	time.Sleep(1 * time.Second)

	// 修复写法2
	fmt.Println("fix 2:")
	for i, s := range arr {
		go func(i int, s string) {
			PrintAB(i, s)
		}(i, s)
	}
	time.Sleep(1 * time.Second)
}

func PrintAB(i int, s string) {
	fmt.Println(i, s)
}
