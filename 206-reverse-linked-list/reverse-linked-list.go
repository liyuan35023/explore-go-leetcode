package _06_reverse_linked_list

/*
	反转单链表
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	pre := (*ListNode)(nil)
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

func ReverseRecur(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := ReverseRecur(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func recurHelper(head *ListNode, pre *ListNode) *ListNode {
	if head == nil {
		return pre
	}

	tmp := head.Next
	head.Next = pre
	return recurHelper(tmp, head)
}