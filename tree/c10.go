package tree

func hasPathSum(root *TreeNode, targetSum int) bool {
	var dfs func(root *TreeNode, sum int) bool
	dfs = func(root *TreeNode, sum int) bool {
		if root == nil {
			return false
		}
		if root.Left == nil && root.Right == nil {
			return sum+root.Val == targetSum
		}
		return dfs(root.Left, sum+root.Val) || dfs(root.Right, sum+root.Val)

	}
	return dfs(root, 0)
}
