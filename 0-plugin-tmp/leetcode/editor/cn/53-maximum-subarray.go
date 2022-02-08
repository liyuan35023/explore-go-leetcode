package cn

//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
// 示例 1： 
//
//输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
//输出：6
//解释：连续子数组[4,-1,2,1] 的和最大，为6 。
//
// 示例 2：
//
//输入：nums = [1]
//输出：1
//
// 示例 3： 
//
//输入：nums = [0]
//输出：0
//
// 示例 4： 
//
//输入：nums = [-1]
//输出：-1
//
// 示例 5： 
//
//输入：nums = [-100000]
//输出：-100000
//
// 提示： 
// 1 <= nums.length <= 3 * 104
// -105 <= nums[i] <= 105 
// 

//leetcode submit region begin(Prohibit modification and deletion)
func maxSubArray(nums []int) int {
	dp := nums[0]
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		dp = max(nums[i], dp+nums[i])
		ans = max(ans, dp)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
//leetcode submit region end(Prohibit modification and deletion)
