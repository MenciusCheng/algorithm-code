package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

func main() {
	testSortIdsIfChangeOrigin()
}

func subSlice() {
	a := []int{0, 1, 2, 3, 4, 5}

	fmt.Println(a[0:1])
	fmt.Println(a[6:6])
	fmt.Println(a[5:6])
}

func sliceItemPointer() {
	arr := []int{1, 2, 3, 4, 5}
	res := make([]*int, 0)

	for _, item := range arr {
		fmt.Println("item=", item, " p=", &item)
		res = append(res, &item)
	}
	fmt.Println("res=", res)
}

func splitArr() {
	arr := strings.Split("", ",")
	fmt.Printf("arr: %+v, len: %d\n", arr, len(arr))
	arr = strings.Split("afwef", ",")
	fmt.Printf("arr: %+v, len: %d\n", arr, len(arr))
}

func findRepeat(arr []string) {
	m := make(map[string]bool)
	for i, str := range arr {
		if !m[str] {
			m[str] = true
		} else {
			fmt.Printf("%d) %s\n", i, str)
		}
	}
}

func splitArrN() {
	arr := strings.Split("a,b,c,d", ",")
	bs, _ := json.Marshal(arr)
	fmt.Printf("arr: %s\n", string(bs))
	arr = strings.SplitN("a,b,c,d", ",", 3)
	bs, _ = json.Marshal(arr)
	fmt.Printf("arr: %s\n", string(bs))
	arr = strings.SplitN("a,b,c,d", ",", 5)
	bs, _ = json.Marshal(arr)
	fmt.Printf("arr: %s\n", string(bs))
}

// 测试对数组排序是否会影响原数组
func testSortIdsIfChangeOrigin() {
	playerIds := []int64{5, 2, 4, 3, 1}
	fmt.Printf("input playerIds: %+v\n", playerIds)
	sortedIds := sortPlayerIds(playerIds)
	fmt.Println("sorted:")
	fmt.Printf("playerIds: %+v\n", playerIds)
	fmt.Printf("sortedIds: %+v", sortedIds)
}

func sortPlayerIds(playerIds []int64) []int64 {
	sort.Slice(playerIds, func(i, j int) bool {
		return playerIds[i] < playerIds[j]
	})
	return playerIds
}
