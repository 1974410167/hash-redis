# hash-redis

通读«Redis设计与实现»哈希表篇章来用 Golang重现Redis的底层哈希表

详见 http://ghyuan.cn/post/64

哈希表几个要素
- 哈希函数
- 如何处理哈希冲突
- 底层数组基于何种策略扩容缩容
- Rehash方法

redis的哈希表的底层实现
- 使用Murmurhash3哈希算法
- 使用链地址法处理哈希冲突
- 负载因子大于1，扩容至第一个大于等于used*2的2次方幂。负载因子小于0.1，缩容至第一个大于等于used的二次方幂
- 维护两个底层哈希表，并以渐进式rehash策略来实现


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
