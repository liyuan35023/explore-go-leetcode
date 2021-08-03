package _83_search_tree_mindis

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	// 中序遍历搞一遍
	// 循环中序
	ans := 1 << 31 - 1
	pre := 0
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = min(ans, root.Val - pre)
		pre = root.Val
		root = root.Right
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
