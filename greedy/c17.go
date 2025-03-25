package greedy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minCameraCover(root *TreeNode) int {

	var res int
	/*
		后序遍历
		0：无覆盖
		1：有摄像头
		2：有覆盖
	*/
	var dfs func(cur *TreeNode) int
	dfs = func(cur *TreeNode) int {
		if cur == nil {
			return 2
		}
		left := dfs(cur.Left)
		right := dfs(cur.Right)
		//
		if left == 2 && right == 2 {
			return 0
		}
		// 在孩子节点存没有被进行覆盖节点，此时父节点需要有摄像头
		if left == 0 || right == 0 {
			res++
			return 1
		}
		// 孩子节点存在一个摄像头，父节点此时就已经有覆盖
		if left == 1 || right == 1 {
			return 2
		}
		// 这里是不会走到这里来的
		return -1
	}
	rootStatus := dfs(root)
	// 在调用dfs(root)后增加根节点的状态检查
	if rootStatus == 0 {
		res++
	}
	return res
}
