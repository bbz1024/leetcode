package codetop

import "testing"

func TestReverseBetween(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	reverseBetween(head, 2, 4)
	for head != nil {
		t.Log(head.Val)
		head = head.Next
	}
}
