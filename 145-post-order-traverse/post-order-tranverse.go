package _45_post_order_traverse

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
func postorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var pre *TreeNode
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		if root.Right == nil || root.Right == pre {
			ans = append(ans, root.Val)
			stack = stack[:len(stack)-1]
			pre = root
			root = nil
		} else {
			root = root.Right
		}
	}
	return ans
}
