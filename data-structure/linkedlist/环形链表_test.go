package linkedlist

import "testing"

func TestHasCycle(t *testing.T) {
	// 1 -> 2 -> 3 -> 2
	second := &ListNode{Val: 2}
	//third := &ListNode{Val: 3, Next: second}
	third := &ListNode{Val: 3}
	second.Next = third
	head := &ListNode{Val: 1, Next: second}
	t.Log(hasCycle(head))
}

// hasCycle 判断是否是环形链表
// 给你一个链表的头节点head，判断链表中是否有环
// 如果链表中有某个节点，可以通过连续跟踪next指针再次到达，则链表中存在环。为了表示给定链表中的环，评测系统内部使用整数pos来表示链表尾连接到链表中的位置(索引从0开始)
// 如果链表中存在环，则返回true，否则，返回false
func hasCycle(head *ListNode) bool {
	// 定义快慢指针，两个指针都从头节点开始出发，快指针每次走两步，慢指针每次走一步，如果有环，两个指针一定会相遇。如果快指针走到尾部了，说明没有环
	fast := head
	slow := head

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

func createLinkedListWithN(n int) *ListNode {
	head := &ListNode{}
	tmp := head
	for i := 0; i < n; i++ {
		tmp.Next = &ListNode{
			Val: i,
		}
		tmp = tmp.Next
	}
	return head.Next
}
