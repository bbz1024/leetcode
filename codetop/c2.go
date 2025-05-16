package codetop

type LRUCache struct {
	capacity int
	size     int
	head     *LinkNode
	tail     *LinkNode
	mp       map[int]*LinkNode
}
type LinkNode struct {
	key   int
	value int
	next  *LinkNode
	prev  *LinkNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		head:     nil,
		tail:     nil,
		mp:       make(map[int]*LinkNode, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	if this.head == nil {
		return -1
	}
	node, ok := this.mp[key]
	if !ok {
		return -1
	}
	// 无需更新状态
	if this.size == 1 || node == this.head {
		return node.value
	}
	// 并非头结点，需要更新状态
	// 删除该节点移动到头结点
	// 如果是尾节点
	if node.next == nil {
		node.prev.next = nil
		this.tail = node.prev
	}
	hd := this.head
	this.head = node
	node.prev = nil
	node.next = hd
	hd.prev = node
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if this.size == 0 {
		this.head = &LinkNode{
			key:   key,
			value: value,
			next:  nil,
			prev:  nil,
		}
		this.tail = this.head
		this.mp[key] = this.head
		this.size++
		return
	}
	node, ok := this.mp[key]
	if !ok {
		// 插入新节点
		if this.size == this.capacity {
			// 删除尾节点
			delete(this.mp, this.tail.key)
			this.tail.prev.next = nil
			this.size--
		}
		newNode := &LinkNode{
			key:   key,
			value: value,
			next:  nil,
			prev:  nil,
		}
		this.mp[key] = newNode

		// 插入到头结点
		hd := this.head
		this.head = newNode
		newNode.next = hd
		hd.prev = newNode
		this.size++
		return
	}
	node.value = value
	if this.size == 1 || node == this.head {
		return
	}
	hd := this.head
	this.head = node
	node.prev = nil
	node.next = hd
	hd.prev = node
	return
}
