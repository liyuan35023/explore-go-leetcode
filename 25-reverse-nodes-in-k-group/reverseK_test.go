package _5_reverse_nodes_in_k_group

import (
	"fmt"
	"testing"
)

func TestReverseK(t *testing.T) {
	five := &ListNode{Val: 5}
	four := &ListNode{Val: 4, Next: five}
	three := &ListNode{Val: 3, Next: four}
	two := &ListNode{Val: 2, Next: three}
	one := &ListNode{Val: 1, Next: two}
	head := reverseKGroup(one, 2)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}

}
