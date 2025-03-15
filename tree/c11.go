package tree

import "slices"

/*
in order: left, root, right
9,3,15,20,7

post order: left, right, root
9,15,7,20,3
*/
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	index := slices.Index(inorder, postorder[len(postorder)-1])
	inLeft := inorder[:index]
	inRight := inorder[index+1:]
	postLeft := postorder[:len(inLeft)]
	postRight := postorder[len(inLeft) : len(postorder)-1]
	node := &TreeNode{
		Val:   postorder[len(postorder)-1],
		Left:  buildTree(inLeft, postLeft),
		Right: buildTree(inRight, postRight),
	}
	return node
}
