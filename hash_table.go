package main

type HashTable struct {
	arr        []*HashNode // 底层数组
	LoadFactor float64     // 负载因子
	cap        int         // 底层数组大小
}

// 得到目前哈希表中哈希键的总数量
func (table *HashTable) getHashKeyTotal() int {
	t := table.arr
	total := 0
	for i := 0; i < table.cap; i++ {
		if t[i] != nil {
			curHashNode := t[i].head
			if curHashNode != nil {
				q := curHashNode.listNodeLength()
				total += q
			}
		}
	}
	return total
}

// 更新哈希表负载因子
func (table *HashTable) updateLoadFactor(loadFactor float64) {
	table.LoadFactor = loadFactor
}

// 得到目前负载因子
func (table *HashTable) getLoadFactor() float64 {
	curNum := table.getHashKeyTotal()
	loadFactor := float64(curNum) / float64(table.cap)
	table.updateLoadFactor(loadFactor)
	return loadFactor
}

// 是否需要扩展
func (table *HashTable) isExtend() bool {
	loadFactor := table.getLoadFactor()
	table.LoadFactor = loadFactor
	if loadFactor >= 1 {
		return true
	}
	return false
}

// 是否需要收缩
func (table *HashTable) isShrink() bool {
	loadFactor := table.getLoadFactor()
	table.LoadFactor = loadFactor
	if loadFactor < 0.1 {
		return true
	}
	return false
}

// 扩展的话扩展至多大? 扩展至第一个大于等于used*2的2次方幂
func (table *HashTable) getExtendNumber() int {
	used := table.getHashKeyTotal()
	t := 2
	for t < used*2 {
		t = t * 2
	}
	return t
}

// 收缩的话收缩至多大？收缩至第一个大于等于used的2次方幂
func (table *HashTable) getShrinkNumber() int {
	used := table.getHashKeyTotal()
	t := 2
	for t < used {
		t = t * 2
	}
	return t
}
