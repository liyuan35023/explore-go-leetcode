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
	var dfs func(l, r int) (int, int, int, int)
	dfs = func(l, r int) (int, int, int, int) {
		if l > r {
			return 0, 0, 0, 0
		}
		if l == r {
			return nums[l], nums[l], nums[l], nums[l]
		}
		mid := l + (r - l) / 2
		lMax1, rMax1, max1, total1 := dfs(l, mid)
		lMax2, rMax2, max2, total2 := dfs(mid+1, r)

		return max(lMax1, total1+lMax2), max(rMax2, total2+rMax1), max(max1, max(max2, rMax1+lMax2)), total1 + total2
	}
	_, _, ans, _ := dfs(0, len(nums)-1)
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
//leetcode submit region end(Prohibit modification and deletion)
