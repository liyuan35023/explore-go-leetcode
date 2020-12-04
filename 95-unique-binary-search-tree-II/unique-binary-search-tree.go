package _5_unique_binary_search_tree_II

/*
	Example:
	Input: 3
	Output:
	[
	  [1,null,3,2],
	  [3,2,null,1],
	  [3,1,null,null,2],
	  [2,1,3],
	  [1,null,2,null,3]
	]
	Explanation:
	The above output corresponds to the 5 unique BST's shown below:

	   1         3     3      2      1
		\       /     /      / \      \
		 3     2     1      1   3      2
		/     /       \                 \
	   2     1         2                 3
	题目大意 #
	给定一个整数 n，生成所有由 1 … n 为节点所组成的二叉搜索树。

 */

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n < 1 {
		return nil
	}

	return helper(1, n)
}

func helper(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	if start == end {
		return []*TreeNode{&TreeNode{Val: start}}
	}
	allTrees := make([]*TreeNode, 0)
	for i := start; i <= end; i++ {
		leftTrees := helper(start, i-1)
		rightTrees := helper(i+1, end)

		for m := 0; m < len(leftTrees); m++ {
			for n := 0; n < len(rightTrees); n++ {
				root := &TreeNode{Val: i}
				root.Left = leftTrees[m]
				root.Right = rightTrees[n]
				//newRoot := copyTree(root)
				allTrees = append(allTrees, root)
			}
		}
	}
	return allTrees
}

func copyTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	return &TreeNode{Val: root.Val, Left: copyTree(root.Left), Right: copyTree(root.Right)}
}




