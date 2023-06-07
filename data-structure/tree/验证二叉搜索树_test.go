package tree

import (
	"math"
	"testing"
)

func TestIsValidBST(t *testing.T) {
	//root := &TreeNode{Val: 5}
	//left := &TreeNode{Val: 1}
	//right := &TreeNode{Val: 4}
	//rightLeft := &TreeNode{Val: 3}
	//rightRight := &TreeNode{Val: 6}

	//root := &TreeNode{Val: 5}
	//left := &TreeNode{Val: 1}
	//right := &TreeNode{Val: 7}
	//rightLeft := &TreeNode{Val: 6}
	//rightRight := &TreeNode{Val: 8}
	//
	//root.Left = left
	//root.Right = right
	//right.Left = rightLeft
	//right.Right = rightRight

	// [1,null,1]
	//root := &TreeNode{Val: 1}
	//right := &TreeNode{Val: 1}
	//
	//root.Right = right

	// [5,4,6,null,null,3,7]
	root := &TreeNode{Val: 5}
	left := &TreeNode{Val: 4}
	right := &TreeNode{Val: 6}
	rightLeft := &TreeNode{Val: 3}
	rightRight := &TreeNode{Val: 7}

	root.Left = left
	root.Right = right
	right.Left = rightLeft
	right.Right = rightRight
	t.Log(isValidBST(root))
}

// isValidBST 验证二叉搜索树
// 给定一个二叉树的根节点root，判断其是否是一个有效的二叉搜索树
// 有效二叉搜索树定义如下：
//
//	节点的左子树只包含小于当前节点的数
//	节点的右子树只包含大于当前节点的数
//	所有左子树和右子树自身必须也是二叉搜索树
//func isValidBST(root *TreeNode) bool {
//	if root == nil {
//		return true
//	}
//	// 这种判断不对，这种只能判断当前节点的左节点是否小于当前节点，右节点是否大于当前节点，但是并不一定满足左节点一定是小于当前节点所属的根节点
//	// [5,4,6,null,null,3,7] 像这种就判断不了，会出错。因为3是小于6，但是3并小于5
//	if (root.Left != nil && root.Left.Val >= root.Val) || (root.Right != nil && root.Right.Val <= root.Val) {
//		return false
//	}
//	return isValidBST(root.Left) && isValidBST(root.Right)
//}

func isValidBST(root *TreeNode) bool {
	return isValidBSTHelper(root, math.MinInt64, math.MaxInt64)
}

func isValidBSTHelper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}

	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return isValidBSTHelper(root.Left, lower, root.Val) && isValidBSTHelper(root.Right, root.Val, upper)
}
