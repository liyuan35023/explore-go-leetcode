package _22_coin_change

import (
	"fmt"
	"math"
)

/*
	Example 1:
	Input: coins = [1, 2, 5], amount = 11
	Output: 3
	Explanation: 11 = 5 + 5 + 1

	Example 2:
	Input: coins = [2], amount = 3
	Output: -1
	Note:

	You may assume that you have an infinite number of each kind of coin.

	题目大意 #
	给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。

 */

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	// dp[amount] 表示组成amount最少使用的硬币数目
	dp := make([]int, amount+1)
	return coinHelper(coins, amount, dp)
}

func coinHelper(coins []int, amount int, dp []int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	if dp[amount] != 0 {
		return dp[amount]
	}
	min := math.MaxInt32
	for _, v := range coins {
		tmp := coinHelper(coins, amount-v, dp)
		if tmp >= 0 && min > tmp + 1 {
			min = tmp + 1
		}
	}
	if amount == 11 {
		fmt.Println("stop")
	}
	if min < math.MaxInt32 {
		dp[amount] = min
	} else {
		dp[amount] = -1
	}
	return dp[amount]
}

func coinChange2(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	// dp[i]表示组成i所需要的最小的硬币数
	// dp[i] = min(dp[i-c1],..., dp[i-cn]) + 1
	dp := make([]int, amount+1)
	for i := 1; i < amount + 1; i++ {
		minNum := math.MaxInt32
		for _, c := range coins {
			if i - c >= 0 && dp[i-c] >= 0 {
				minNum = min(minNum, dp[i-c])
			}
		}
		if minNum != math.MaxInt32 {
			dp[i] = minNum + 1
		} else {
			dp[i] = -1
		}
	}
	return dp[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

