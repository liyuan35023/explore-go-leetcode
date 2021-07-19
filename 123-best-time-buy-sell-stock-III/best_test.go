package _23_best_time_buy_sell_stock_III

import (
	"fmt"
	"testing"
)

func TestBestTime(t *testing.T) {
	fmt.Println(maxProfit222([]int{1,2,3,4,5}))

}

func maxProfit222(prices []int) int {
	// dp 5种状态
	dp := make([][5]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = [5]int{0, 0, 0, 0, 0}
	}
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
	}
	return max(dp[len(prices)-1][0], max(dp[len(prices)-1][2], dp[len(prices)-1][4]))
}

