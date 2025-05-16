package codetop

import "testing"

func TestReorderList(t *testing.T) {
	list := &ListNode{
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
	reorderList(list)
	for list != nil {
		t.Log(list.Val)
		list = list.Next
	}
}
