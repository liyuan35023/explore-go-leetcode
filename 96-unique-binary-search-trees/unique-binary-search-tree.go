package _6_unique_binary_search_trees

/*
	Example:

	Input: 3
	Output: 5
	Explanation:
	Given n = 3, there are a total of 5 unique BST's:

	   1         3     3      2      1
		\       /     /      / \      \
		 3     2     1      1   3      2
		/     /       \                 \
	   2     1         2                 3
	题目大意 #
	给定一个整数 n，求以 1 … n 为节点组成的二叉搜索树有多少种？

 */

func numTrees(n int) int {
	num := make([]int, n+1)
	num[0] = 1
	//num[1] = 1
	for i := 1; i < n+1; i++ {

		for j := 1; j <= i; j++ {
			num[i] += num[j-1] * num[i-j]
		}

	}
	return num[n]
}
