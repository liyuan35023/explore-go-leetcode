package _48_sort_list

import (
	"fmt"
	"testing"
)

func TestSortListLoop(t *testing.T) {
	node4 := &ListNode{Val: 3}
	node3 := &ListNode{Val: 1, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	node1 := &ListNode{Val: 4, Next: node2}

	head := SortListLoop(node1)
	fmt.Println(head.Val)
}
