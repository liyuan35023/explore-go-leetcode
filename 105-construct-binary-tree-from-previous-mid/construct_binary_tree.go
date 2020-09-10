package _05_construct_binary_tree_from_previous_mid
/*
根据一棵树的前序遍历与中序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：

    3
   / \
  9  20
    / \
   15 7

https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal

 */
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := new(TreeNode)
	cur := root
	cur.Val = preorder[0]
	indexInOrder := findIndexInInorder(inorder, cur.Val)
	left := buildTree(preorder[1:indexInOrder+1], inorder[0:indexInOrder])
	right := buildTree(preorder[indexInOrder+1:], inorder[indexInOrder+1:])
	cur.Left = left
	cur.Right = right
	return cur
}

func findIndexInInorder(inorder []int, val int) int {
	for k, v := range inorder {
		if v == val {
			return k
		}
	}
	return 0
}
