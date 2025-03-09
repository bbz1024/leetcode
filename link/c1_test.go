package link

import "testing"

func TestRemoveElements(t *testing.T) {
	head := &node{
		val: 1,
		next: &node{
			val: 2,
			next: &node{
				val: 6,
				next: &node{
					val: 3,
					next: &node{
						val: 4,
						next: &node{
							val: 5,
						},
					},
				},
			},
		},
	}
	removeElements(head, 6)
	for head != nil {
		t.Log(head.val)
		head = head.next
	}
}
