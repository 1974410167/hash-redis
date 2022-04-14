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

func (h *Hash) Get(key string) any {
	return h.Table.Get(key)

}

func (h *Hash) Put(key string, val any) {
	h.Table.Put(key, val)
}

func main() {
	hash := &Hash{}
	hash.init()
	hash.Put("age", 100)
	hash.Put("name", "嘻嘻")
	hash.Put("address", "东炉村")
	hash.Put("gender", "男")
	q1 := hash.Get("address")
	f.Println(q1)
	hash.Put("address", "啦啦啦")
	q2 := hash.Get("address")
	f.Println(q2)
	q3 := hash.Get("age")
	f.Println(q3)

	// 以下为测试数据
	// Put进666666个测试数据
	for i := 0; i < 666666; i++ {
		a := i
		hash.Put(string(a), a*10)
	}
	qq := hash.Get(string(123))
	f.Println(qq)
	f.Println(hash.Get("address"))
	qq1 := hash.Get(string(12334))
	f.Println(qq1) // 预计返回123340
	qq2 := hash.Get(string(66665))
	f.Println(qq2) // 预计返回666650
	qq3 := hash.Get(string(666651))
	f.Println(qq3) // 预计返回6666510
}
