package _34_palindrome_linked_list

/*
	Example 1:


	Input: 1->2
	Output: false

	Example 2:


	Input: 1->2->2->1
	Output: true

	Follow up:

	Could you do it in O(n) time and O(1) space?

	题目大意 #
	判断一个链表是否是回文链表。要求时间复杂度 O(n)，空间复杂度 O(1)。

 */

type ListNode struct {
    Val int
    Next *ListNode
}

func isPalindrome234(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return true
    }
    // total number
    num := 0
    cur := head
    for cur != nil {
        num++
        cur = cur.Next
    }
    // find mid
    // reverse half linked list
    cur = head
    var pre *ListNode
    count := 0
    var head1, head2 *ListNode
    for ; count < num / 2; count++ {
        tmp := cur.Next
        cur.Next = pre
        pre = cur
        cur = tmp
    }
    if num % 2 == 0 {
        head1 = pre
        head2 = cur
    } else {
        head1 = pre
        head2 = cur.Next
    }
    for head1 != nil {
        if head1.Val != head2.Val {
            return false
        }
        head1 = head1.Next
        head2 = head2.Next
    }
    return true
}
