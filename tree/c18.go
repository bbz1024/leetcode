package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 是否在p的路径上找到q
	var find func(p *TreeNode, v int) bool
	find = func(p *TreeNode, v int) bool {
		if p == nil {
			return false
		}
		if p.Val == v {
			return true
		}
		return find(p.Left, v) || find(p.Right, v)
	}
	if find(q, p.Val) {
		return q
	}
	if find(p, q.Val) {
		return p
	}
	// 这里可以确保pq一定是有公共的父节点的。
	// 在root先找到p的节点
	// 后序遍历
	var endDfs func(root *TreeNode) *TreeNode
	endDfs = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		res := endDfs(root.Left)
		if res != nil {
			return res
		}
		res = endDfs(root.Right)
		if res != nil {
			return res
		}
		// 在这里进行查找是否存在公共节点
		if find(root, p.Val) && find(root, q.Val) || find(root, q.Val) && find(root, p.Val) {
			return root
		}
		return nil
	}
	return endDfs(root)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor2(root.Left, p, q)
	right := lowestCommonAncestor2(root.Right, p, q)
	// 遇到了left和遇到了right说明就是公共节点
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}
