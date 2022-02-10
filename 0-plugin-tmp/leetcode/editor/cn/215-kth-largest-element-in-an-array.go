package cn

import "math/rand"

// 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//
// 示例 1: 
//
//输入: [3,2,1,5,6,4] 和 k = 2
//输出: 5
//
// 示例 2: 
//
//输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
//输出: 4 
//
// 提示： 
//
// 1 <= k <= nums.length <= 104 
// -104 <= nums[i] <= 104 


//leetcode submit region begin(Prohibit modification and deletion)
func findKthLargest(nums []int, k int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := partition(nums, l, r)
		if mid == k - 1 {
			return nums[mid]
		} else if mid > k - 1 {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return nums[l]
}

func partition(nums []int, l, r int) int {
	randIdx := rand.Intn(r - l + 1) + l
	nums[l], nums[randIdx] = nums[randIdx], nums[l]
	pivot := nums[l]
	for l < r {
		for l < r && nums[r] <= pivot {
			r--
		}
		nums[l] = nums[r]
		for l < r && nums[l] >= pivot {
			l++
		}
		nums[r] = nums[l]
	}
	nums[l] = pivot
	return l
}


//leetcode submit region end(Prohibit modification and deletion)
