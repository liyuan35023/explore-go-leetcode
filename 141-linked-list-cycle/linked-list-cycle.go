package _41_linked_list_cycle

/*
题目大意 #
判断链表是否有环，不能使用额外的空间。

 */

type ListNode struct {
    Val int
    Next *ListNode
}

func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }
    // 反转
    var pre *ListNode
    cur := head
    for cur != nil {
        tmp := cur.Next
        cur.Next = pre
        pre = cur
        cur = tmp
    }
    if pre == head {
        return true
    }
    return false
}

func hasCycleII(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }
    fast, slow := head.Next, head
    for fast != nil && fast.Next != nil && fast != slow {
        fast = fast.Next.Next
        slow = slow.Next
    }
    if fast == slow {
        return true
    }
    return false
}
