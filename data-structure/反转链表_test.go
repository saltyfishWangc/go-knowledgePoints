package data_structure

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		cur.Next, pre, cur = pre, cur, cur.Next
	}
	return cur
}
