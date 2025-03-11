package queue_stack

import (
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		nums := []int{1, 1, 1, 2, 2, 3}
		k := 2
		res := topKFrequent(nums, k)
		t.Log(res)
	})
}
