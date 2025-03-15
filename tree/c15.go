package tree

import "math"

func isValidBST(root *TreeNode) bool {
	// 中序遍历，搜索树就是待有序的
	if root == nil {
		return true
	}
	pre := math.MinInt64
	var inOrder func(node *TreeNode)
	order := true
	inOrder = func(node *TreeNode) {
		if node == nil || !order {
			return
		}
		inOrder(node.Left)
		if pre >= node.Val {
			order = false
			return
		}
		pre = node.Val
		inOrder(node.Right)
	}
	inOrder(root)
	return order
}
