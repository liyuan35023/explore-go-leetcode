package cn

import (
	"bytes"
	"strconv"
	"strings"
)

//序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方
//式重构得到原数据。 
//
// 请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串
//反序列化为原始的树结构。 
//
// 提示: 输入输出格式与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的
//方法解决这个问题。 
//
// 
//
// 示例 1： 
//
// 
//输入：root = [1,2,3,null,null,4,5]
//输出：[1,2,3,null,null,4,5]
// 
//
// 示例 2： 
//
// 
//输入：root = []
//输出：[]
// 
//
// 示例 3： 
//
// 
//输入：root = [1]
//输出：[1]
// 
//
// 示例 4： 
//
// 
//输入：root = [1,2]
//输出：[1,2]
// 
// 提示：
//
// 
// 树中结点数在范围 [0, 10⁴] 内 
// -1000 <= Node.val <= 1000 
// 


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	ans := bytes.NewBuffer([]byte{})
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			if queue[i] == nil {
				ans.WriteString("null,")
			} else {
				ans.WriteString(strconv.Itoa(queue[i].Val)+",")
				queue = append(queue, queue[i].Left)
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[n:]
	}
	ret := strings.TrimRight(ans.String(), ",")
	return ret

}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	myAtoi := func(s string) int {
		ret, _ := strconv.Atoi(s)
		return ret
	}
	cells := strings.Split(data, ",")
	queue := make([]*TreeNode, 0)
	root := &TreeNode{Val: myAtoi(cells[0])}
	queue = append(queue, root)
	idx := 1
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if cells[idx] == "null" {
			node.Left = nil
		} else {
			node.Left = &TreeNode{Val: myAtoi(cells[idx])}
			queue = append(queue, node.Left)
		}
		if cells[idx+1] == "null" {
			node.Right = nil
		} else {
			node.Right = &TreeNode{Val: myAtoi(cells[idx+1])}
			queue = append(queue, node.Right)
		}
		idx = idx + 2
	}
	return root





}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
//leetcode submit region end(Prohibit modification and deletion)
