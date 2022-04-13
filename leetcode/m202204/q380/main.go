package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
https://leetcode-cn.com/problems/insert-delete-getrandom-o1/

380. O(1) 时间插入、删除和获取随机元素
实现RandomizedSet 类：

RandomizedSet() 初始化 RandomizedSet 对象
bool insert(int val) 当元素 val 不存在时，向集合中插入该项，并返回 true ；否则，返回 false 。
bool remove(int val) 当元素 val 存在时，从集合中移除该项，并返回 true ；否则，返回 false 。
int getRandom() 随机返回现有集合中的一项（测试用例保证调用此方法时集合中至少存在一个元素）。每个元素应该有 相同的概率 被返回。
你必须实现类的所有函数，并满足每个函数的 平均 时间复杂度为 O(1) 。

示例：

输入
["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
[[], [1], [2], [2], [], [1], [2], []]
输出
[null, true, false, true, 2, true, false, 2]

解释
RandomizedSet randomizedSet = new RandomizedSet();
randomizedSet.insert(1); // 向集合中插入 1 。返回 true 表示 1 被成功地插入。
randomizedSet.remove(2); // 返回 false ，表示集合中不存在 2 。
randomizedSet.insert(2); // 向集合中插入 2 。返回 true 。集合现在包含 [1,2] 。
randomizedSet.getRandom(); // getRandom 应随机返回 1 或 2 。
randomizedSet.remove(1); // 从集合中移除 1 ，返回 true 。集合现在包含 [2] 。
randomizedSet.insert(2); // 2 已在集合中，所以返回 false 。
randomizedSet.getRandom(); // 由于 2 是集合中唯一的数字，getRandom 总是返回 2 。

提示：

-2^31 <= val <= 2^31 - 1
最多调用 insert、remove 和 getRandom 函数 2 * 105 次
在调用 getRandom 方法时，数据结构中 至少存在一个 元素。
*/
func main() {
	randomizedSet := Constructor()
	fmt.Println(randomizedSet.Insert(1) == true)                                  // 向集合中插入 1 。返回 true 表示 1 被成功地插入。
	fmt.Println(randomizedSet.Remove(2) == false)                                 // 返回 false ，表示集合中不存在 2 。
	fmt.Println(randomizedSet.Insert(2) == true)                                  // 向集合中插入 2 。返回 true 。集合现在包含 [1,2] 。
	fmt.Println(randomizedSet.GetRandom() == 1 || randomizedSet.GetRandom() == 2) // getRandom 应随机返回 1 或 2 。
	fmt.Println(randomizedSet.Remove(1) == true)                                  // 从集合中移除 1 ，返回 true 。集合现在包含 [2] 。
	fmt.Println(randomizedSet.Insert(2) == false)                                 // 2 已在集合中，所以返回 false 。
	fmt.Println(randomizedSet.GetRandom() == 2)                                   // 由于 2 是集合中唯一的数字，getRandom 总是返回 2 。
}

type RandomizedSet struct {
	m   map[int]int
	arr []int
}

func Constructor() RandomizedSet {
	r := RandomizedSet{}
	r.m = make(map[int]int)
	r.arr = []int{0}
	return r
}

func (this *RandomizedSet) Insert(val int) bool {
	if this.m[val] == 0 {
		this.m[val] = len(this.arr)
		this.arr = append(this.arr, val)
		return true
	}
	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	if this.m[val] > 0 {
		i := this.m[val]
		this.m[val] = 0
		if i != len(this.arr)-1 {
			this.arr[i], this.arr[len(this.arr)-1] = this.arr[len(this.arr)-1], this.arr[i]
			this.m[this.arr[i]] = i
		}
		this.arr = this.arr[:len(this.arr)-1]
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	rand.Seed(time.Now().UnixNano())
	i := 1 + rand.Intn(len(this.arr)-1)
	return this.arr[i]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
