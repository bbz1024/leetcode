package tree

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	maxIdx := 0
	for idx, num := range nums {
		if num > nums[maxIdx] {
			maxIdx = idx
		}
	}
	leftNums := nums[:maxIdx]
	rightNums := nums[maxIdx+1:]
	node := &TreeNode{
		Val: nums[maxIdx],
	}
	node.Left = constructMaximumBinaryTree(leftNums)
	node.Right = constructMaximumBinaryTree(rightNums)
	return node
}
func constructMaximumBinaryTree2(nums []int) *TreeNode {
	return nil
}
