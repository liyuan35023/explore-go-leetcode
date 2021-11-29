package cn

import "container/heap"

//请判断一个链表是否为回文链表。
//
// 示例 1: 
//
// 输入: 1->2
//输出: false 
//
// 示例 2: 
//
// 输入: 1->2->2->1
//输出: true
// 
// 进阶：
//你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？ 

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	dummy := new(ListNode)
	dummy.Next = head
	slow, fast := head, dummy
	pre := dummy
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		tmp := slow.Next
		slow.Next = pre
		pre = slow
		slow = tmp
	}
	var start1, start2 *ListNode
	if fast == nil {
		start1, start2 = pre.Next, slow
	} else {
		start1, start2 = pre, slow
	}
	for start1 != nil && start2 != nil {
		if start1.Val != start2.Val {
			return false
		}
		start1 = start1.Next
		start2 = start2.Next
	}
	return true




}
//leetcode submit region end(Prohibit modification and deletion)
