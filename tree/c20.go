package tree

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	dummy := root
	for dummy != nil {
		if dummy.Val > val {
			if dummy.Left == nil {
				dummy.Left = &TreeNode{Val: val}
				break
			}
			dummy = dummy.Left
		} else {
			if dummy.Right == nil {
				dummy.Right = &TreeNode{Val: val}
				break
			}
			dummy = dummy.Right
		}

	}
	return root
}
