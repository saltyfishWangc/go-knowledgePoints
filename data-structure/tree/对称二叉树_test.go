package tree

import "testing"

func TestIsSymmetric(t *testing.T) {
	root := &TreeNode{Val: 1}
	left := &TreeNode{Val: 2}
	right := &TreeNode{Val: 2}
	leftLeft := &TreeNode{Val: 3}
	leftRight := &TreeNode{Val: 4}
	rightLeft := &TreeNode{Val: 4}
	rightRight := &TreeNode{Val: 3}

	root.Left = left
	root.Right = right
	left.Left = leftLeft
	left.Right = leftRight
	right.Left = rightLeft
	right.Right = rightRight
	t.Log(isSymmetric(root))
}

// isSymmetric 判断是否是对称二叉树
// 给定一个二叉树的根节点root，检查它是否轴对称
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricHelper(root.Left, root.Right)
}

func isSymmetricHelper(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil || left.Val != right.Val {
		return false
	}
	// 左子树的左节点和右子树的右节点比较
	return isSymmetricHelper(left.Left, right.Right) && isSymmetricHelper(left.Right, right.Left)
}
