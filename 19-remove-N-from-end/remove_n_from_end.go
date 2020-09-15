package _9_remove_N_from_end

type ListNode struct {
	Val int
    Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	i, j := dummy, dummy
	step := 0
	for j.Next != nil {
		j = j.Next
		step++
		if step > n {
			i = i.Next
		}
	}
	if i.Next != nil {
		i.Next = i.Next.Next
	}
	return dummy.Next
}
