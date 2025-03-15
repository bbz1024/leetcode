package tree

func countNodes(root *TreeNode) int {
	// 通过前序遍历
	if root == nil {
		return 0
	}

	return countNodes(root.Left) + countNodes(root.Right) + 1
}
