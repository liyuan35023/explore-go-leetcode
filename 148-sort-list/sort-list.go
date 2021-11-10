package _48_sort_list

/*
	Example 1:

	Input: 4->2->1->3
	Output: 1->2->3->4

	Example 2:

	Input: -1->5->3->4->0
	Output: -1->0->3->4->5

	题目大意 #
	链表的排序，要求时间复杂度必须是 	O(n log n)，空间复杂度是 O(1)
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
	total := 0
	cur := head
	for cur != nil {
		cur = cur.Next
		total++
	}
	dummy := new(ListNode)
	dummy.Next = head
	for block := 1; block < total; block *= 2 {
		pre := dummy
		cur := dummy.Next
		for cur != nil {
			var firstHead, secondHead *ListNode
			var firstTail, secondTail *ListNode
			firstHead = cur
			for i := 1; i < block && cur != nil; i++ {
				cur = cur.Next
			}
			if cur == nil || cur.Next == nil {
				pre.Next = firstHead
				break
			}
			firstTail = cur
			secondHead = cur.Next

			cur = cur.Next
			firstTail.Next = nil
			for i := 1; i < block && cur != nil; i++ {
				cur = cur.Next
			}
			secondTail = cur
			if cur != nil {
				cur = cur.Next
				secondTail.Next = nil
			}
			newHead, newTail := merge(firstHead, secondHead)
			pre.Next = newHead
			pre = newTail
		}
	}
	return dummy.Next
}

func merge(head1 *ListNode, head2 *ListNode) (*ListNode, *ListNode) {
	dummy := new(ListNode)
	pre := dummy
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			pre.Next = head1
			head1 = head1.Next
		} else {
			pre.Next = head2
			head2 = head2.Next
		}
		pre = pre.Next
	}
	if head1 != nil {
		pre.Next = head1
		for ; head1 != nil; head1 = head1.Next {
			pre = pre.Next
		}
	}
	if head2 != nil {
		pre.Next = head2
		for ; head2 != nil; head2 = head2.Next {
			pre = pre.Next
		}
	}
	return dummy.Next, pre
}