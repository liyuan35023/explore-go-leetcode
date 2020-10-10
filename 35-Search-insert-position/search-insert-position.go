package _5_Search_insert_position

/*
	给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
	你可以假设数组中无重复元素。

	Example 1:
	Input: [1,3,5,6], 5
	Output: 2

	Example 2:
	Input: [1,3,5,6], 2
	Output: 1

	Example 3:
	Input: [1,3,5,6], 7
	Output: 4

	Example 4:
	Input: [1,3,5,6], 0
	Output: 0

 */
func searchInsert(nums []int, target int) int {
	if len(nums) < 0 {
		return -1
	}
	i, j := 0, len(nums)-1
	for i <= j {
		mid := (i + j) / 2
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			if mid == 0 || target > nums[mid-1] {
				return mid
			}
			j = mid - 1
		} else if target > nums[mid] {
			if mid == len(nums)-1 || target < nums[mid+1] {
				return mid + 1
			}
			i = mid + 1
		}
	}
	return -1
}