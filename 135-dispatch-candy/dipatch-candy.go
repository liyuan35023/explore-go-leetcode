package _35_dispatch_candy

/*

	老师想给孩子们分发糖果，有 N个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。

	你需要按照以下要求，帮助老师给这些孩子分发糖果：

	每个孩子至少分配到 1 个糖果。
	相邻的孩子中，评分高的孩子必须获得更多的糖果。
	那么这样下来，老师至少需要准备多少颗糖果呢？

	示例1:

	输入: [1,0,2]
	输出: 5
	解释: 你可以分别给这三个孩子分发 2、1、2 颗糖果。
	示例2:

	输入: [1,2,2]
	输出: 4
	解释: 你可以分别给这三个孩子分发 1、2、1 颗糖果。
		 第三个孩子只得到 1 颗糖果，这已满足上述两个条件。
 */
func Candy(ratings []int) int {
	// 动态规划
	if len(ratings) < 2 {
		return len(ratings)
	}
	n := len(ratings)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			dp[i] = dp[i-1] + 1
		}
	}
	// 再从右向左
	for i := n-2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && dp[i] <= dp[i+1] {
			dp[i] = dp[i+1] + 1
		}
	}

	tmpTotal := 0
	for i := 0; i < n; i++ {
		tmpTotal += dp[i]
	}
	return tmpTotal
}