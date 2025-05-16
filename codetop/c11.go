package codetop

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	q := []*TreeNode{root}
	for len(q) > 0 {
		size := len(q)
		var curLevel []int

		for size > 0 {
			node := q[0]
			q = q[1:]
			curLevel = append(curLevel, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
			size--
		}

		res = append(res, curLevel)
	}
	return res
}
