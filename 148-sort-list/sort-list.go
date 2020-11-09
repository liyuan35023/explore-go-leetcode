package _48_sort_list

/*
	Example 1:

	Input: 4->2->1->3
	Output: 1->2->3->4

	Example 2:

	Input: -1->5->3->4->0
	Output: -1->0->3->4->5

	题目大意 #
	链表的排序，要求时间复杂度必须是 O(n log n)，空间复杂度是 O(1)
    要满足
 */

type ListNode struct {
    Val int
    Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	// merge sort
    // k个一组进行排序
	// 递归
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	k := 0
	for cur != nil {
		k++
		cur = cur.Next
	}
	mid := head
	if k % 2 == 0 {
		for i := 0; i < k/2 - 1; i++ {
			mid = mid.Next
		}
	} else {
		for i := 0; i < k/2; i++ {
			mid = mid.Next
		}
	}
	rightHead := mid.Next
	mid.Next = nil
	l := sortList(head)
	r := sortList(rightHead)
	return mergeTwo(l, r)
}

func mergeTwo(l *ListNode, r *ListNode) *ListNode {
	dummyHead := new(ListNode)
	pre := dummyHead
	for l != nil && r != nil {
		if l.Val < r.Val {
			pre.Next = l
			l = l.Next
		} else {
			pre.Next = r
			r = r.Next
		}
		pre = pre.Next
	}
	if l != nil {
		pre.Next = l
	}
	if r != nil {
		pre.Next = r
	}

	return dummyHead.Next
}

func SortListLoop(head *ListNode) *ListNode {
	// 迭代：因为要求空间复杂度为O(1)
	if head == nil || head.Next == nil {
		return head
	}
	n := 0
	cur := head
	for cur != nil {
		n++
		cur = cur.Next
	}
	dummyHead := new(ListNode)
	pre := dummyHead
	dummyHead.Next = head

	firstBlockHead, firstBlockTail := head, head
	secondBlockHead, secondBlockTail := head, head
	for block := 1; block < n; block *= 2 {
		firstBlockHead = dummyHead.Next
		pre = dummyHead
		for firstBlockHead != nil {
			// 比较 两个相邻block，进行merge操作
			// 首先找到两个block的head以及tail
			firstBlockTail = firstBlockHead
			secondBlockHead, secondBlockTail = firstBlockHead, firstBlockHead
			for k := 1; k < block && firstBlockTail != nil; k++ {
				firstBlockTail = firstBlockTail.Next
			}
			if firstBlockTail == nil || firstBlockTail.Next == nil {
				pre.Next = firstBlockHead
				break
			}
			secondBlockHead = firstBlockTail.Next
			secondBlockTail = secondBlockHead
			for k := 1; k < block && secondBlockTail != nil; k++ {
				secondBlockTail = secondBlockTail.Next
			}
			// 断开链表 && 记录下次的firstBlockHead
			firstBlockTail.Next = nil
			tmpNextBlockHead := secondBlockTail
			if secondBlockTail != nil {
				tmpNextBlockHead = secondBlockTail.Next
				secondBlockTail.Next = nil
			}
			newHead, newTail := mergeTwoRetHeadAndTail(firstBlockHead, secondBlockHead)
			firstBlockHead = tmpNextBlockHead
			pre.Next = newHead
			pre = newTail
		}
	}
	return dummyHead.Next
}

func mergeTwoRetHeadAndTail(l *ListNode, r *ListNode) (*ListNode, *ListNode) {
	dummyHead := new(ListNode)
	tail := dummyHead
	pre := dummyHead
	for l != nil && r != nil {
		if l.Val < r.Val {
			pre.Next = l
			l = l.Next
		} else {
			pre.Next = r
			r = r.Next
		}
		pre = pre.Next
	}
	if l != nil {
		pre.Next = l
		for l.Next != nil {
			l = l.Next
		}
		tail = l
	}
	if r != nil {
		pre.Next = r
		for r.Next != nil {
			r = r.Next
		}
		tail = r
	}

	return dummyHead.Next, tail
}



