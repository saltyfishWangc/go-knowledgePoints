package practise

import (
	"fmt"
	"testing"
)

// TestReverseSingleLinkedList
// 单链表12345678910反转成10987654321
func TestReverseSingleLinkedList(t *testing.T) {
	head := NewListNode(1)

	next := head
	for i := 2; i < 11; i++ {
		tmp := NewListNode(i)
		next.Next = tmp
		next = next.Next
	}
	t.Log("反转前：", printLinkedList(head))
	t.Log("反转后：", printLinkedList(ReverseLinkedList(head)))
}

func ReverseLinkedList(head *ListNode) *ListNode {
	// 前一个结点
	var pre *ListNode
	cur := head
	// 注意：每一次操作只是把当前节点的Next保存给下一次当做cur，当前节点指向前一个结点，并没有处理当前节点和下一个节点的关系，是放在下一次循环做的
	for cur != nil {
		cur.Next, pre, cur = pre, cur, cur.Next
	}
	return cur
}

// ReverseLinkedListWithFinalN 从链表尾部开始，每n个节点成一组，反转链表，如果最后头部剩余的节点不足n个，不反转
func ReverseLinkedListWithFinalN(head *ListNode, n int) *ListNode {
	// 因为要从尾部开始，需要用前后指针，找到后N个节点，进行反转
	return nil
}

func NewListNode(data interface{}) *ListNode {
	return &ListNode{Data: data}
}

// ListNode 链表结点
type ListNode struct {
	Data interface{}
	Next *ListNode
}

// printLinkedList 打印链表
func printLinkedList(head *ListNode) string {
	res := ""
	for head != nil {
		if res == "" {
			res = fmt.Sprintf("%d", head.Data.(int))
		} else {
			res = fmt.Sprintf("%s -> %d ", res, head.Data.(int))
		}
		head = head.Next
	}
	return res
}
