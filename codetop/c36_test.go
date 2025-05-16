package codetop

import "testing"

func TestDeleteDuplicates(t *testing.T) {
	// 1 2 3 3 4 4 5
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
		},
	}

	deleteDuplicates(head)
}
