package _16_partition_equal_subset_sum

/*
	Example 1:

	Input: [1, 5, 11, 5]

	Output: true

	Explanation: The array can be partitioned as [1, 5, 5] and [11].
	Example 2:

	Input: [1, 2, 3, 5]

	Output: false

	Explanation: The array cannot be partitioned into equal sum subsets.
	题目大意 #
	给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

	注意:
	每个数组中的元素不会超过 100
	数组的大小不会超过 200

 */

func canPartition(nums []int) bool {
	// calculate total and get max Number
	total := 0
	maxNum := 0
	for _, v := range nums {
		total += v
		maxNum = max(maxNum, v)
	}
	if total % 2 != 0 {
		return false
	}
	if maxNum > total / 2 {
		return false
	} else if maxNum == total / 2 {
		return true
	}
	dp := make([][]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]bool, total / 2 + 1)
	}
	dp[0][nums[0]] = true
	for i := 1; i < len(nums); i++ {
		for j := 1; j < total / 2 + 1; j++ {
			dp[i][j] = dp[i-1][j] || j > nums[i] && dp[i-1][j-nums[i]]
		}
	}
	return dp[len(nums)-1][total/2]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
