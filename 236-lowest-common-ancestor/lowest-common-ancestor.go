package _36_lowest_common_ancestor

/*
	给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

	百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

	例如，给定如下二叉树: root =[3,5,1,6,2,0,8,null,null,7,4]

	示例 1:

	输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
	输出: 3
	解释: 节点 5 和节点 1 的最近公共祖先是节点 3。

	示例2:

	输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
	输出: 5
	解释: 节点 5 和节点 4 的最近公共祖先是节点 5。因为根据定义最近公共祖先节点可以为节点本身。

	说明:
	所有节点的值都是唯一的。
	p、q 为不同节点且均存在于给定的二叉树中。

 */

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
var ans *TreeNode

func lowestCommonAncestorRec(root, p, q *TreeNode) *TreeNode {
	// 递归
	if root == nil {
		return nil
	}
	lNode := lowestCommonAncestor(root.Left, p, q)
	rNode := lowestCommonAncestor(root.Right, p, q)
	if lNode != nil && rNode != nil || root == p || root == q {
		return root
	}
	if lNode != nil {
		return lNode
	}
	return rNode
}

func helper(node, p, q *TreeNode) bool {
	// helper 代表 node 节点的子树或其本身包不包含p 或 q
	if node == nil {
		return false
	}
	l := helper(node.Left, p, q)
	r := helper(node.Right, p, q)

	if l && r || (node.Val == p.Val || node.Val == q.Val) && (l || r) {
		ans = node
		return true
	}
	return l || r || node.Val == p.Val || node.Val == q.Val
}


func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}
	pre := preOrder(root)
	in, pIdx, qIdx := inOrder(root, p, q)
	midNode := pre[0]
	for {
		midIdx := findMidInInorder(midNode, in)
		var ifMid bool
		var newMidNode *TreeNode
		ifMid, newMidNode, pre, in = midBetweenPQ(midIdx, pIdx, qIdx, pre, in)
		if ifMid {
			return midNode
		}
		midNode = newMidNode
	}
}

func findMidInInorder(node *TreeNode, inorder []*TreeNode) int {
	for k, v := range inorder {
		if v == node {
			return k
		}
	}
	return -1
}

func midBetweenPQ(midIdx, pIdx, qIdx int, preOrder []*TreeNode, inOrder []*TreeNode) (bool, *TreeNode, []*TreeNode, []*TreeNode)  {
	if pIdx <= midIdx && qIdx >= midIdx || qIdx <= midIdx && pIdx >= midIdx {
		return true, nil, preOrder, inOrder
	}
	if pIdx < midIdx && qIdx < midIdx {
		preOrder = preOrder[1:midIdx+1]
		inOrder = inOrder[:midIdx]
		return false, preOrder[1], preOrder, inOrder
	} else {
		preOrder = preOrder[midIdx+1:]
		inOrder = inOrder[midIdx+1:]
		return false, preOrder[midIdx+1], preOrder, inOrder
	}
}

func preOrder(root *TreeNode) []*TreeNode {
	ans := make([]*TreeNode, 0)
	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}
		ans = append(ans, node)
		helper(node.Left)
		helper(node.Right)
	}
	helper(root)
    return ans
}

func inOrder(root *TreeNode, p, q *TreeNode) ([]*TreeNode, int, int) {
	ans := make([]*TreeNode, 0)
	var pIdx int
	var qIdx int
	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}
		helper(node.Left)
		if node == p {
			pIdx = len(ans)
		}
		if node == q {
			qIdx = len(ans)
		}
		ans = append(ans, node)
		helper(node.Right)
	}
	helper(root)
	return ans, pIdx, qIdx
}
