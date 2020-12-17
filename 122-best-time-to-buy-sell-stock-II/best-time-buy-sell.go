package _22_best_time_to_buy_sell_stock_II

/*
	Example 1:

	Input: [7,1,5,3,6,4]
	Output: 7
	Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
				 Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
	Example 2:

	Input: [1,2,3,4,5]
	Output: 4
	Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
				 Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are
				 engaging multiple transactions at the same time. You must sell before buying again.
	Example 3:

	Input: [7,6,4,3,1]
	Output: 0
	Explanation: In this case, no transaction is done, i.e. max profit = 0.
	题目大意 #
	给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
	注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

 */

func maxProfit122(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	profit := 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] > min && (i == len(prices)-1 || prices[i] > prices[i+1]) {
			profit += prices[i]-min
			min = prices[i]
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return profit
}

func maxProfit122DP(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	// dp[i][0]表示第i天处于空仓状态的最大收益
	// dp[i][1]表示第i天处于持有股票状态的最大收益
	dp := make([][2]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = [2]int{0, 0}
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[len(prices)-1][0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y

}

