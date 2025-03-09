package link

type MyLinkedList struct {
	head *node
	tail *node
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if index == 0 {
		if this.head == nil {
			return -1
		}
		return this.head.val
	}
	return this.getNode(index).val
}
func (this *MyLinkedList) getNode(index int) *node {
	if index == 0 {
		return this.head
	}
	nd := this.head
	for i := 0; i < index; i++ {
		if nd == nil {
			return nil
		}
		nd = nd.next
	}
	return nd
}
func (this *MyLinkedList) AddAtHead(val int) {
	node := &node{val: val}
	if this.head == nil {
		this.head = node
		this.tail = node
		return
	}
	hd := this.head
	node.next = hd
	hd.prev = node
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &node{val: val}
	if this.head == nil {
		this.head = node
		this.tail = node
		return
	}
	tl := this.tail
	tl.next = node
	node.prev = tl
	this.tail = node
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	nd := this.getNode(index)
	if nd == nil {
		return
	}
	newnode := &node{val: val}
	pre := nd.prev
	pre.next = newnode
	newnode.prev = pre
	nd.prev = newnode
	newnode.next = nd
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index == 0 {
		this.head = this.head.next
		return
	}
	node := this.getNode(index)
	if node == nil {
		return
	}
	pre := node.prev
	pre.next = node.next
	node.next.prev = pre
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
