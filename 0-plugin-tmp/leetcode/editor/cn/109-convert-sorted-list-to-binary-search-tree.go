package cn
//给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。 
//
// 本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。 
//
// 示例: 
//
// 给定的有序链表： [-10, -3, 0, 5, 9],
//
//一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：
//
//      0
//     / \
//   -3   9
//   /   /
// -10  5
// 


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {




}
//func sortedListToBST(head *ListNode) *TreeNode {
//	node := head
//	total := 0
//	for head != nil {
//		head = head.Next
//		total++
//	}
//	var helper func(left, right int) *TreeNode
//	helper = func(left, right int) *TreeNode {
//		if left > right {
//			return nil
//		}
//		mid := (left + right + 1) / 2
//		l := helper(left, mid-1)
//		root := &TreeNode{Val: node.Val}
//		node = node.Next
//		r := helper(mid+1, right)
//		root.Left = l
//		root.Right = r
//		return root
//	}
//	return helper(1, total)
//}

//leetcode submit region end(Prohibit modification and deletion)

