package tree

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
	minVal := math.MaxInt
	var pre *TreeNode
	var getMin func(root *TreeNode)
	getMin = func(root *TreeNode) {
		if root == nil {
			return
		}
		getMin(root.Left)
		if pre != nil {
			minVal = min(minVal, root.Val-pre.Val)
		}
		pre = root // 记录上一个节点，对于中来说是左节点，对于右来说是中节点
		getMin(root.Right)
	}
	getMin(root)
	return minVal
}
