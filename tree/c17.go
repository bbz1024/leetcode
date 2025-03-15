package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
	var maxCnt int
	var cnt int
	var res []int
	var pre *TreeNode
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if pre == nil {
			cnt = 1
		} else if pre.Val == root.Val {
			cnt++
		} else {
			cnt = 1
		}
		pre = root
		if cnt == maxCnt {
			res = append(res, root.Val)
		} else if cnt > maxCnt {
			maxCnt = cnt
			res = []int{root.Val}
		}
		dfs(root.Right)
	}
	dfs(root)
	return res

}
