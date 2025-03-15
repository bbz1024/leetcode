package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	left := sortedArrayToBST(nums[:mid])
	right := sortedArrayToBST(nums[mid+1:])
	// 后续遍历，根节点就是栈的最后一个元素
	return &TreeNode{
		Val:   nums[mid],
		Left:  left,
		Right: right,
	}
}
