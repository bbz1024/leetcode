package codetop

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	var queue []*TreeNode
	queue = append(queue, root)
	status := true
	for len(queue) > 0 {
		var curLevel []int
		size := len(queue)
		for size > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			curLevel = append(curLevel, node.Val)
			size--
		}
		if !status {
			temp := make([]int, len(curLevel))
			for i, v := range curLevel {
				temp[len(curLevel)-i-1] = v
			}
			curLevel = temp
		}
		res = append(res, curLevel)
		status = !status
	}
	return res
}
