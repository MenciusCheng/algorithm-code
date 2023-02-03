package main

import "time"

func main() {
	//operateMap()
	//concurrencyRWMap()
	concurrencyRRMap()
}

func operateMap() {
	// 初始化
	hash1 := make(map[string]int)
	hash2 := make(map[string]int, 10)
	hash3 := map[string]int{
		"wei": 1,
	}

	// 访问
	v1 := hash3["wei"]
	v2, ok := hash3["wei"]

	// 赋值
	hash3["mei"] = 2

	// 删除键
	delete(hash3, "mei")

	println(hash1, hash2, hash3, v1, v2, ok)
}

// fatal error: concurrent map read and map write
func concurrencyRWMap() {
	cnt := make(map[int]int)
	go func() {
		for {
			cnt[0] = 5
		}
	}()
	go func() {
		for {
			_ = cnt[1]
		}
	}()
	time.Sleep(1 * time.Second)
}

func concurrencyRRMap() {
	cnt := make(map[int]int)
	go func() {
		for {
			_ = cnt[0]
		}
	}()
	go func() {
		for {
			_ = cnt[1]
		}
	}()
	time.Sleep(1 * time.Second)
}
