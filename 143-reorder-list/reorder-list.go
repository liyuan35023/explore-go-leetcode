package _43_reorder_list

/*
	Example 1:
	Given 1->2->3->4, reorder it to 1->4->2->3.

	Example 2:
	Given 1->2->3->4->5, reorder it to 1->5->2->4->3.

	题目大意 #
	按照指定规则重新排序链表：第一个元素和最后一个元素排列在一起，接着第二个元素和倒数第二个元素排在一起，接着第三个元素和倒数第三个元素排在一起。

 */

type ListNode struct {
    Val int
    Next *ListNode
}

func reorderList(head *ListNode) {
	// 找到链表的中点
    // 反转后半部分链表
    // 合并两个链表
    if head == nil || head.Next == nil {
        return
    }
    dummyHead := &ListNode{Next: head}
    fast, slow := dummyHead, dummyHead
    for fast != nil && fast.Next != nil {
    	fast = fast.Next.Next
    	slow = slow.Next
    }
    mid := slow
    // 反转
    newTail := mid.Next
    mid.Next = nil // 断链
    newHead := reverse(newTail)
    mergeTwoOneByOne(head, newHead)
}

func reverse(head *ListNode) *ListNode {
    var pre *ListNode
    cur := head
    for cur != nil {
        tmp := cur.Next
        cur.Next = pre
        pre = cur
        cur = tmp
    }
    return pre
}

func mergeTwoOneByOne(l1 *ListNode, l2 *ListNode) *ListNode {
    dummyHead := new(ListNode)
    pre := dummyHead
    for l1 != nil || l2 != nil {
        if l1 != nil {
            pre.Next = l1
            pre = l1
            l1 = l1.Next
        }
        if l2 != nil {
            pre.Next = l2
            pre = l2
            l2 = l2.Next
        }
    }
    return dummyHead.Next
}
