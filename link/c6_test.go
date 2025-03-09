package link

import "testing"

func TestDetectCycle(t *testing.T) {
}
func TestDetectCycle2(t *testing.T) {
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
	//[1,2,3,4]
	// pos 2
	head.Next.Next.Next.Next = head.Next
	println(detectCycle2(head).Val)
}
