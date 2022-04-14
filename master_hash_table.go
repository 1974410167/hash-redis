package main

import (
	"github.com/spaolacci/murmur3"
	"math"
)

/*
Redis底层维护两个哈希表,这里我们对其进行封装，使它看起来像是在操作一个哈希表
*/

type MasterHashTable struct {
	HashTable1   *HashTable
	HashTable2   *HashTable
	counter      int  // 计数器，用于渐进式rehash
	sign         bool // 用于标注目前哈希表是否处于rehash阶段
	threshold    int  // 阈值，当前阈值和阈值一样的时候，判断一次是否需要扩容或者缩容
	cap          int
	curThreshold int // 当前阈值
}

func (m *MasterHashTable) Get(key string) any {
	if m.sign == false {
		m.handleThreshold()
	}
	return m.GetValFromHashNode(key)
}

func (m *MasterHashTable) Put(key string, val any) {
	if m.sign == false {
		m.handleThreshold()
		m.PutInHashTable(key, val, m.HashTable1)
	} else {
		m.PutInHashTable(key, val, m.HashTable2)
		m.gradualHash()

	}
}

// 获得哈希表的节点和索引

func (m *MasterHashTable) GetHashNode(key string, table *HashTable) (*HashNode, int64) {
	index := m.getIndex(key, table.cap)
	node := table.arr[index]
	return node, index
}

// 根据key从哈希表节点拿值

func (m *MasterHashTable) GetValFromHashNode(key string) any {
	if m.sign == true {
		m.gradualHash()
	}
	// 不管何种情况都先到HashTable1中查找

	curHashNode1, _ := m.GetHashNode(key, m.HashTable1)
	if curHashNode1 != nil && curHashNode1.head.exist(key) {
		return curHashNode1.head.searchInListNode(key)
	}
	// 处于渐进式哈希状态，并且没在HashTable1中找到，继续从HashTable2中查找
	if m.sign == true {
		curHashNode2, _ := m.GetHashNode(key, m.HashTable2)
		if curHashNode1 == nil && curHashNode2 == nil {
			return nil
		} else if curHashNode2 != nil && curHashNode2.head.exist(key) {
			return curHashNode2.head.searchInListNode(key)
		}
	}
	return nil
}

// 向HashTable中Put一个键值对

func (m *MasterHashTable) PutInHashTable(key string, val any, table *HashTable) {

	curHashNode, index := m.GetHashNode(key, table)
	newNode := &ListNode{
		key: key,
		val: val,
	}
	if curHashNode == nil {
		head := &ListNode{}
		tail := &ListNode{}
		head.next = tail
		tail.pre = head
		newHashNode := &HashNode{
			head: head,
			tail: tail,
		}
		newHashNode.tail.insertInListNode(newNode)
		m.SetHashNode(index, newHashNode, table)
	} else {
		// 不存在链表中，插入
		if !curHashNode.head.exist(key) {
			curHashNode.tail.insertInListNode(newNode)
		} else {
			// key在链表里，更新
			curHashNode.head.update(newNode)
		}
	}
}

func (m *MasterHashTable) gradualHash() {
	if m.sign == false {
		return
	}
	// 渐进式哈希完毕
	if m.counter >= m.HashTable1.cap {
		m.sign = false
		m.changeHashTable()
		m.curThreshold = 0
		m.counter = 0
		m.threshold = m.cap / 10
		return
	}
	hashNode := m.HashTable1.arr[m.counter]
	if hashNode != nil {
		t := hashNode.head
		for t.next != nil {
			m.PutInHashTable(t.next.key, t.next.val, m.HashTable2)
			t = t.next
		}
		m.HashTable1.arr[m.counter] = nil
	}
	m.counter += 1

}

func (m *MasterHashTable) changeHashTable() {
	m.HashTable1 = m.HashTable2
	m.HashTable2 = nil
	m.cap = m.HashTable1.cap
}

func (m *MasterHashTable) SetHashNode(index int64, node *HashNode, table *HashTable) {
	table.arr[index] = node
}

func (m *MasterHashTable) handleThreshold() {
	m.curThreshold += 1
	// 是否需要判断负载因子
	if m.curThreshold == m.threshold {
		// 是否需要扩展
		isExtend := m.HashTable1.isExtend()
		isShrink := m.HashTable1.isShrink()
		var n int
		// 是否需要扩容或者缩容
		if isShrink || isExtend {
			// 扩容
			if isExtend {
				n = m.HashTable1.getExtendNumber()
			} else if isShrink {
				n = m.HashTable2.getShrinkNumber()
			}
			newArr := make([]*HashNode, n)
			m.HashTable2 = &HashTable{
				arr: newArr,
				cap: n,
			}
			m.sign = true
		}
		m.curThreshold = 0
	}
}

// 索引
func (m *MasterHashTable) getIndex(key string, cap int) int64 {
	q := m.getHashCode(key)
	a := math.Abs(float64(q))
	res := int64(a) % int64(cap)
	return res
}

// 哈希值
func (m *MasterHashTable) getHashCode(key string) int {
	by := []byte(key)
	p := murmur3.New32()
	p.Write(by)
	res := int(p.Sum32())
	return res
}
