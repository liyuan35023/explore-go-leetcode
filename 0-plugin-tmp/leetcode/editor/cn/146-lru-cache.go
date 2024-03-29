package cn

//运用你所掌握的数据结构，设计和实现一个 LRU (最近最少使用) 缓存机制 。
//
// 实现 LRUCache 类：
//
// LRUCache(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
// int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
// void put(int key, int value) 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。当缓存容量达到上
//限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
//
// 进阶：你是否可以在 O(1) 时间复杂度内完成这两种操作？
//
// 示例：
//
//输入
//["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
//[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
//输出
//[null, null, null, 1, null, -1, null, -1, 3, 4]
//
//解释
//LRUCache lRUCache = new LRUCache(2);
//lRUCache.put(1, 1); // 缓存是 {1=1}
//lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
//lRUCache.get(1);    // 返回 1
//lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
//lRUCache.get(2);    // 返回 -1 (未找到)
//lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
//lRUCache.get(1);    // 返回 -1 (未找到)
//lRUCache.get(3);    // 返回 3
//lRUCache.get(4);    // 返回 4
//
// 提示：
//
// 1 <= capacity <= 3000
// 0 <= key <= 10000
// 0 <= value <= 105
// 最多调用 2 * 105 次 get 和 put
//
//leetcode submit region begin(Prohibit modification and deletion)

type dequeue struct {
	key, val int
	pre, next *dequeue
}

type LRUCache struct {
	capacity int
	cache map[int]*dequeue
	dummyHead, dummyTail *dequeue
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity:  capacity,
		cache:     make(map[int]*dequeue),
		dummyHead: new(dequeue),
		dummyTail: new(dequeue),
	}
	lru.dummyHead.next = lru.dummyTail
	lru.dummyTail.pre = lru.dummyHead
	return lru
}

func (this *LRUCache) moveToHead(d *dequeue) {
	if d.pre != nil {
		d.pre.next = d.next
		d.next.pre = d.pre
	}
	d.pre = this.dummyHead
	d.next = this.dummyHead.next
	this.dummyHead.next.pre = d
	this.dummyHead.next = d
}

func (this *LRUCache) Get(key int) int {
	if d, ok := this.cache[key]; ok {
		// move to head
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
		newNode := &dequeue{
			key: key,
			val: value,
		}
		this.moveToHead(newNode)
		this.cache[key] = newNode
		if len(this.cache) > this.capacity {
			// move tail
			tail := this.dummyTail.pre
			delete(this.cache, tail.key)
			tail.pre.next = this.dummyTail
			this.dummyTail.pre = tail.pre
		}
	}
}




//type Deque struct {
//	Key int
//	Val  int
//	Pre  *Deque
//	Next *Deque
//}
//
//type LRUCache struct {
//	capacity             int
//	cache                map[int]*Deque
//	dummyHead, dummyTail *Deque
//}
//
//func Constructor(capacity int) LRUCache {
//	lru := LRUCache{
//		capacity:  capacity,
//		cache:     make(map[int]*Deque),
//		dummyHead: new(Deque),
//		dummyTail: new(Deque),
//	}
//	lru.dummyHead.Next = lru.dummyTail
//	lru.dummyTail.Pre = lru.dummyHead
//	return lru
//}
//
//func (this *LRUCache) Get(key int) int {
//	if d, ok := this.cache[key]; ok {
//		this.moveToHead(d)
//		return d.Val
//	} else {
//		return -1
//	}
//}
//
//func (this *LRUCache) Put(key int, value int) {
//	if d, ok := this.cache[key]; ok {
//		d.Val = value
//	} else {
//		if len(this.cache) == this.capacity {
//			// delete tail
//			delete(this.cache, this.dummyTail.Pre.Key)
//			this.dummyTail.Pre.Pre.Next = this.dummyTail
//			this.dummyTail.Pre = this.dummyTail.Pre.Pre
//		}
//		this.cache[key] = &Deque{Key: key, Val: value}
//	}
//	this.moveToHead(this.cache[key])
//}
//
//func (this *LRUCache) moveToHead(d *Deque) {
//	if d.Pre != nil {
//		d.Pre.Next = d.Next
//	}
//	if d.Next != nil {
//		d.Next.Pre = d.Pre
//	}
//	d.Next = this.dummyHead.Next
//	this.dummyHead.Next.Pre = d
//	d.Pre = this.dummyHead
//	this.dummyHead.Next = d
//}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)
