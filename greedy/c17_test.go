package greedy

import "testing"

func TestMinCameraCover(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	t.Log(minCameraCover(tree))
}
