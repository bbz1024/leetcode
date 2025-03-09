package link

import "testing"

func TestRemoveNthFromEnd(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}},
			},
		},
	}
	hd := removeNthFromEnd(head, 5)
	for hd != nil {
		t.Log(hd.Val)
		hd = hd.Next
	}
}
