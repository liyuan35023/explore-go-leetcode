package __Add_Two_Numbers

/**
 * You are given two non-empty linked lists representing two non-negative integers.
 * The digits are stored in reverse order and each of their nodes contain a single digit.
 * Add the two numbers and return it as a linked list.

 * You may assume the two numbers do not contain any leading zero, except the number 0 itself.

 * Example:

 * Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
 * Output: 7 -> 0 -> 8
 * Explanation: 342 + 465 = 807.
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	cur := dummyHead
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		cur.Next = &ListNode{Val:carry%10}
		carry = carry / 10
		cur = cur.Next
	}
	return dummyHead.Next


	//cur1 := l1
	//cur2 := l2
	//startNode := &ListNode{}
	//curNode := startNode
	//carry := 0
	//v1 := 0
	//v2 := 0
	//
	//for cur1 != nil || cur2 != nil {
	//	if cur1 != nil {
	//		v1 = cur1.Val
	//	} else {
	//		v1 = 0
	//	}
	//
	//	if cur2 != nil {
	//		v2 = cur2.Val
	//	} else {
	//		v2 = 0
	//	}
	//	curNode.Val = (v1 + v2 + carry) % 10
	//	carry = (v1 + v2 + carry) / 10
	//	if cur1 != nil {
	//		cur1 = cur1.Next
	//	}
	//	if cur2 != nil {
	//		cur2 = cur2.Next
	//	}
	//	if cur1 == nil && cur2 == nil {
	//		break
	//	}
	//	nextNode := &ListNode{}
	//	curNode.Next = nextNode
	//	curNode = nextNode
	//}
	//if carry == 1 {
	//	curNode.Next = &ListNode{Val:1}
	//}
	//return startNode

}