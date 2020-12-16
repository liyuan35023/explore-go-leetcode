package _03_binary_tree_zigzag_order

/*
	For Example: Given binary tree [3,9,20,null,null,15,7],


		3
	   / \
	  9  20
		/  \
	   15   7

	return its zigzag level order traversal as:


	[
	  [3],
	  [20,9],
	  [15,7]
	]

	题目大意 #
	按照 Z 字型层序遍历一棵树。
 */


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// 解法一
func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}

	queue := []*TreeNode{root}
	goRight := true
	for len(queue) != 0 {
		n := len(queue)
		level := make([]int, n)

		for i := 0; i < n; i++ {
			if goRight {
				level[i] = queue[i].Val
				//level = append(level, queue[i].Val)
			} else {
				level[n-i-1] = queue[i].Val
				//level = append([]int{queue[i].Val}, level...)
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}

		ans = append(ans, level)
		queue = queue[n:]
		goRight = !goRight
	}
	return ans
}

func zigzagLevelOrderRec(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	zigHelper(root, 0, &ans)
	return ans
}

func zigHelper(root *TreeNode, level int, ans *[][]int) {
	// 偶数层，从左向右
	if root == nil {
		return
	}
	for level + 1 > len(*ans) {
		*ans = append(*ans, []int{})
	}
	if level % 2 == 0 {
		(*ans)[level] = append((*ans)[level], root.Val)
	} else {
		(*ans)[level] = append([]int{root.Val}, (*ans)[level]...)
	}
	zigHelper(root.Left, level+1, ans)
	zigHelper(root.Right, level+1, ans)
}
