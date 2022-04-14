package cn

import "math/rand"

//给你一个整数数组 nums，请你将该数组升序排列。
//
// 示例 1：
//
// 输入：nums = [5,2,3,1]
//输出：[1,2,3,5]
//
// 示例 2： 
//
// 输入：nums = [5,1,1,2,0,0]
//输出：[0,0,1,1,2,5]
//
// 提示： 
//
// 1 <= nums.length <= 50000 
// -50000 <= nums[i] <= 50000 
// 
// Related Topics 数组 分治 桶排序 计数排序 基数排序 排序 堆（优先队列） 归并排序 👍 359 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func sortArray(nums []int) []int {
	// 三路快排
	var sortHelper func(nums []int, left, right int)
	sortHelper = func(nums []int, left, right int) {
		if left >= right {
			return
		}
		p0, p2 := partition(nums, left, right)
		sortHelper(nums, left, p0-1)
		sortHelper(nums, p2+1, right)
	}
	sortHelper(nums, 0, len(nums)-1)
	return nums
}

func partition(nums []int, left, right int) (int, int) {
	randIdx := left + rand.Intn(right - left + 1)
	nums[left], nums[randIdx] = nums[randIdx], nums[left]
	pivot := nums[left]
	p0, p2 := left, right
	for i := left; i <= p2; {
		if nums[i] == pivot {
			i++
			continue
		} else if nums[i] > pivot {
			nums[i], nums[p2] = nums[p2], nums[i]
			p2--
		} else {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
			i++
		}
	}
	return p0, p2
}

//leetcode submit region end(Prohibit modification and deletion)
