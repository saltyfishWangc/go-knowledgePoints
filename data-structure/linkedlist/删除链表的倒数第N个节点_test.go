package linkedlist

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 如何找到倒数第n个节点：让让一个指针走n个节点，然后再一个节点从头开始相同的步长一起走，等第一个指针走到最后一个节点时，第二个指针的节点就是倒数第N个节点
	tmp := &ListNode{
		Next: head,
	}
	firstPtr := tmp
	secPtr := tmp

	// 先让第一个指针走n个节点
	for i := 0; i < n; i++ {
		firstPtr = firstPtr.Next
	}

	// 两个指针一起走
	for firstPtr.Next != nil {
		firstPtr = firstPtr.Next
		secPtr = secPtr.Next
	}

	secPtr.Next = secPtr.Next.Next
	return head.Next
}
