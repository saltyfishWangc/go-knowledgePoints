package tree

import "testing"

func TestMaxDepth(t *testing.T) {
	root := &TreeNode{Val: 3}
	left := &TreeNode{Val: 9}
	right := &TreeNode{Val: 20}
	rightLeft := &TreeNode{Val: 15}
	rightRight := &TreeNode{Val: 7}

	root.Left = left
	root.Right = right
	right.Left = rightLeft
	right.Right = rightRight

	t.Log(maxDepth(root))
}

// maxDepth 给定一个二叉树，找出其最大深度
// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
