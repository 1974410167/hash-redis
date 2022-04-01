package main

import (
	f "fmt"
	"github.com/spaolacci/murmur3"
	"math"
)

/*

定义哈希表以及哈希表的节点以及实现操作哈希表的函数

*/

type HashNode struct {
	head *ListNode
	tail *ListNode
}

type Hash struct {
	HashTable1   *HashTable
	HashTable2   *HashTable
	cap          int
	counter      int  // 计数器，用于渐进式rehash
	sign         bool // 用于标注目前哈希表是否处于rehash阶段
	threshold    int  // 阈值，当前阈值和阈值一样的时候，判断一次是否需要扩容或者缩容
	curThreshold int  // 当前阈值

}

// 初始化哈希表
func (h *Hash) init() {
	h.cap = 64
	h.threshold = h.cap / 10
	h.curThreshold = 0
	arr1 := make([]*HashNode, h.cap)
	h.HashTable1 = &HashTable{
		arr1: arr1,
		cap:  h.cap,
	}
}

// Get 从哈希表拿值,
func (h *Hash) Get(key string) any {

	index := h.getIndex(key)
	node := h.HashTable1.arr1[index]
	if node == nil {
		return nil
	} else {
		res := node.head.searchInListNode(key)
		return res
	}
}

// Put 向哈希表中Put一个键值对，如果key存在则更新
func (h *Hash) Put(key string, val any) {
	index := h.getIndex(key)
	hashNode := h.HashTable1.arr1[index]
	newNode := &ListNode{
		key: key,
		val: val,
	}
	// 索引节点处为空  也就是没有哈希冲突
	if hashNode == nil {
		head := &ListNode{}
		tail := &ListNode{}
		head.next = tail
		tail.pre = head
		h.HashTable1.arr1[index] = &HashNode{
			head: head,
			tail: tail,
		}
		hashNode = h.HashTable1.arr1[index]
		hashNode.tail.insertInListNode(newNode)
		return
	}
	// 有哈希冲突
	//key不存在链表里，直接插入链表末尾
	if !hashNode.head.exist(key) {
		hashNode.tail.insertInListNode(newNode)
	} else {
		// key在链表里，更新
		hashNode.head.update(newNode)
	}

}

// 索引
func (h *Hash) getIndex(key string) int64 {
	q := h.getHashCode(key)
	a := math.Abs(float64(q))
	cap := h.HashTable1.cap
	res := int64(a) % int64(cap)
	return res
}

// 哈希值
func (h *Hash) getHashCode(key string) int {
	by := []byte(key)
	p := murmur3.New32()
	p.Write(by)
	res := int(p.Sum32())
	return res
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
