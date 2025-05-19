package main

import "fmt"

func main() {
	testErr1()
	//testErr2()
}

// 在函数中覆盖err
func testErr1() {
	pid1, err := getError()
	if true {
		pid2, err := getErrorNil()
		if err != nil {
			fmt.Printf("getErrorNil err=%s\n", err)
			return
		}
		_ = pid2
	}
	if err != nil {
		fmt.Printf("getError err=%s\n", err)
		return
	}
	_ = pid1
	fmt.Printf("success")
}

// 同一作用域覆盖err
func testErr2() {
	pid1, err := getError()
	_ = pid1

	pid2, err := getErrorNil()
	if err != nil {
		fmt.Printf("getErrorNil err=%s\n", err)
		return
	}
	_ = pid2
	if err != nil {
		fmt.Printf("getError err=%s\n", err)
		return
	}
	fmt.Printf("success")
}

func getError() (int, error) {
	return 1, fmt.Errorf("getError")
}

func getErrorNil() (int, error) {
	return 2, nil
}
