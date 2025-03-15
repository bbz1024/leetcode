package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor3(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		val := root.Val
		if p.Val > val && q.Val > val {
			root = root.Right
		} else if p.Val < val && q.Val < val {
			root = root.Left
		} else {
			return root
		}
	}
	return root
}
