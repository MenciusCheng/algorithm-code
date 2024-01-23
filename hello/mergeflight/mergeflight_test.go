package mergeflight

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestGroup_Do(t *testing.T) {
	wg := sync.WaitGroup{}

	fn := func(params []interface{}) (interface{}, error) {
		res := make([]interface{}, 0)
		for _, item := range params {
			res = append(res, fmt.Sprintf("%v (%T)", item, item))
		}
		fmt.Printf("fn, len:%d, params:%+v\n", len(params), params)
		time.Sleep(10 * time.Millisecond)
		return res, nil
	}

	g := NewGroup()
	n := 200
	capacity := 10
	delay := 10 * time.Millisecond
	for i := 0; i < n; i++ {
		go func(param int) {
			wg.Add(1)
			defer func() {
				wg.Done()
			}()

			want := fmt.Sprintf("%v (%T)", param, param)
			got, err, index := g.Do("key", fn, param, capacity, delay)
			if err != nil {
				t.Errorf("Do error = %v", err)
				return
			}
			arr, ok := got.([]interface{})
			if !ok {
				t.Errorf("got type error, got = %v", got)
				return
			}
			if index >= len(arr) {
				t.Errorf("index out of range, index = %v", index)
				return
			}
			if arr[index] != want {
				t.Errorf("v = %v, want = %v", arr[index], want)
				return
			}
		}(i)
	}
	wg.Wait()
}
