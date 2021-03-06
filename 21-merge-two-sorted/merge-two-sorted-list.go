package _1_merge_two_sorted

/*
 	Merge two sorted linked lists and return it as a new list.
	The new list should be made by splicing together the nodes of the first two lists.

	Input: 1->2->4, 1->3->4
	Output: 1->1->2->3->4->4
 */


type ListNode struct {
	Val int
    Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	pre := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			pre.Next = l1
			pre = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			pre = l2
			l2 = l2.Next
		}
	}
	if l1 == nil {
		pre.Next = l2
	} else if l2 == nil {
		pre.Next = l1
	}
	return dummy.Next

	//for l1 != nil || l2 != nil {
	//	if l1 == nil {
	//		pre.Next = l2
	//		break
	//	}
	//	if l2 == nil {
	//		pre.Next = l1
	//		break
	//	}
	//	if l1.Val < l2.Val {
	//		pre.Next = l1
	//		pre = l1
	//		l1 = l1.Next
	//	} else {
	//		pre.Next = l2
	//		pre = l2
	//		l2 = l2.Next
	//	}
	//}
	//return dummy.Next
	for l1 != nil || l2 != nil {
		if l1 == nil {
			pre.Next = l2
			break
		}
		if l2 == nil {
			pre.Next = l1
			break
		}
		if l1.Val < l2.Val {
			pre.Next = l1
			pre = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			pre = l2
			l2 = l2.Next
		}
	}
	return dummy.Next
}

func mergeTwoListsRec(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoListsRec(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoListsRec(l1, l2.Next)
		return l2
	}
}
