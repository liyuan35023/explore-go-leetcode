package cn

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
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil
	var pre *ListNode
	for mid != nil {
		tmp := mid.Next
		mid.Next = pre
		pre = mid
		mid = tmp
	}
	halfHead := pre

	for head != nil && halfHead != nil {
		if head.Val != halfHead.Val {
			return false
		}
		head = head.Next
		halfHead = halfHead.Next
	}
	return true
}
//leetcode submit region end(Prohibit modification and deletion)
