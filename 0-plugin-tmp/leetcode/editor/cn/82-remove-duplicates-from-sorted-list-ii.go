package cn
//存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除链表中所有存在数字重复情况的节点，只保留原始链表中 没有重复出现 的数字。 
//
// 返回同样按升序排列的结果链表。 
//
// 示例 1： 
//
//输入：head = [1,2,3,3,4,4,5]
//输出：[1,2,5]
//
// 示例 2： 
//
//输入：head = [1,1,1,2,3]
//输出：[2,3]
// 
// 提示：
//
// 链表中节点数目在范围 [0, 300] 内 
// -100 <= Node.val <= 100 
// 题目数据保证链表已经按升序排列 
// 

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := new(ListNode)
	pre := dummy
	cur := head
	for cur != nil {
		if cur.Next != nil && cur.Val == cur.Next.Val {
			tmp := cur.Next
			for ; tmp != nil && tmp.Val == cur.Val; tmp = tmp.Next {
			}
			cur = tmp
		} else {
			pre.Next = cur
			pre = cur
			cur = cur.Next
		}
	}
	pre.Next = nil
	return dummy.Next


}
//func deleteDuplicates(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	dummy := &ListNode{Val: 101}
//	pre := dummy
//	cur := head
//	for cur != nil && cur.Next != nil {
//		if cur.Val != pre.Val && cur.Val != cur.Next.Val {
//			pre.Next = cur
//			pre = cur
//			cur = cur.Next
//		} else {
//			// find first number not equal to pre
//			for cur.Next != nil && cur.Val == cur.Next.Val {
//				cur = cur.Next
//			}
//			cur = cur.Next
//		}
//	}
//	pre.Next = cur
//	return dummy.Next
//}
//leetcode submit region end(Prohibit modification and deletion)
