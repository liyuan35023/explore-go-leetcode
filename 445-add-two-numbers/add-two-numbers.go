package _45_add_two_numbers

/*
	Example:

	Input: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
	Output: 7 -> 8 -> 0 -> 7

	题目大意 #
	这道题是第 2 题的变种题，第 2 题中的 2 个数是从个位逆序排到高位，这样相加只用从头交到尾，进位一直进位即可。
	这道题目中强制要求不能把链表逆序。2 个数字是从高位排到低位的，这样进位是倒着来的。

	You may assume the two numbers do not contain any leading zero, except the number 0 itself.

	Follow up: What if you cannot modify the input lists? In other words, reversing the lists is not allowed.

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers445(l1 *ListNode, l2 *ListNode) *ListNode {
	// stack
	stack1 := make([]*ListNode, 0)
	stack2 := make([]*ListNode, 0)
	for l1 != nil {
		stack1 = append(stack1, l1)
		l1 = l1.Next
	}
	for l2 != nil {
		stack2 = append(stack2, l2)
		l2 = l2.Next
	}
	var pre *ListNode
	var cur1, cur2 int
	carry := 0
	for len(stack1) != 0 || len(stack2) != 0 || carry != 0 {
		if len(stack1) != 0 {
			cur1 = stack1[len(stack1)-1].Val
			stack1 = stack1[:len(stack1)-1]
		} else {
			cur1 = 0
		}
		if len(stack2) != 0 {
			cur2 = stack2[len(stack2)-1].Val
			stack2 = stack2[:len(stack2)-1]
		} else {
			cur2 = 0
		}
		node := &ListNode{Val: (cur1 + cur2 + carry) % 10, Next: pre}
		carry = (cur1 + cur2 + carry) / 10
		pre = node
	}
	return pre
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1 := make([]int, 0)
	stack2 := make([]int, 0)
	for l1 != nil || l2 != nil {
		if l1 != nil {
			stack1 = append(stack1, l1.Val)
			l1 = l1.Next
		}
		if l2 != nil {
			stack2 = append(stack2, l2.Val)
			l2 = l2.Next
		}
	}
	m, n := len(stack1), len(stack2)
	var next *ListNode
	carry := 0
	for i, j := m-1, n-1; i >= 0 || j >= 0 || carry > 0; i, j = i-1, j-1 {
		v1, v2 := 0, 0
		if i >= 0 {
			v1 = stack1[i]
		}
		if j >= 0 {
			v2 = stack2[j]
		}
		cur := &ListNode{Val: (v1 + v2 + carry) % 10, Next: next}
		carry = (v1 + v2 + carry) / 10
		next = cur
	}
	return next
}

func addTwoNumbersIII(l1 *ListNode, l2 *ListNode) *ListNode {
	m, n := 0, 0
	cur1, cur2 := l1, l2
	for cur1 != nil || cur2 != nil {
		if cur1 != nil {
			m++
			cur1 = cur1.Next
		}
		if cur2 != nil {
			n++
			cur2 = cur2.Next
		}
	}
	var next *ListNode
	cur1, cur2 = l1, l2
	if m > n {
		dis := m - n
		for dis > 0 {
			tmp := &ListNode{Val: cur1.Val, Next: next}
			next = tmp
			cur1 = cur1.Next
			dis--
		}
	} else {
		dis := n - m
		for dis > 0 {
			tmp := &ListNode{Val: cur2.Val, Next: next}
			next = tmp
			cur2 = cur2.Next
			dis--
		}
	}
	for cur1 != nil && cur2 != nil {
		tmp := &ListNode{Val: cur1.Val + cur2.Val, Next: next}
		next = tmp
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	return reverseAndCal(next)
}

func reverseAndCal(node *ListNode) *ListNode {
	var pre *ListNode
	carry := 0
	for node != nil || carry != 0 {
		if node == nil {
			node = &ListNode{Val: carry % 10}
			carry = carry / 10
		} else {
			node.Val, carry = (node.Val + carry) % 10, (node.Val + carry) / 10
		}
		tmp := node.Next
		node.Next = pre
		pre = node
		node = tmp
	}
	return pre
}
