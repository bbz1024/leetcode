package link

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	cur := dummy

	for cur.Next != nil && cur.Next.Next != nil {
		n1 := cur.Next
		n2 := cur.Next.Next
		cur.Next = n2
		n1.Next = n2.Next
		n2.Next = n1
		cur = n1
	}
	return dummy.Next
}

/*
1234
13
*/
