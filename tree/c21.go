package tree

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		// 左不空，右为空
		if root.Left != nil && root.Right == nil {
			return root.Left
		} else if root.Left == nil && root.Right != nil {
			return root.Right
		} else if root.Left == nil && root.Right == nil {
			return nil
		} else {
			// 左右都不为空
			node := root.Right
			for node.Left != nil {
				node = node.Left
			}
			node.Left = root.Left
			return root.Right
		}

	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
	}
	return root

}
