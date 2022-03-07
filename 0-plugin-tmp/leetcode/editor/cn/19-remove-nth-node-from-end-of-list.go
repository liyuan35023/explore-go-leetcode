package cn
//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。 
//
// 进阶：你能尝试使用一趟扫描实现吗？ 
//
//
// 示例 1： 
//
//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]
// 
// 示例 2：
//
//输入：head = [1], n = 1
//输出：[]
//
// 示例 3： 
//
// 
//输入：head = [1,2], n = 1
//输出：[1]
//
//
// 提示： 
//
// 链表中结点的数目为 sz 
// 1 <= sz <= 30 
// 0 <= Node.val <= 100 
// 1 <= n <= sz 
// 


//leetcode submit region begin(Prohibit modification and deletion)

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	slow, fast := dummy, dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
//leetcode submit region end(Prohibit modification and deletion)
