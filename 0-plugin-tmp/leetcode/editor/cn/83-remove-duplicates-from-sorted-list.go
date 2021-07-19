package cn
//存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除所有重复的元素，使每个元素 只出现一次 。 
//
// 返回同样按升序排列的结果链表。 
//
// 示例 1： 
//
//输入：head = [1,1,2]
//输出：[1,2]
// 
// 示例 2：
//
//输入：head = [1,1,2,3,3]
//输出：[1,2,3]
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
	if head == nil || head.Next == nil {
		return head
	}
	pre := head
	cur := head.Next
	for cur != nil {
		if cur.Val != pre.Val {
			pre.Next = cur
			pre = cur
		}
		cur = cur.Next
	}
	pre.Next = nil
	return head
}
//leetcode submit region end(Prohibit modification and deletion)
