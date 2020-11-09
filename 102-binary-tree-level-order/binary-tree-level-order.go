package _02_binary_tree_level_order

/*
	给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。


	示例：
	二叉树：[3,9,20,null,null,15,7],

		3
	   / \
	  9  20
		/  \
	   15   7
	返回其层次遍历结果：

	[
	  [3],
	  [9,20],
	  [15,7]
	]

 */

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
	    return nil
    }
    queue := make([]*TreeNode, 0)
    ans := make([][]int, 0)
    queue = append(queue, root)
    for len(queue) != 0 {
    	n := len(queue)
    	level := make([]int, 0)
    	for i := 0; i < n; i++ {
    		level = append(level, queue[i].Val)
    		if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[n:]
		ans = append(ans, level)
	}
	return ans
}
