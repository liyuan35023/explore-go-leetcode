package _2_reverse_linkedList_II

/*
	Note: 1 ≤ m ≤ n ≤ length of list.

	Example:

	Input: 1->2->3->4->5->NULL, m = 2, n = 4
	Output: 1->4->3->2->5->NULL

    206 : 反转单链表
	题目大意 #
	给定 2 个链表中结点的位置 m, n，反转这个两个位置区间内的所有结点。 请使用一趟扫描实现
 */

type ListNode struct {
    Val int
    Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
    // 没有用一趟扫描
    dummyHead := &ListNode{Next: head}
    fast := dummyHead
    slowBeforeM := dummyHead
    slowN := dummyHead

    k := 1
    for k <= n {
        if k == m {
            slowBeforeM = fast
        }
        fast = fast.Next
        k++
    }
    slowN = fast
    M := slowBeforeM.Next
    NextN := slowN.Next
    slowBeforeM.Next = nil
    slowN.Next = nil
    reverseK(M)
    slowBeforeM.Next = slowN
    M.Next = NextN
    return dummyHead.Next
}

func reverseK(head *ListNode) *ListNode {
    pre := (*ListNode)(nil)
    cur := head
    for cur != nil {
        tmp := cur.Next
        cur.Next = pre
        pre = cur
        cur = tmp
    }
    return pre
}


func ReverseBetweenOneTurn(head *ListNode, m int, n int) *ListNode {
    // 用一趟扫描完成反转
	if head == nil || head.Next == nil {
	    return head
    }
    dummyHead := &ListNode{Next: head}
    beforeM, afterN := dummyHead, dummyHead
    M, N := dummyHead, dummyHead
    cur := head
    k := 0
    pre := dummyHead
    for {
        k++
        if k + 1 == m {
            beforeM = cur
        }
        if k == m {
            M = cur
        }
        if k == n {
            N = cur
        }
        if k == n + 1 {
            afterN = cur
            beforeM.Next = N
            M.Next = afterN
            break
        }
        if k > m && k <= n {
        	tmp := cur.Next
        	cur.Next = pre
            pre = cur
            cur = tmp
        } else {
            pre = cur
            cur = cur.Next
        }
    }
    return dummyHead.Next
}
