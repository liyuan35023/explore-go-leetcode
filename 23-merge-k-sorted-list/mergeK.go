package _3_merge_k_sorted_list

/*
	Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

	Example:


	Input:
	[
  		1->4->5,
  		1->3->4,
  		2->6
	]
	Output: 1->1->2->3->4->4->5->6
 */
type ListNode struct {
    Val int
    Next *ListNode
 }

func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) < 1 {
        return nil
    }
    ans := lists[0]
    for i := 1; i < len(lists); i++ {
        ans = mergeTwo(ans, lists[i])
    }
    return ans
}

func mergeTwo(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := new(ListNode)
    cur := dummy
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            cur.Next = l1
            l1 = l1.Next
        } else {
            cur.Next = l2
            l2 = l2.Next
        }
        cur = cur.Next
    }
    if l1 != nil {
        cur.Next = l1
    }
    if l2 != nil {
        cur.Next = l2
    }
    return dummy.Next
}