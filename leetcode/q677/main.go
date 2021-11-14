package main

import (
	"fmt"
)

/*
https://leetcode-cn.com/problems/map-sum-pairs/

677. 键值映射
实现一个 MapSum 类，支持两个方法，insert 和 sum：

MapSum() 初始化 MapSum 对象
void insert(String key, int val) 插入 key-val 键值对，字符串表示键 key ，整数表示值 val 。如果键 key 已经存在，那么原来的键值对将被替代成新的键值对。
int sum(string prefix) 返回所有以该前缀 prefix 开头的键 key 的值的总和。
*/
func main() {
	mapSum := Constructor()
	mapSum.Insert("apple", 3)
	fmt.Println(mapSum.Sum("ap"))
	mapSum.Insert("app", 2)
	fmt.Println(mapSum.Sum("ap"))
}

type MapSum struct {
	Map map[string]int
}

func Constructor() MapSum {
	m := MapSum{}
	m.Map = make(map[string]int)
	return m
}

func (this *MapSum) Insert(key string, val int) {
	this.Map[key] = val
}

func (this *MapSum) Sum(prefix string) int {
	sum := 0
	for k, v := range this.Map {
		if len(k) >= len(prefix) && k[:len(prefix)] == prefix {
			sum += v
		}
	}
	return sum
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
