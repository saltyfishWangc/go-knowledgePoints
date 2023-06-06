package tree

import "testing"

func TestSortedArrayToBST(t *testing.T) {
	nums := []int{-10, -3, 0, 5, 9}
	sortedArrayToBST(nums)
}

// sortedArrayToBST 将有序数组转换为二叉搜索树
// 给定一个整数数组，其中元素已经按升序排列，将其转换为一棵 高度平衡 二叉搜索树
// 高度平衡二叉树是一棵满足【每个节点的左右两个子树的高度差的绝对值不超过1】的二叉树
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}

	return sortedArrayToBSTIt(nums, 0, len(nums)-1)
}

// sortedArrayToBSTIt 因为数组是有序的，按照二叉树的特点来，左子树一定是小于根结点的，右子树一定是大于根结点的
// 所以有序数组中的中间结点就是树的根结点，然后用中间结点的左边作为树的左子树，右边作为树的右子树
func sortedArrayToBSTIt(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := (end-start)/2 + start
	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBSTIt(nums, start, mid-1)
	root.Right = sortedArrayToBSTIt(nums, mid+1, end)
	return root
}
