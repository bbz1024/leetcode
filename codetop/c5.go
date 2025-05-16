package codetop

func reverseKGroup(head *ListNode, k int) *ListNode {
	var cnt int
	for cur := head; cur != nil; cur = cur.Next {
		cnt++
	}

	var pre *ListNode
	dummy := &ListNode{Next: head}
	p0 := dummy
	cur := p0.Next

	for i := 0; i < cnt/k; i++ {
		for j := 0; j < k; j++ {
			next := cur.Next
			cur.Next = pre
			pre = cur
			cur = next
		}
		next := p0.Next
		p0.Next.Next = cur // 头部第一个指向尾部第一个
		p0.Next = pre      // 将头部指向尾部
		p0 = next          // 移动到下一个头部
	}
	return dummy.Next
}
