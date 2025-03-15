package tree

import "math"

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	return math.Abs(float64(left-right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}
