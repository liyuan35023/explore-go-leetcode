package cn
//给定整数 n 和 k，找到 1 到 n 中字典序第 k 小的数字。 
//
// 注意：1 ≤ k ≤ n ≤ 109。 
//
// 示例 : 
//
//输入:
//n: 13   k: 2
//
//输出:
//10
//
//解释:
//字典序的排列是 [1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9]，所以第二小的数字是 10。
//

//leetcode submit region begin(Prohibit modification and deletion)
func findKthNumber(n int, k int) int {







}







//func findKthNumber(n int, k int) int {
//	count := 0
//	prefix := 1
//	preNum := 1
//	for count < k {
//		curPrefixCount := getPrefixCount(prefix, n)
//		preNum = prefix
//		if count + curPrefixCount < k {
//			// 从兄弟节点找
//			prefix = prefix + 1
//			count += curPrefixCount
//		} else {
//			prefix = prefix * 10
//			count++
//		}
//	}
//	return preNum
//}
//
//func getPrefixCount(prefix int, n int) int {
//	//// 在十叉树中寻找以prefix开头的数字数目，包括他自己
//	first := prefix
//	second := first + 1
//	count := 0
//	for first <= n {
//		//count += min(second, n+1) - first
//		if second > n {
//			count += n - first + 1
//		} else {
//			count += second - first
//		}
//		first *= 10
//		second *= 10
//	}
//	return count
//}

//func min(x, y int) int {
//	if x < y {
//		return x
//	}
//	return y
//}

//leetcode submit region end(Prohibit modification and deletion)

