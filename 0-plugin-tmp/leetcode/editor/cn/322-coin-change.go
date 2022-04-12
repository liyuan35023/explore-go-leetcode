package cn

//给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
//
// 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。 
//
// 你可以认为每种硬币的数量是无限的。 
//
// 示例 1：
//
//输入：coins = [1, 2, 5], amount = 11
//输出：3 
//解释：11 = 5 + 5 + 1 
//
// 示例 2： 
//
//输入：coins = [2], amount = 3
//输出：-1 
//
// 示例 3： 
//
//输入：coins = [1], amount = 0
//输出：0
// 
// 示例 4：
//
//输入：coins = [1], amount = 1
//输出：1
//
// 示例 5： 
//
//输入：coins = [1], amount = 2
//输出：2
//
// 提示： 
//
// 1 <= coins.length <= 12 
// 1 <= coins[i] <= 231 - 1 
// 0 <= amount <= 104 

//leetcode submit region begin(Prohibit modification and deletion)
func coinChange(coins []int, amount int) int {

}




//func coinChange(coins []int, amount int) int {
//	// dp
//	// dp[i] 表示用coins组成i，所需要的最少硬币数
//	dp := make([]int, amount+1)
//	for i := 1; i < amount+1; i++ {
//		minCount := 1 << 31 - 1
//		for _, v := range coins {
//			if i >= v && dp[i-v] >= 0 {
//				minCount = min(minCount, dp[i-v]+1)
//			}
//		}
//		if minCount == 1 << 31 - 1 {
//			dp[i] = -1
//		} else {
//			dp[i] = minCount
//		}
//	}
//	return dp[amount]
//
//}
//func coinChange(coins []int, amount int) int {
//	// dp
//	// dp[i] 表示硬币组成i，所需要的最少数目
//	dp := make([]int, amount+1)
//	for i := 1; i < amount+1; i++ {
//		minCount := 1<<31 - 1
//		for _, v := range coins {
//			if i - v >= 0 && dp[i-v] != -1 {
//				minCount = min(minCount, dp[i-v])
//			}
//		}
//		if minCount == 1 << 31 - 1 {
//			dp[i] = -1
//		} else {
//			dp[i] = minCount + 1
//		}
//	}
//	return dp[amount]

	//// dfs (超时了)
	//ans := 1<<31-1
	//sort.Ints(coins)
	//var dfs func(idx int, remain int, curCount int)
	//dfs = func(idx int, remain int, curCount int) {
	//	if idx < 0 || curCount + remain / coins[idx] >= ans {
	//		return
	//	}
	//	if remain % coins[idx] == 0 {
	//		ans = min(ans, curCount+remain/coins[idx])
	//		return
	//	}
	//	maxCount := remain / coins[idx]
	//	for maxCount >= 0 {
	//		dfs(idx - 1, remain - coins[idx]*maxCount, curCount + maxCount)
	//		maxCount--
	//	}
	//}
	//dfs(len(coins)-1, amount, 0)
	//if ans == 1 << 31 - 1 {
	//	return -1
	//}
	//return ans
//}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
//leetcode submit region end(Prohibit modification and deletion)
