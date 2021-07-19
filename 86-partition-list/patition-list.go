package _6_partition_list

type ListNode struct {
	Val int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	dummy := new(ListNode)
	lessTail := dummy
	var biggerHead, biggerTail *ListNode
	cur := head
	for cur != nil {
		if cur.Val >= x {
			if biggerHead == nil {
				biggerHead = cur
			}
			if biggerTail == nil {
				biggerTail = cur
			} else {
				biggerTail.Next = cur
				biggerTail = cur
			}
			cur = cur.Next
			biggerTail.Next = nil
		} else {
			tmp := cur.Next
			lessTail.Next = cur
			lessTail = cur
			cur.Next = nil
			cur = tmp
		}
	}
	lessTail.Next = biggerHead
	return dummy.Next

}