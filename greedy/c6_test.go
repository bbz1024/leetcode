package greedy

import "testing"

//2,3,1,1,4

func TestJump(t *testing.T) {
	t.Log(jump([]int{7, 0, 9, 6, 9, 6, 1, 7, 9, 0, 1, 2, 9, 0, 3}))
}
