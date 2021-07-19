package cn

//实现获取 下一个排列 的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
//
// 如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
//
// 必须 原地 修改，只允许使用额外常数空间。
//
// 示例 1：
//
//输入：nums = [1,2,3]
//输出：[1,3,2]
//
// 示例 2：
//
//输入：nums = [3,2,1]
//输出：[1,2,3]
//
// 示例 3：
//
//输入：nums = [1,1,5]
//输出：[1,5,1]
//
// 示例 4：
//
//输入：nums = [1]
//输出：[1]
//
// 提示：
//
// 1 <= nums.length <= 100
// 0 <= nums[i] <= 100
//

//leetcode submit region begin(Prohibit modification and deletion)
func nextPermutation(nums []int) {
	// find last 升序序列
	lastAscendIdx := -1
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			lastAscendIdx = i - 1
			nextBiggerIdx := binarySearchNextBiggerIdx(nums[i:], nums[lastAscendIdx]) + lastAscendIdx + 1
			nums[lastAscendIdx], nums[nextBiggerIdx] = nums[nextBiggerIdx], nums[lastAscendIdx]
			break
		}
	}
	// swap
	left, right := lastAscendIdx+1, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

func binarySearchNextBiggerIdx(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target && (mid == n-1 || nums[mid+1] <= target) {
			return mid
		}
		if nums[mid] <= target && (mid == 0 || nums[mid-1] > target) {
			return mid - 1
		}

		if nums[mid] <= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

//leetcode submit region end(Prohibit modification and deletion)
