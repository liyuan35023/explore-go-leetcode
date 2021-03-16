package _40_kth_smallest_order

/*
	给定整数n和k，找到1到n中字典序第k小的数字。

	注意：1 ≤ k ≤ n ≤ 10^9。

	示例 :

	输入:
	n: 13   k: 2

	输出:
	10

	解释:
	字典序的排列是 [1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9]，所以第二小的数字是 10。
 */

var totalNum int
func FindKthNumber(n int, k int) int {
	// backTrace 回溯加剪枝 可以ac一半，剩下的超时
	//totalNum = 0
	//return backTrace(n, 0, k)

	// 十叉树
	curCount := 1
	prefix := 1
	for curCount < k {
		count := getPrefixCount(prefix, n)
		if curCount + count > k {
			// 在prefix的子节点中寻找
			curCount++
			prefix *= 10
		} else {
			// 在prefix的兄弟节点中寻找
			curCount += count
			prefix++
		}
	}

	return prefix
}

func getPrefixCount(prefix int, maxNum int) int {
	// 在十叉树中寻找以prefix开头的数字数目，包括他自己

	first := prefix
	second := first + 1
	count := 0
	for first <= maxNum {
		count += min(second, maxNum+1) - first
		first *= 10
		second *= 10
	}
	return count
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func backTrace(n int, cur int, k int) int {
	 if cur >= 1 && cur <= n {
	 	totalNum++
	 	if totalNum == k {
	 		return cur
		}
	 } else if cur > n {
	 	return -2
	 }
	 for i := 0; i <= 9; i++ {
	 	if cur == 0 && i == 0 {
	 		continue
		}
	 	ans := backTrace(n, cur*10+i, k)
		 if ans == -2 {
		 	break
		 } else if ans != -1 {
		 	return ans
		 }
	 }
	 return -1
}
