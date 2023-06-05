package linkedlist

import "testing"

func TestHello(t *testing.T) {
	t.Log("nihao")
}

// mergeTwoLists 将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	// 递归法：如果list1的值小于list2的值，就比较list.Next和list2的结果作为list1.Next，直到有一个为空
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}
