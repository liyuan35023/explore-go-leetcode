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

}

