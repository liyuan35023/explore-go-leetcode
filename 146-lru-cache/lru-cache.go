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
type LRUCache struct {
	capacity int
	// head 表示最新的值  head -> <- node -> <- tail
	dummyHead, dummyTail *Dequeue
	cache map[int]*Dequeue
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		cache: make(map[int]*Dequeue),
		dummyHead: new(Dequeue),
		dummyTail: new(Dequeue),
	}
	l.dummyTail.pre = l.dummyHead
	l.dummyHead.next = l.dummyTail
	return l
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToHead(node)
		return node.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int)  {
	if node, ok := this.cache[key]; ok {
		this.moveToHead(node)
		node.value = value
	} else {
		node = &Dequeue{key: key, value: value}
		this.moveToHead(node)
		this.cache[key] = node
		if len(this.cache) > this.capacity {
			// delete tail
			tail := this.dummyTail.pre
			delete(this.cache, tail.key)
			this.dummyTail.pre = tail.pre
			tail.pre.next = this.dummyTail
		}
	}
}

func (this *LRUCache) moveToHead(node *Dequeue) {
	if node.pre != nil {
		node.pre.next = node.next
		node.next.pre = node.pre
	}
	node.pre = this.dummyHead
	node.next = this.dummyHead.next
	this.dummyHead.next.pre = node
	this.dummyHead.next = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

type Dequeue struct {
	key, value int
	pre, next *Dequeue
}
