package main

import f "fmt"

/*
定义双向链表，并实现一系列操作双向链表的函数
*/

type ListNode struct {
	key  string
	val  any
	next *ListNode
	pre  *ListNode
}

// 判断key是否存在于链表
func (node *ListNode) exist(key string) bool {
	t := node
	for t.next != nil {
		if t.next.key == key {
			return true
		}
		t = t.next
	}
	return false
}

// 新节点插入链表
func (node *ListNode) insertInListNode(newNode *ListNode) {
	newNode.pre = node.pre
	node.pre.next = newNode
	node.pre = newNode
	newNode.next = node
}

// 更新相同key对应的val
func (node *ListNode) update(newNode *ListNode) {
	t := node.next
	for t != nil {
		if t.key == newNode.key {
			t.val = newNode.val
		}
		t = t.next
	}
}

// 根据key找到val,key不存在返回nil
func (node *ListNode) searchInListNode(key string) any {
	t := node
	for t.next != nil {
		if t.next.key == key {
			return t.next.val
		}
		t = t.next
	}
	return nil
}

// 打印链表
func (node *ListNode) printListNode() {
	t := node
	for t.next != nil {
		f.Println(t.next.val)
		t = t.next
	}
}

// 得到链表长度
func (node *ListNode) listNodeLength() int {
	t := node
	n := 0
	for t.next != nil {
		n += 1
		t = t.next
	}
	return n
}
