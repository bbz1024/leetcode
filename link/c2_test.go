package link

import (
	"testing"
)

func TestMyLinkedList(t *testing.T) {
	list := Constructor()
	list.AddAtHead(1)
	list.AddAtTail(3)
	list.AddAtIndex(1, 2)
	t.Log(list.Get(1))
	list.DeleteAtIndex(1)
	t.Log(list.Get(1))
}
