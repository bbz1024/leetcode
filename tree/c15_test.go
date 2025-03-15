package tree

import (
	"fmt"
	"testing"
)

func TestIsValidBST(t *testing.T) {
	bst := isValidBST(&TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 5,
		},
		Right: &TreeNode{
			Val: 3,
		},
	})
	fmt.Println(bst)
}
