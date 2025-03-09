package link

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// 栈
	var stack []*ListNode
	for head != nil {
		stack = append(stack, head)
		head = head.Next
	}

	for i := len(stack) - 1; i > 0; i-- {
		stack[i].Next = stack[i-1]
		stack[i-1].Next = nil
	}
	return stack[len(stack)-1]
}

// 1,2,3,4,5
func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	var pre *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
