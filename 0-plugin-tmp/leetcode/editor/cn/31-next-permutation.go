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
	idx := len(nums) - 2
	for idx >= 0 && nums[idx] >= nums[idx+1] {
		idx--
	}
	if idx >= 0 {
		firstBiggerIdx := getFirstBigger(nums[idx+1:], nums[idx]) + idx + 1
		nums[idx], nums[firstBiggerIdx] = nums[firstBiggerIdx], nums[idx]
	}
	for i, j := idx + 1, len(nums) - 1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func getFirstBigger(nums []int, target int) int {
	l, r := 0, len(nums) - 1
	for l <= r {
		mid := l + (r - l) / 2
		if nums[mid] <= target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return r

}


//leetcode submit region end(Prohibit modification and deletion)
