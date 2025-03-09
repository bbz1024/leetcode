package link

func detectCycle(head *ListNode) *ListNode {
	// 哈希表做法
	var mp = make(map[*ListNode]struct{})
	for head != nil {
		_, ok := mp[head]
		if ok {
			// 环
			return head
		}
		//
		mp[head] = struct{}{}
		head = head.Next
	}
	return nil
}
func detectCycle2(head *ListNode) *ListNode {
	/*
		TODO m
	*/
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		// 碰撞了
		if fast == slow {
			index := head
			index2 := slow
			for index != index2 {
				index = index.Next
				index2 = index2.Next
			}
			return index
		}
	}
	return nil
}
