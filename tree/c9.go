package tree

func findBottomLeftValue(root *TreeNode) int {
	var left int
	var maxDepth int
	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if depth > maxDepth {
			maxDepth = depth
			left = root.Val
		}
		dfs(root.Left, depth+1)
		dfs(root.Right, depth+1)
	}
	dfs(root, 1)
	return left
}
