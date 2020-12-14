package _2_remove_duplicated_sorted_listII

/*
	Example 1:


	Input: 1->2->3->3->4->4->5
	Output: 1->2->5

	Example 2:


	Input: 1->1->1->2->3
	Output: 2->3

	题目大意 #
	删除链表中重复的结点，只要是有重复过的结点，全部删除。

 */

type ListNode struct {
    Val int
    Next *ListNode
}

func deleteDuplicates1(head *ListNode) *ListNode {
	dummy := new(ListNode)
	pre := dummy
	cur := head
	for cur != nil {
		if cur.Next == nil || cur.Val != cur.Next.Val {
			pre.Next = cur
			pre = cur
			cur = cur.Next
		} else {
			val := cur.Val
			for cur != nil && cur.Val == val {
				cur = cur.Next
			}
			pre.Next = cur
		}
	}

	return dummy.Next
}