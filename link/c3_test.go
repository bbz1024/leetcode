package link

import "testing"

func TestReverseList(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}
	hd := reverseList(head)
	for hd != nil {
		t.Log(hd.Val)
		hd = hd.Next
	}
}
