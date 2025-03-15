package tree

import "strconv"

func binaryTreePaths(root *TreeNode) []string {
	var res []string
	var dfs func(node *TreeNode, path string)
	dfs = func(node *TreeNode, path string) {
		if node != nil {
			// 到达叶子节点
			if node.Left == nil && node.Right == nil {
				res = append(res, path+strconv.Itoa(node.Val))
				return
			}
			//
			if node.Left != nil {
				dfs(node.Left, path+strconv.Itoa(node.Val)+"->")
			}
			if node.Right != nil {
				dfs(node.Right, path+strconv.Itoa(node.Val)+"->")
			}
		}
	}
	dfs(root, "")
	return res
}
