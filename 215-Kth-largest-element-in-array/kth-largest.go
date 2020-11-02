package _15_Kth_largest_element_in_array

import (
	"math/rand"
	"time"
)

/*
	Example 1:
	Input: [3,2,1,5,6,4] and k = 2
	Output: 5

	Example 2:
	Input: [3,2,3,1,2,4,5,5,6] and k = 4
	Output: 4

	Note:
	You may assume k is always valid, 1 ≤ k ≤ array’s length.

	题目大意 #
	找出数组中第 K 大的元素。这一题非常经典。可以用 O(n) 的时间复杂度实现。
 */

func FindKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSearch(nums, k, 0, len(nums)-1)
}

func quickSearch(nums []int, k, left, right int) int {
	mid := partition(nums, left, right)
	if mid+1 < k {
		return quickSearch(nums, k, mid+1, right)
	} else if mid+1 > k {
		return quickSearch(nums, k, left, mid-1)
	} else {
		return nums[mid]
	}
}

func randomPartition(nums []int, left, right int) int {
	randomIndex := rand.Int() % (right - left + 1) + left
	nums[randomIndex], nums[right] = nums[right], nums[randomIndex]
	return partition(nums, left, right)
}

func partition(nums []int, left, right int) int {
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
