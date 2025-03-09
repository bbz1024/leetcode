package link

type node struct {
	val  int
	next *node
	prev *node
}

func removeElements(head *node, val int) *node {
	dummy := &node{next: head}
	cur := dummy
	for cur != nil && cur.next != nil {
		if cur.next.val == val {
			cur.next = cur.next.next
		} else {
			cur = cur.next
		}
	}
	return dummy.next
}
