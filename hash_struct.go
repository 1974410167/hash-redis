package main

import (
	f "fmt"
)

/*

定义哈希表以及哈希表的节点以及实现操作哈希表的函数

*/

type HashNode struct {
	head *ListNode
	tail *ListNode
}

type Hash struct {
	Table *MasterHashTable
}

// 初始化哈希表
func (h *Hash) init() {
	cap := 64
	arr := make([]*HashNode, cap)
	newHashTable := &HashTable{
		arr: arr,
		cap: cap,
	}
	h.Table = &MasterHashTable{
		HashTable1:   newHashTable,
		threshold:    cap / 10,
		curThreshold: 0,
		cap:          cap,
	}
}

// Get 从哈希表拿值,
func (h *Hash) Get(key string) any {
	return h.Table.Get(key)

}

// Put 向哈希表中Put一个键值对，如果key存在则更新
func (h *Hash) Put(key string, val any) {
	h.Table.Put(key, val)
}

func main() {
	hash := &Hash{}
	hash.init()
	hash.Put("age", 100)
	hash.Put("name", "葛浩源")
	hash.Put("address", "东炉村")
	hash.Put("gender", "男")
	q1 := hash.Get("address")
	f.Println(q1)
	hash.Put("address", "啦啦啦")
	q2 := hash.Get("address")
	f.Println(q2)

	q3 := hash.Get("age")
	f.Println(q3)
}
