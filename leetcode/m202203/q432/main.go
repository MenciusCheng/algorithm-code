package main

import (
	"container/heap"
	"fmt"
)

/*
https://leetcode-cn.com/problems/all-oone-data-structure/

432. 全 O(1) 的数据结构
请你设计一个用于存储字符串计数的数据结构，并能够返回计数最小和最大的字符串。

实现 AllOne 类：

AllOne() 初始化数据结构的对象。
inc(String key) 字符串 key 的计数增加 1 。如果数据结构中尚不存在 key ，那么插入计数为 1 的 key 。
dec(String key) 字符串 key 的计数减少 1 。如果 key 的计数在减少后为 0 ，那么需要将这个 key 从数据结构中删除。测试用例保证：在减少计数前，key 存在于数据结构中。
getMaxKey() 返回任意一个计数最大的字符串。如果没有元素存在，返回一个空字符串 "" 。
getMinKey() 返回任意一个计数最小的字符串。如果没有元素存在，返回一个空字符串 "" 。

示例：

输入
["AllOne", "inc", "inc", "getMaxKey", "getMinKey", "inc", "getMaxKey", "getMinKey"]
[[], ["hello"], ["hello"], [], [], ["leet"], [], []]
输出
[null, null, null, "hello", "hello", null, "hello", "leet"]

解释
AllOne allOne = new AllOne();
allOne.inc("hello");
allOne.inc("hello");
allOne.getMaxKey(); // 返回 "hello"
allOne.getMinKey(); // 返回 "hello"
allOne.inc("leet");
allOne.getMaxKey(); // 返回 "hello"
allOne.getMinKey(); // 返回 "leet"

提示：

1 <= key.length <= 10
key 由小写英文字母组成
测试用例保证：在每次调用 dec 时，数据结构中总存在 key
最多调用 inc、dec、getMaxKey 和 getMinKey 方法 5 * 10^4 次
*/
func main() {
	allOne := Constructor()
	allOne.Inc("hello")
	allOne.Inc("hello")
	fmt.Println(allOne.GetMaxKey() == "hello") // 返回 "hello"
	fmt.Println(allOne.GetMinKey() == "hello") // 返回 "hello"
	allOne.Inc("leet")
	fmt.Println(allOne.GetMaxKey() == "hello") // 返回 "hello"
	fmt.Println(allOne.GetMinKey() == "leet")  // 返回 "leet"
}

type AllOne struct {
	cnt     map[string]int
	minHeap *Heap
	maxHeap *MaxHeap
}

func Constructor() AllOne {
	cnt := make(map[string]int)
	a := AllOne{
		cnt: cnt,
		minHeap: &Heap{
			arr: []string{},
			cnt: cnt,
		},
		maxHeap: &MaxHeap{
			Heap: Heap{
				arr: []string{},
				cnt: cnt,
			},
		},
	}
	return a
}

func (this *AllOne) Inc(key string) {
	if count, ok := this.cnt[key]; ok {
		this.cnt[key] = count + 1
		index := 0
		for i, s := range this.minHeap.arr {
			if s == key {
				index = i
				break
			}
		}
		heap.Fix(this.minHeap, index)
		for i, s := range this.maxHeap.arr {
			if s == key {
				index = i
				break
			}
		}
		heap.Fix(this.maxHeap, index)
	} else {
		this.cnt[key] = 1
		heap.Push(this.minHeap, key)
		heap.Push(this.maxHeap, key)
	}
}

func (this *AllOne) Dec(key string) {
	if count, ok := this.cnt[key]; ok {
		if count > 1 {
			this.cnt[key] = count - 1

			index := 0
			for i, s := range this.minHeap.arr {
				if s == key {
					index = i
					break
				}
			}
			heap.Fix(this.minHeap, index)
			for i, s := range this.maxHeap.arr {
				if s == key {
					index = i
					break
				}
			}
			heap.Fix(this.maxHeap, index)
		} else {
			delete(this.cnt, key)

			index := 0
			for i, s := range this.minHeap.arr {
				if s == key {
					index = i
					break
				}
			}
			heap.Remove(this.minHeap, index)
			for i, s := range this.maxHeap.arr {
				if s == key {
					index = i
					break
				}
			}
			heap.Remove(this.maxHeap, index)
		}
	}
}

func (this *AllOne) GetMaxKey() string {
	if len(this.cnt) == 0 {
		return ""
	}
	return this.maxHeap.arr[0]
}

func (this *AllOne) GetMinKey() string {
	if len(this.cnt) == 0 {
		return ""
	}
	return this.minHeap.arr[0]
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */

// Heap 优先队列
type Heap struct {
	arr []string
	cnt map[string]int
}

func (x Heap) Len() int           { return len(x.arr) }
func (x Heap) Less(i, j int) bool { return x.cnt[x.arr[i]] < x.cnt[x.arr[j]] }
func (x Heap) Swap(i, j int)      { x.arr[i], x.arr[j] = x.arr[j], x.arr[i] }

func (x *Heap) Push(v interface{}) {
	x.arr = append(x.arr, v.(string))
}

func (x *Heap) Pop() interface{} {
	a := x.arr
	v := a[len(a)-1]
	x.arr = a[:len(a)-1]
	return v
}

type MaxHeap struct {
	Heap
}

func (x MaxHeap) Less(i, j int) bool {
	return x.cnt[x.arr[i]] > x.cnt[x.arr[j]]
}
