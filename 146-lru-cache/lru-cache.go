package _46_lru_cache

/*
*	运用你所掌握的数据结构，设计和实现一个 LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
*
*	获取数据 get(key) - 如果关键字 (key) 存在于缓存中，则获取关键字的值（总是正数），否则返回 -1。
*	写入数据 put(key, value) - 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字/值」。
*	当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
*
*	进阶:
*
*	你是否可以在 O(1) 时间复杂度内完成这两种操作？
*
*	示例:
*
*	LRUCache cache = new LRUCache( 2  缓存容量  )
*	cache.put(1, 1);
*	cache.put(2, 2);
*	cache.get(1);       // 返回  1
*	cache.put(3, 3);    // 该操作会使得关键字 2 作废
*	cache.get(2);       // 返回 -1 (未找到)
*	cache.put(4, 4);    // 该操作会使得关键字 1 作废
*	cache.get(1);       // 返回 -1 (未找到)
*	cache.get(3);       // 返回  3
*	cache.get(4)       // 返回  4
*
*/
type Dequeue struct {
	key, val int
	Pre, Next *Dequeue
}

type LRUCache struct {
	cache map[int]*Dequeue
	capacity int
	dummyHead, dummyTail *Dequeue
}


func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		cache:     make(map[int]*Dequeue),
		capacity:  capacity,
		dummyHead: new(Dequeue),
		dummyTail: new(Dequeue),
	}
	lru.dummyHead.Next = lru.dummyTail
	lru.dummyTail.Pre = lru.dummyHead
	return lru
}

func (this *LRUCache) moveToHead(d *Dequeue) {
	if d.Pre != nil {
		d.Pre.Next = d.Next
		d.Next.Pre = d.Pre
	}
	d.Pre = this.dummyHead
	d.Next = this.dummyHead.Next
	this.dummyHead.Next.Pre = d
	this.dummyHead.Next = d
}

func (this *LRUCache) Get(key int) int {
	if d, ok := this.cache[key]; ok {
		this.moveToHead(d)
		return d.val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int)  {
	if d, ok := this.cache[key]; ok {
		d.val = value
		this.moveToHead(d)
	} else {
		if len(this.cache) == this.capacity {
			// deleteTail
			delete(this.cache, this.dummyTail.Pre.key)
			this.dummyTail.Pre.Pre.Next = this.dummyTail
			this.dummyTail.Pre = this.dummyTail.Pre.Pre
		}
		dequeue := &Dequeue{key: key, val: value}
		this.cache[key] = dequeue
		this.moveToHead(dequeue)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
