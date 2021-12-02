package cn
//给你一个链表数组，每个链表都已经按升序排列。 
//
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。 
//
// 示例 1： 
//
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
//输出：[1,1,2,3,4,4,5,6]
//解释：链表数组如下：
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//将它们合并到一个有序链表中得到。
//1->1->2->3->4->4->5->6
//
// 示例 2： 
//
// 输入：lists = []
//输出：[]
//
// 示例 3： 
//
// 输入：lists = [[]]
//输出：[]
//
// 提示： 
// k == lists.length
// 0 <= k <= 10^4 
// 0 <= lists[i].length <= 500 
// -10^4 <= lists[i][j] <= 10^4 
// lists[i] 按 升序 排列 
// lists[i].length 的总和不超过 10^4 
// 
// 👍 1345 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	total := len(lists)
	for _, n := range lists {
		if n != nil {
			lists = append(lists, n)
		}
	}
	lists = lists[total:]

	constructPriorQueue(lists)
	dummy := new(ListNode)
	pre := dummy
	for len(lists) > 0 {
		ans := lists[0]
		lists[0] = lists[0].Next
		if lists[0] == nil {
			lists[0] = lists[len(lists)-1]
			lists = lists[:len(lists)-1]
		}
		downAdjustQueue(lists, 0, len(lists))
		pre.Next = ans
		pre = pre.Next
	}
	return dummy.Next
}

func constructPriorQueue(lists []*ListNode) {
	for i := len(lists) / 2 - 1; i >= 0; i-- {
		downAdjustQueue(lists, i, len(lists))
	}
}

func downAdjustQueue(lists []*ListNode, parent int, length int) {
	for parent <= length / 2 - 1 {
		minChild := getMinChild(lists, parent, length)
		if lists[parent].Val > lists[minChild].Val {
			lists[parent], lists[minChild] = lists[minChild], lists[parent]
			parent = minChild
		} else {
			break
		}
	}
}

func getMinChild(lists []*ListNode, parent int, length int) int {
	child1 := parent * 2 + 1
	child2 := parent * 2 + 2
	if child2 > length - 1 || lists[child1].Val < lists[child2].Val {
		return child1
	} else {
		return child2
	}

}

//func mergeKLists(lists []*ListNode) *ListNode {
//	// 分治
//	if len(lists) < 1 {
//		return nil
//	}
//	for len(lists) != 1 {
//		n := len(lists)
//		for i := 0; i < n; i = i+2 {
//			if i < n - 1 {
//				newHead := mergeTwo(lists[i], lists[i+1])
//				lists = append(lists, newHead)
//			} else {
//				lists = append(lists, lists[i])
//			}
//		}
//		lists = lists[n:]
//	}
//	return lists[0]
//}
//
//func mergeTwo(l1 *ListNode, l2 *ListNode) *ListNode {
//	dummy := new(ListNode)
//	cur := dummy
//	for l1 != nil && l2 != nil {
//		if l1.Val < l2.Val {
//			cur.Next = l1
//			l1 = l1.Next
//		} else {
//			cur.Next = l2
//			l2 = l2.Next
//		}
//		cur = cur.Next
//	}
//	if l1 != nil {
//		cur.Next = l1
//	}
//	if l2 != nil {
//		cur.Next = l2
//	}
//	return dummy.Next
//}

//leetcode submit region end(Prohibit modification and deletion)
