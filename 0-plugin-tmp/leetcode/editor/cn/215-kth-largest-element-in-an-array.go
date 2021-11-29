package cn

import (
	"math/rand"
	"time"
)

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
	rand.Seed(time.Now().UnixNano())
	left, right := 0, len(nums) - 1
	for left < right {
		mid := randomPartition(nums, left, right)
		if mid == k - 1 {
			return nums[mid]
		} else if mid > k - 1 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}

func randomPartition(nums []int, left, right int) int {
	randomIdx := rand.Intn(right-left+1) + left
	nums[left], nums[randomIdx] = nums[randomIdx], nums[left]
	pivot := nums[left]
	for left < right {
		for left < right && nums[right] <= pivot {
			right--
		}
		nums[left] = nums[right]
		for left < right && nums[left] >= pivot {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = pivot
	return left
}

//leetcode submit region end(Prohibit modification and deletion)
