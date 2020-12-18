package _42_linked_cycle_II

/*
	Example 1:
	Input: head = [3,2,0,-4], pos = 1
	Output: tail connects to node index 1
	Explanation: There is a cycle in the linked list, where tail connects to the second node.

	Example 2:

	Input: head = [1,2], pos = 0
	Output: tail connects to node index 0
	Explanation: There is a cycle in the linked list, where tail connects to the first node.

	Example 3:

	Input: head = [1], pos = -1
	Output: no cycle
	Explanation: There is no cycle in the linked list.

    Note: Do not modify the linked list.

	题目大意 #
	判断链表是否有环，不能使用额外的空间。如果有环，输出环的起点指针，如果没有环，则输出空。
 */

type ListNode struct {
    Val int
    Next *ListNode
}


func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
	    return nil
    }
    // 判断有没有环
    slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if slow != fast {
		return nil
	}
	prev := head
	for slow != prev {
		slow = slow.Next
		prev = prev.Next
	}
	return slow
}

