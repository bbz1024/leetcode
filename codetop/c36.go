package codetop

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head, Val: -1}
	cur := dummy.Next
	pre := dummy
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			for cur.Next != nil && cur.Val == cur.Next.Val {
				cur = cur.Next
			}
			cur = cur.Next
			pre.Next = cur
		} else {
			pre = cur
			cur = cur.Next
		}
	}
	return dummy.Next
}
