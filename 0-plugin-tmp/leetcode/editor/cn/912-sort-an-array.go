package cn

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
	// merge sort

	var sortHelper func(left, right int)
	sortHelper = func(left, right int) {
		if left >= right {
			return
		}
		mid := left + (right - left) / 2
		sortHelper(left, mid)
		sortHelper(mid+1, right)
		tmp := make([]int, right - left + 1)
		i, j := left, mid + 1
		for cur := 0; cur < right - left + 1; cur++ {
			if j > right {
				tmp[cur] = nums[i]
				i++
			} else if i > mid {
				tmp[cur] = nums[j]
				j++
			} else {
				if nums[i] < nums[j] {
					tmp[cur] = nums[i]
					i++
				} else {
					tmp[cur] = nums[j]
					j++
				}
			}
		}
		copy(nums[left:right+1], tmp)
	}
	sortHelper(0, len(nums)-1)
	return nums
}

//leetcode submit region end(Prohibit modification and deletion)
