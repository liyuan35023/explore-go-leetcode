package cn
//给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。 
//
// 你可以假设除了数字 0 之外，这两个数字都不会以零开头。 
//
// 示例1： 
//
//输入：l1 = [7,2,4,3], l2 = [5,6,4]
//输出：[7,8,0,7]
// 
// 示例2：
//
//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[8,0,7]
// 
// 示例3：
//
//输入：l1 = [0], l2 = [0]
//输出：[0]
//
// 提示： 
//
// 链表的长度范围为 [1, 100] 
// 0 <= node.val <= 9 
// 输入数据保证链表代表的数字无前导 0 
//
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1 := make([]int, 0)
	s2 := make([]int, 0)
	for l1 != nil || l2 != nil {
		if l1 != nil {
			s1 = append(s1, l1.Val)
			l1 = l1.Next
		}
		if l2 != nil {
			s2 = append(s2, l2.Val)
			l2 = l2.Next
		}
	}
	var next *ListNode
	carry := 0
	for len(s1) != 0 || len(s2) != 0 || carry != 0 {
		var d1, d2 int
		if len(s1) > 0 {
			d1 = s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
		}
		if len(s2) > 0 {
			d2 = s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
		}
		tmp := d1 + d2 + carry
		cur := &ListNode{Val: tmp % 10}
		cur.Next = next
		next = cur
		carry = tmp / 10
	}
	return next
}

//leetcode submit region end(Prohibit modification and deletion)
