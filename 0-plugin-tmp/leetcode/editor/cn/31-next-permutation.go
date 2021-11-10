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
	if len(nums) < 2 {
		return
	}
	lastAsc := len(nums) - 2
	for ; lastAsc >= 0 && nums[lastAsc] >= nums[lastAsc+1]; lastAsc-- {
	}
	if lastAsc >= 0 {
		tmpIdx := findFirstBigger(nums[lastAsc+1:], nums[lastAsc])
		realIdx := lastAsc + tmpIdx + 1
		nums[lastAsc], nums[realIdx] = nums[realIdx], nums[lastAsc]
	}
	// reverse
	for left, right := lastAsc+1, len(nums)-1; left < right; left, right = left+1, right-1 {
		nums[left], nums[right] = nums[right], nums[left]
	}
}

func findFirstBigger(nums []int, target int) int {
	left, right := 0, len(nums) - 1
	for left <= right {
		mid := left + (right - left) / 2
		if nums[mid] <= target {
			right = mid - 1
		} else if nums[mid] > target {
			left = mid + 1
		}
	}
	return right

}

//leetcode submit region end(Prohibit modification and deletion)
