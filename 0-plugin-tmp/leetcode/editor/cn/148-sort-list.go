package cn


//给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
//
// 进阶： 
// 你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？
//
// 示例 1： 
//
//输入：head = [4,2,1,3]
//输出：[1,2,3,4]
//
// 示例 2： 
//
//输入：head = [-1,5,3,4,0]
//输出：[-1,0,3,4,5]
// 
// 示例 3：
//
//输入：head = []
//输出：[]
//
// 提示： 
//
// 链表中节点的数目在范围 [0, 5 * 104] 内 
// -105 <= Node.val <= 105 
//

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func sortList(head *ListNode) *ListNode {
}


//func sortList(head *ListNode) *ListNode {
//	total := 0
//	for cur := head; cur != nil; cur = cur.Next {
//		total++
//	}
//	dummy := &ListNode{Next: head}
//	for block := 1; block < total; block *= 2 {
//		pre := dummy
//		cur := dummy.Next
//		for cur != nil {
//			var head1, head2 *ListNode
//			head1 = cur
//			for i := 1; i < block && cur != nil; i++ {
//				cur = cur.Next
//			}
//			if cur == nil {
//				pre.Next = head1
//				break
//			}
//			head2 = cur.Next
//			cur.Next = nil
//			cur = head2
//			for i := 1; i < block && cur != nil; i++ {
//				cur = cur.Next
//			}
//			if cur != nil {
//				tmp := cur.Next
//				cur.Next = nil
//				cur = tmp
//			}
//			newHead, newTail := sortTwo(head1, head2)
//			pre.Next = newHead
//			pre = newTail
//		}
//	}
//	return dummy.Next
//}
//
//func sortTwo(head1, head2 *ListNode) (*ListNode, *ListNode) {
//	dummy := new(ListNode)
//	pre := dummy
//	for head1 != nil && head2 != nil {
//		if head1.Val < head2.Val {
//			pre.Next = head1
//			pre = head1
//			head1 = head1.Next
//		} else {
//			pre.Next = head2
//			pre = head2
//			head2 = head2.Next
//		}
//	}
//	for head1 != nil {
//		pre.Next = head1
//		pre = head1
//		head1 = head1.Next
//	}
//	for head2 != nil {
//		pre.Next = head2
//		pre = head2
//		head2 = head2.Next
//	}
//	return dummy.Next, pre
//}



//func sortList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	total := 0
//	for cur := head; cur != nil; cur = cur.Next {
//		total++
//	}
//	dummy := &ListNode{Next: head}
//	// block 代表一次排序的一个链表的节点的数目
//	// 最大不能超过 total
//	for block := 1; block < total; block *= 2 {
//		// 一次寻找两个链表的 head节点
//		pre := dummy
//		cur := dummy.Next
//		for cur != nil {
//			head1 := cur
//			// 寻找block个节点
//			for count := 1; count < block && cur != nil; count++ {
//				cur = cur.Next
//			}
//			// 第一个链表就数目不足，那么不需要继续排序，break即可
//			if cur == nil {
//				pre.Next = head1
//				break
//			}
//			head2 := cur.Next
//			cur.Next = nil  // 断链
//			cur = head2
//
//			// 从head2开始找第二个链表的block个节点
//			for count := 1; count < block && cur != nil; count++ {
//				cur = cur.Next
//			}
//			// 第二个链表节点数目不足block个, 直接排序即可
//			if cur != nil {
//				tmp := cur.Next
//				cur.Next = nil
//				cur = tmp
//			}
//			newHead, newTail := sortTwo(head1, head2)
//			pre.Next = newHead
//			pre = newTail
//		}
//	}
//	return dummy.Next
//}
//
//func sortTwo(head1, head2 *ListNode) (*ListNode, *ListNode) {
//	dummy := new(ListNode)
//	pre := dummy
//	for head1 != nil && head2 != nil {
//		if head1.Val < head2.Val {
//			pre.Next = head1
//			pre = head1
//			head1 = head1.Next
//		} else {
//			pre.Next = head2
//			pre = head2
//			head2 = head2.Next
//		}
//	}
//	for head1 != nil {
//		pre.Next = head1
//		pre = head1
//		head1 = head1.Next
//	}
//	for head2 != nil {
//		pre.Next = head2
//		pre = head2
//		head2 = head2.Next
//	}
//	return dummy.Next, pre
//}

//leetcode submit region end(Prohibit modification and deletion)
