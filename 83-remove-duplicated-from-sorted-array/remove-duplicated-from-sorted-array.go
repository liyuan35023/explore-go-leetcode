package _3_remove_duplicated_from_sorted_array

/*
	Example 1:


	Input: 1->1->2
	Output: 1->2

	Example 2:


	Input: 1->1->2->3->3
	Output: 1->2->3

	题目大意 #
	删除链表中重复的结点，以保障每个结点只出现一次。
 */

type ListNode struct {
    Val int
    Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
    cur := head
    for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

func deleteDuplicatesII(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	for next != nil && head.Val == next.Val {
		next = next.Next
	}
	head.Next = deleteDuplicatesII(next)
	return head
}
