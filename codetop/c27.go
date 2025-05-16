package codetop

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var cntA int
	for cur := headA; cur != nil; cur = cur.Next {
		cntA++
	}
	var cntB int
	for cur := headB; cur != nil; cur = cur.Next {
		cntB++
	}
	diff := cntA - cntB
	if diff > 0 {
		for i := 0; i < diff; i++ {
			headA = headA.Next
		}
	} else {
		for i := 0; i < -diff; i++ {
			headB = headB.Next
		}
	}
	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}
	return headA
}
