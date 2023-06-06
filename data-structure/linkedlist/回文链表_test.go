package linkedlist

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	// 1 -> 3 -> 5 -> 3 -> 1
	//sl := []int{1, 3, 5, 3, 1}
	//sl := []int{1, 3, 5, 3}
	//sl := []int{1, 3, 1}
	sl := []int{1, 2, 2, 1}
	head := createLinkedListWithSlice(sl)
	t.Log("链表：", printLinkedList(head))

	//t.Logf("链表：%s 是否是回文链表：%t", printLinkedList(head), isPalindrome(head))
	t.Logf("链表：%s 是否是回文链表：%t", printLinkedList(head), isPalindromeWithArray(head))
}

// createLinkedListWithSlice 根据数组指定的值创建链表，返回头节点
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

// isPalindrome 判断链表是否是回文链表
// 采用快慢指针找到中间节点，然后从中间节点开始将后半链表反转，分别遍历前、后半链表，比较每个节点的值
// 时间复杂度为O(N) + O(N) = O(N)，但是没有额外引入别的数据结构，始终在链表上操作，所以空间复杂度为O(1)
func isPalindrome(head *ListNode) bool {
	// 空链表不是回文链表
	if head == nil {
		return false
	}

	// 只有一个节点的链表是回文链表
	if head.Next == nil {
		return true
	}

	// 通过快慢指针找到中间节点。因为不是环形，所以当快指针跑到最后一个节点时，慢指针所在的就是中间节点
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// 从慢指针所在的节点也就是中间节点开始，把链表的后半部进行反转，这样慢指针走到尾部节点时就是后半链表的头节点
	var pre *ListNode
	cur := slow
	for cur != nil {
		// 将当前指针的下一个指针临时存起来，因为下面要将当前指针的下一个指针指向前一个节点的，避免断层
		tmp := cur.Next
		// 把当前指针
		cur.Next = pre
		// 前一个节点就变成了当前节点
		pre = cur
		// 当前节点就是上面的临时节点，这样就可以继续遍历链表
		cur = tmp
	}
	// 反转完后，pre就是反转后的头节点

	// 开始遍历后半链表和原来的链表，直到后半链表遍历完，中间比较值，如果存在不相等，说明不是回文链表
	// 这里不能只判断pre是否为空。因为我们在处理slow节点开始反转后半链表时，前半链表是断开了的，所以如果链表的长度是复数且链表真的是回文链表的话，那么后半链表就会比前半链表多一个节点，那这里要是只判断pre是否为空的话，head会报空指针
	for pre != nil && head != nil {
		if head.Val != pre.Val {
			return false
		}
		pre = pre.Next
		head = head.Next
	}
	return true
}

// 用数组的方式判断是否是回文链表，时间复杂度为n+n也就是O(1)，但是额外引入了数组，空间复杂度为O(n)
// 还有说法是引入栈结构(数组实现)，先遍历链表，把节点入栈。然后同时遍历栈顶和链表，比较。其实也是和数组是一样的道理
func isPalindromeWithArray(head *ListNode) bool {
	// 空链表不是回文链表
	if head == nil {
		return false
	}

	// 只有一个节点的链表是回文链表
	if head.Next == nil {
		return true
	}

	// 定义整数数组，用来装每个节点的Val
	var sl []int

	for head != nil {
		sl = append(sl, head.Val)
		head = head.Next
	}

	// 遍历数组，判断前后元素是否相等
	for i := 0; i < len(sl)/2; i++ {
		if sl[i] != sl[len(sl)-1-i] {
			return false
		}
	}
	return true
}
