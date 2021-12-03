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
	// 找最后一个升序对
	lastBiggerIdx := -1
	for i := len(nums)-1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			lastBiggerIdx = i - 1
			idx := findFirstBiggerIdx(nums[lastBiggerIdx], nums[lastBiggerIdx+1:])
			realIdx := lastBiggerIdx + idx + 1
			nums[lastBiggerIdx], nums[realIdx] = nums[realIdx], nums[lastBiggerIdx]
			break
		}
	}
	//reverse

	l, r := lastBiggerIdx+1, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l, r = l+1, r-1
	}

}

func findFirstBiggerIdx(target int, nums []int) int {
	// 降序的nums
	left, right := 0, len(nums) - 1
	for left <= right {
		mid := left + (right - left) / 2
		if nums[mid] <= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}

//leetcode submit region end(Prohibit modification and deletion)
