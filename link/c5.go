package link

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 0,1,2,3,4,5
	dummy := &ListNode{Next: head}
	fast := dummy
	slow := dummy
	// 快指针先走n
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	if fast == nil {
		return head.Next
	}
	for fast != nil && fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	if slow == dummy {
		return head.Next
	}
	slow.Next = slow.Next.Next
	return head
}
