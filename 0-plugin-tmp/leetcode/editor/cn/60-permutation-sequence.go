package cn

import "strconv"

//给出集合 [1,2,3,...,n]，其所有元素共有 n! 种排列。
//
// 按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下： 
//
// "123" 
// "132" 
// "213" 
// "231" 
// "312" 
// "321" 
//
// 给定 n 和 k，返回第 k 个排列。 
//
// 示例 1： 
//
//输入：n = 3, k = 3
//输出："213"
//
// 示例 2： 
//
//输入：n = 4, k = 9
//输出："2314"
// 
// 示例 3：
//
//输入：n = 3, k = 1
//输出："123"
//
// 提示： 
//
// 1 <= n <= 9 
// 1 <= k <= n! 

//leetcode submit region begin(Prohibit modification and deletion)
func getPermutation(n int, k int) string {
	seq := make([]int, n)
	for i := 0; i < n; i++ {
		seq[i] = i + 1
	}

	ans := ""
	for n > 1 {
		perNum := jiecheng(n-1)
		curSeqIdx := (k - 1) / perNum
		ans += strconv.Itoa(seq[curSeqIdx])
		seq = append(seq[:curSeqIdx], seq[curSeqIdx+1:]...)
		k = k - curSeqIdx * perNum
		n--
	}
	ans += strconv.Itoa(seq[0])
	return ans
}

func jiecheng(n int) int {
	ans := 1
	for i := 1; i <= n; i++ {
		ans *= i
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
