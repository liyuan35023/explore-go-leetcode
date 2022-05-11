package cn
//给定一个单链表 L 的头节点 head ，单链表 L 表示为： 
//
// L0 → L1 → … → Ln-1 → Ln 
//请将其重新排列后变为： 
//
// L0 → Ln → L1 → Ln-1 → L2 → Ln-2 → … 
//
// 不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。 
//
// 示例 1: 
//
//输入: head = [1,2,3,4]
//输出: [1,4,2,3] 
//
// 示例 2: 
//
//输入: head = [1,2,3,4,5]
//输出: [1,5,2,4,3] 
//
// 提示： 
//
// 链表的长度范围为 [1, 5 * 104]
// 1 <= node.val <= 1000 

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {



}
//func reorderList(head *ListNode)  {
//	// find mid
//	slow, fast := head, head
//	for fast != nil && fast.Next != nil {
//		slow = slow.Next
//		fast = fast.Next.Next
//	}
//	var pre *ListNode
//	for slow != nil {
//		tmp := slow.Next
//		slow.Next = pre
//		pre = slow
//		slow = tmp
//	}
//	first := head
//	last := pre
//	for first != last && first.Next != last {
//		tmp1 := first.Next
//		tmp2 := last.Next
//		first.Next = last
//		last.Next = tmp1
//		first = tmp1
//		last = tmp2
//	}
//	return
//}
//leetcode submit region end(Prohibit modification and deletion)
