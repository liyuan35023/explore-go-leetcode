package _5_reverse_nodes_in_k_group

/*
	Example:

	Given this linked list: 1->2->3->4->5

	For k = 2, you should return: 2->1->4->3->5

	For k = 3, you should return: 3->2->1->4->5

	Note:
	Only constant extra memory is allowed.
	You may not alter the values in the list’s nodes, only nodes itself may be changed.
	题目大意 #
	按照每 K 个元素翻转的方式翻转链表。如果不满足 K 个元素的就不翻转。
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k < 2 {
		return head
	}
	dummyHead := &ListNode{Next: head}
	cur := head
	pre := dummyHead
	for {
		groupTail := pre
		for i := 0; i < k ; i++ {
			groupTail = groupTail.Next
			if groupTail == nil {
				return dummyHead.Next
			}
		}
		newHead, newTail := reverseK(cur, groupTail)
		pre.Next = newHead
		cur = newTail.Next
		pre = newTail
	}
}

func reverseK(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
	pre := tail.Next
	cur := head
	for pre != tail {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return tail, head
}