package _43_diameter_of_binary_tree

/*
	给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。

	示例 :
	给定二叉树

			  1
			 / \
			2   3
		   / \
		  4   5
	返回3, 它的长度是路径 [4,2,1,3] 或者[5,2,1,3]。

	注意：两结点之间的路径长度是以它们之间边的数目表示。
 */

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0
	var helper func(node *TreeNode) int
    helper = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        left := helper(node.Left)
        right := helper(node.Right)
        ans = max(ans, left+right)
        return max(left, right) + 1
    }
    helper(root)
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
