package codetop

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode
	pre, cur := nil, head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}
