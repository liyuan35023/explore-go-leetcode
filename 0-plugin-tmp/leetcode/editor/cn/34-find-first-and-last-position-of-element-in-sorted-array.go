package cn

//给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
//
// 如果数组中不存在目标值 target，返回 [-1, -1]。 
//
// 进阶： 
//
// 你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？ 
//
// 示例 1： 
//
//输入：nums = [5,7,7,8,8,10], target = 8
//输出：[3,4] 
//
// 示例 2： 
//
//输入：nums = [5,7,7,8,8,10], target = 6
//输出：[-1,-1] 
//
// 示例 3： 
//
// 
//输入：nums = [], target = 0
//输出：[-1,-1] 
//
// 提示： 
//
// 
// 0 <= nums.length <= 105 
// -109 <= nums[i] <= 109 
// nums 是一个非递减数组 
// -109 <= target <= 109 
// 
// Related Topics 数组 二分查找 
// 👍 1065 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func searchRange(nums []int, target int) []int {
	ans := []int{-1, -1}
	if len(nums) < 1 {
		return ans
	}
	ans[0], ans[1] = findFirstTarget(nums, target), findLastTarget(nums, target)
	return ans
}

func findFirstTarget(nums []int, target int) int {
	ans := -1
	left, right := 0, len(nums) - 1
	for left <= right {
		mid := left + (right - left) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			if mid == 0 || nums[mid-1] < target {
				return mid
			} else {
				right = mid - 1
			}
		}
	}
	return ans
}

func findLastTarget(nums []int, target int) int {
	ans := -1
	left, right := 0, len(nums) - 1
	for left <= right {
		mid := left + (right - left) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			if mid == len(nums) - 1 || nums[mid+1] > target {
				return mid
			} else {
				left = mid + 1
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
