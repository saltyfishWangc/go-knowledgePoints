package linkedlist

import "fmt"

// ListNode 链表结点
type ListNode struct {
	Val  int
	Next *ListNode
}

// createLinkedListWithSlice 根据数组指定的值构建链表，返回头节点
func createLinkedListWithSlice(sl []int) *ListNode {
	head := &ListNode{}
	tmp := head
	for _, val := range sl {
		tmp.Next = &ListNode{Val: val}
		tmp = tmp.Next
	}
	return head.Next
}

// printLinkedList 打印链表
func printLinkedList(head *ListNode) string {
	str := ""
	p := head
	for p != nil {
		if str == "" {
			str = fmt.Sprintf("%d", p.Val)
		} else {
			str = fmt.Sprintf("%s -> %d", str, p.Val)
		}
		p = p.Next
	}
	return str
}

// createLinkedListWithN 构建指定数量节点的链表，返回头节点
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
