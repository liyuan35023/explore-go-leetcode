package _10_balanced_binary_tree

/*
	Example 1:

	Given the following tree [3,9,20,null,null,15,7]:


		3
	   / \
	  9  20
		/  \
	   15   7

	Return true.

	Example 2:

	Given the following tree [1,2,2,3,3,null,null,4,4]:


		   1
		  / \
		 2   2
		/ \
	   3   3
	  / \
	 4   4

	Return false.

	题目大意 #
	判断一棵树是不是平衡二叉树。平衡二叉树的定义是：树中每个节点都满足左右两个子树的高度差 <= 1 的这个条件。
 */

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }
    var dfs func(node *TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        lDepth := dfs(node.Left)
        rDepth := dfs(node.Right)
        if lDepth != -1 && rDepth != -1 && abs(lDepth, rDepth) <= 1 {
            return max(lDepth, rDepth) + 1
        }
        return -1
    }
    balanced := dfs(root)
    return balanced >= 0
}

func max(x ,y int) int {
    if x > y {
        return x
    }
    return y
}

func abs(x, y int) int {
    if x > y {
        return x - y
    } else {
        return y - x
    }
}

