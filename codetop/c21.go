package codetop

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}
	if left == right {
		return head
	}
	dummy := &ListNode{Next: head}
	first := dummy
	for i := 0; i < left-1; i++ {
		first = first.Next
	}
	var pre *ListNode
	cur := first.Next
	temp := cur
	for i := 0; i < right-left+1; i++ {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	first.Next = pre
	temp.Next = cur
	return dummy.Next
}
