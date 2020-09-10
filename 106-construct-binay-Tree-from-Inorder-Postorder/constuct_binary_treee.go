package _06_construct_binay_Tree_from_Inorder_Postorder
/*
根据一棵树的中序遍历与后序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

中序遍历 inorder = [9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7

 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}



func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 && len(postorder) == 0 {
		return nil
	}

	root := new(TreeNode)
	root.Val = postorder[len(postorder)-1]
	index := findRootInInOrder(inorder, root.Val)
	left := buildTree(inorder[:index], postorder[:index])
	right := buildTree(inorder[index+1:], postorder[index:len(postorder)-1])
	root.Left = left
	root.Right = right
	return root
}

func findRootInInOrder(inorder []int, val int) int {
	for k, v := range inorder {
		if v == val {
			return k
		}
	}
	return 0
}

