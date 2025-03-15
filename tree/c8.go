package tree

func sumOfLeftLeaves(root *TreeNode) int {
	var dfs func(root *TreeNode, isLeft bool) int
	dfs = func(root *TreeNode, isLeft bool) int {
		if root == nil {
			return 0
		}
		if root.Left == nil && root.Right == nil && isLeft {
			return root.Val
		}
		return dfs(root.Left, true) + dfs(root.Right, false)
	}

	return dfs(root, false)
}
