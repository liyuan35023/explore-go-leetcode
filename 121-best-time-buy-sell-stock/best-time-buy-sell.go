package _21_best_time_buy_sell_stock

/*
	Say you have an array for which the ith element is the price of a given stock on day i.

	If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock),
	design an algorithm to find the maximum profit.

	Note that you cannot sell a stock before you buy one.

	Example 1:

	Input: [7,1,5,3,6,4]
	Output: 5
	Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
				 Not 7-1 = 6, as selling price needs to be larger than buying price.
	Example 2:

	Input: [7,6,4,3,1]
	Output: 0
	Explanation: In this case, no transaction is done, i.e. max profit = 0.
*/

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	n := len(prices)

	dp := make([]int, n) // dp[i] 表示第i+1天卖出，可以获得的最大利润

	ans := 0
	dp[0] = 0
	minimumPrice := prices[0]
	for i := 1; i < n; i++ {
		dp[i] = max(prices[i]-minimumPrice, 0)
		ans = max(ans, dp[i])
		minimumPrice = min(minimumPrice, prices[i])
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
func maxProfitIINoDp(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	ans := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		ans = max(ans, prices[i]-minPrice)
		minPrice = min(minPrice, prices[i])
	}
	return ans
}
