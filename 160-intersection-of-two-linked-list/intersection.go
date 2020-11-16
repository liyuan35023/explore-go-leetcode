package _60_intersection_of_two_linked_list

/*
    找到 2 个链表的交叉点。

    note:
    1. 如果两个链表没有交点，返回null
    2. 在返回结果后，两个链表仍保持原结构
    3. 没有循环
    4. 程序尽量满足O(n), O(1)

 */

type ListNode struct {
    Val int
    Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
        return nil
    }
    curA := headA
    curB := headB
    m := 1
    n := 1
    for curA.Next != nil || curB.Next != nil {
        if curA.Next != nil {
            m++
            curA = curA.Next
        }
        if curB.Next != nil {
            n++
            curB = curB.Next
        }
    }
    if curA != curB {
        return nil
    }
    preA := &ListNode{Next: headA}
    preB := &ListNode{Next: headB}
    curA = preA
    curB = preB

    if m > n {
        diff := m - n
        for curA != curB {
            if diff <= 0 {
                curB = curB.Next
            }
            curA = curA.Next
            diff--
        }
    } else {
    	diff := n - m
        for curA != curB {
            if diff <= 0 {
                curA = curA.Next
            }
            curB = curB.Next
            diff--
        }
    }
    return curA
}

func getIntersectionNodeII(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
        return nil
    }
    curA := headA
    curB := headB
    var tailA *ListNode
    var tailB *ListNode
    for curA != curB {
        if curA.Next == nil {
            tailA = curA
            curA = headB
        } else {
            curA = curA.Next
        }

        if curB.Next == nil {
            tailB = curB
            curB = headA
        } else {
            curB = curB.Next
        }
        if tailA != nil && tailB != nil {
            if tailA != tailB {
                return nil
            }
        }
    }
    if tailA == tailB {
        return curA
    } else {
        return nil
    }
}