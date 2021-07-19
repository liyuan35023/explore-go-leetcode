package cn

import "golang.org/x/net/html/atom"

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
// 进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。


//leetcode submit region begin(Prohibit modification and deletion)
func maxSubArray(nums []int) int {
	var sumHelper func(l, r int) (lSum, rSum, maxSub, totalSum int)
	sumHelper = func(l, r int) (lSum, rSum, maxSub, totalSum int) {
		if l == r {
			return nums[l], nums[l], nums[l], nums[l]
		}
		mid := l + (r - l) / 2
		lSum1, rSum1, maxSub1, totalSum1 := sumHelper(l, mid)
		lSum2, rSum2, maxSub2, totalSum2 := sumHelper(mid+1, r)
		return max(lSum1, totalSum1+lSum2), max(rSum2, totalSum2+rSum1), max(rSum1+lSum2, max(maxSub1, maxSub2)), totalSum1+totalSum2
	}
	_, _, ans, _ := sumHelper(0, len(nums)-1)
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
//leetcode submit region end(Prohibit modification and deletion)
