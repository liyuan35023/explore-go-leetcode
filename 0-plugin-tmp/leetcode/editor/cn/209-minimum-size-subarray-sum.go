package cn
//给定一个含有 n 个正整数的数组和一个正整数 target 。 
//
// 找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长
//度。如果不存在符合条件的子数组，返回 0 。 
//
//
// 示例 1： 
//
//输入：target = 7, nums = [2,3,1,2,4,3]
//输出：2
//解释：子数组 [4,3] 是该条件下的长度最小的子数组。
//
// 示例 2： 
//
//输入：target = 4, nums = [1,4,4]
//输出：1
// 
// 示例 3：
//
//输入：target = 11, nums = [1,1,1,1,1,1,1,1]
//输出：0
//
// 提示： 
//
// 1 <= target <= 10⁹
// 1 <= nums.length <= 10⁵ 
// 1 <= nums[i] <= 10⁵ 
// 
//
// 进阶：
//
// 如果你已经实现 O(n) 时间复杂度的解法, 请尝试设计一个 O(n log(n)) 时间复杂度的解法。


//leetcode submit region begin(Prohibit modification and deletion)
func minSubArrayLen(target int, nums []int) int {
	sum := make([]int, 0)
	ans := 0

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			sum = append(sum, nums[i])
		} else {
			sum = append(sum, sum[i-1]+nums[i])
		}
		if sum[i] >= target {
			idx := findFirstBiggerIdx(sum, sum[i]-target)
			if ans == 0 {
				ans = i - idx + 1
			} else {
				ans = min(ans, i - idx + 1)
			}
		}
	}
	return ans
}

func findFirstBiggerIdx(nums []int, target int) int {
	left, right := 0, len(nums) - 1
	for left < right {
		mid := left + (right - left) / 2
		if nums[mid] <= target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		}
	}
	return left
}


func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
//leetcode submit region end(Prohibit modification and deletion)
