package tree

import "testing"

func TestLevelOrder(t *testing.T) {
	root := &TreeNode{Val: 3}
	left := &TreeNode{Val: 9}
	right := &TreeNode{Val: 20}
	rightLeft := &TreeNode{Val: 15}
	rightRight := &TreeNode{Val: 7}

	root.Left = left
	root.Right = right

	right.Left = rightLeft
	right.Right = rightRight

	t.Log(levelOrder(root))
}

// levelOrder 给你一个二叉树的根节点root，返回其节点值的层序遍历。(即逐层地，从左到右访问所有节点)
func levelOrder(root *TreeNode) [][]int {

	var res [][]int
	// 定义节点数组p，用来存储每一层的节点
	var p []*TreeNode
	// 初始化时只有根结点
	p = append(p, root)
	for len(p) > 0 {
		// 用来存储每一层节点的值
		var eachLevel []int

		// 用来存储每一层节点的子节点，最后赋值给p，达到每一层一层遍历的效果
		var tmp []*TreeNode
		for _, node := range p {
			eachLevel = append(eachLevel, node.Val)
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
			//tmp = append(tmp, node.Left, node.Right)
		}
		res = append(res, eachLevel)

		p = tmp
	}
	return res
}
