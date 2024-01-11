package csingleflight

import (
	"fmt"
	"golang.org/x/sync/singleflight"
)

// 官方包示例
func GroupExample() {
	g := new(singleflight.Group)

	block := make(chan struct{})
	res1c := g.DoChan("key", func() (interface{}, error) {
		<-block
		return "func 1", nil
	})
	res2c := g.DoChan("key", func() (interface{}, error) {
		<-block
		return "func 2", nil
	})
	close(block)

	res1 := <-res1c
	res2 := <-res2c

	// Results are shared by functions executed with duplicate keys.
	fmt.Println("Shared:", res2.Shared)
	// Only the first function is executed: it is registered and started with "key",
	// and doesn't complete before the second funtion is registered with a duplicate key.
	fmt.Println("Equal results:", res1.Val.(string) == res2.Val.(string))
	fmt.Println("Result:", res1.Val)
}
