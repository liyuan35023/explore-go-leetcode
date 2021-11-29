package cn
//给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。 
//
// 示例 1： 
//
//输入：nums = [1,5,11,5]
//输出：true
//解释：数组可以分割成 [1, 5, 5] 和 [11] 。 
//
// 示例 2： 
//
//输入：nums = [1,2,3,5]
//输出：false
//解释：数组不能分割成两个元素和相等的子集。
//
// 提示： 
//
// 1 <= nums.length <= 200
// 1 <= nums[i] <= 100 


//leetcode submit region begin(Prohibit modification and deletion)

func canPartition(nums []int) bool {
	total := 0
	maxNum := 0
	n := len(nums)
	for _, v := range nums {
		total += v
		maxNum = max(maxNum, v)
	}
	if total % 2 != 0 || maxNum > total / 2 {
		return false
	}
	if total / 2 == maxNum {
		return true
	}
	// dp
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, total / 2 + 1)
	}
	dp[0][nums[0]] = true
	for i := 1; i < n; i++ {
		for j := 1; j < total / 2 + 1; j++ {
			if j >= nums[i] {
				dp[i][j] = dp[i-1][j-nums[i]] || dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n-1][total/2]
}






//func canPartition(nums []int) bool {
//	total := nums[0]
//	maxNum := nums[0]
//	for i := 1; i < len(nums); i++ {
//		total += nums[i]
//		maxNum = max(maxNum, nums[i])
//	}
//	if total % 2 != 0 || total / 2 < maxNum {
//		return false
//	}
//	if total / 2 == maxNum {
//		return true
//	}
//	half := total / 2
//	// dp
//	// dp[i][j] 表示 下标从0 到 i是否可以选出一些数字的和==j
//	dp := make([][]bool, len(nums))
//	for i := 0; i < len(nums); i++ {
//		dp[i] = make([]bool, half+1)
//	}
//	dp[0][nums[0]] = true
//	for i := 1; i < len(nums); i++ {
//		for j := 1; j < half+1; j++ {
//			dp[i][j] = dp[i-1][j] || dp[i][half-j] || j >= nums[i] && dp[i-1][j-nums[i]]
//		}
//	}
//	return dp[len(nums)-1][half]
//}
//

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//leetcode submit region end(Prohibit modification and deletion)
