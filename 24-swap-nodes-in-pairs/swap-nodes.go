package _4_swap_nodes_in_pairs

/*
	Given a linked list, swap every two adjacent nodes and return its head.
	You may not modify the values in the list’s nodes, only nodes itself may be changed.

	Example:
	Given 1->2->3->4, you should return the list as 2->1->4->3.

 */

type ListNode struct {
	Val int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{Next: head}
	pre := dummyHead
	cur := head
	for cur != nil && cur.Next != nil {
		tmp := cur.Next.Next
		pre.Next = cur.Next
		cur.Next.Next = cur
		cur.Next = tmp
		pre = cur
		cur = tmp
	}
	return dummyHead.Next
}

func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	tmp := head.Next.Next
	head.Next.Next = head
	head.Next = swapPairs2(tmp)
	return newHead
}