package _1_rotate_list

/*
	Example 1:
	Input: 1->2->3->4->5->NULL, k = 2
	Output: 4->5->1->2->3->NULL
	Explanation:
	rotate 1 steps to the right: 5->1->2->3->4->NULL
	rotate 2 steps to the right: 4->5->1->2->3->NULL

	Example 2:


	Input: 0->1->2->NULL, k = 4
	Output: 2->0->1->NULL
	Explanation:
	rotate 1 steps to the right: 2->0->1->NULL
	rotate 2 steps to the right: 1->2->0->NULL
	rotate 3 steps to the right: 0->1->2->NULL
	rotate 4 steps to the right: 2->0->1->NULL

	题目大意 #
	旋转链表 K 次
 */

type ListNode struct {
    Val int
    Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	n := 1
	var tail *ListNode
	for tail = head; tail.Next != nil; tail = tail.Next {
		n++
	}
	tail.Next = head
	realK := k % n
	// real K 代表倒数第k个节点为新的head
	cur := tail
	for i := 0; i < n - realK; i++ {
		cur = cur.Next
	}
	newHead := cur.Next
	cur.Next = nil

	return newHead
}
