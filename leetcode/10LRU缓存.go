package leetcode

/*
设计和构建一个“最近最少使用”缓存，该缓存会删除最近最少使用的项目。缓存应该从键映射到值(允许你插入和检索特定键对应的值)，并在初始化时指定最大容量。当缓存被填满时，它应该删除最近最少使用的项目。

它应该支持以下操作： 获取数据 get 和 写入数据 put 。

获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。

示例:

LRUCache cache = new LRUCache( 2  );

//cache.put(1, 1);
//cache.put(2, 2);
//cache.get(1);       // 返回  1
//cache.put(3, 3);    // 该操作会使得密钥 2 作废
//cache.get(2);       // 返回 -1 (未找到)
//cache.put(4, 4);    // 该操作会使得密钥 1 作废
//cache.get(1);       // 返回 -1 (未找到)
//cache.get(3);       // 返回  3
//cache.get(4);       // 返回  4
解题思路
map+双向链表，将每次使用的节点移动到链表头部，则尾部节点就是在容量满时，需要淘汰的节点。
*/
type LRUCacheNode struct {
	key   int
	value int
	pre   *LRUCacheNode
	next  *LRUCacheNode
}

// 双向链表+map
type LRUCache struct {
	md       map[int]*LRUCacheNode
	len      int
	capacity int
	head     *LRUCacheNode
	tail     *LRUCacheNode
}

func Constructor(capacity int) LRUCache {
	lc := LRUCache{
		md:       make(map[int]*LRUCacheNode),
		len:      0,
		capacity: capacity,
	}
	lc.head = &LRUCacheNode{}
	lc.tail = &LRUCacheNode{}

	lc.head.pre = nil
	lc.head.next = lc.tail

	lc.tail.pre = lc.head
	lc.tail.next = nil

	return lc
}

func (this *LRUCache) Get(key int) int {
	// 不存在，就返回-1
	if _, exist := this.md[key]; !exist {
		return -1
	}
	// 存在就添加到头部，返回值
	this.AddToHead(this.md[key])
	return this.md[key].value
}

func (this *LRUCache) Put(key int, value int) {
	// 先判断是否存在，存在则更新，并且移到头部
	if _, exist := this.md[key]; exist {
		// 移除该节点
		this.RemoveNode(this.md[key])
		// 添加到头部
	} else { // 不存在，则直接添加到头部
		// 先判断容量是否满
		if this.len > this.capacity {
			panic("capacity err")
		}
		// 已经满，先剔除尾部，再添加到头部
		if this.len == this.capacity {
			// 剔除尾部
			this.RemoveTail(this.md[key])
			this.len--
			// 添加到头部
			newNode := &LRUCacheNode{
				key:   key,
				value: value,
				pre:   nil,
				next:  nil,
			}
			this.AddToHead(newNode)
			this.md[key] = newNode
			this.len++
		} else { // 未满，直接添加到头部
			// 添加到头部
		}
	}
}

// 添加节点到头部
func (this *LRUCache) AddToHead(node *LRUCacheNode) {
	node.next = this.head.next
	node.pre = this.head
	this.head.next.pre = node
	this.head.next = node
}

// 移动到头部
func (this *LRUCache) MoveToHead(node *LRUCacheNode) {
	this.RemoveNode(node)
	this.AddToHead(node)
}

// 移除当前节点
func (this *LRUCache) RemoveNode(node *LRUCacheNode) {
	delete(this.md, node.key) // 在map删除
	// 删除链表节点
	node.pre.next = node.next
	node.next.pre = node.pre

	node.pre = nil
	node.next = nil
}

// 移除尾部节点
func (this *LRUCache) RemoveTail(node *LRUCacheNode) {
	if this.len > 0 {
		this.RemoveNode(this.tail.pre)
	}
}
