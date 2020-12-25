package _3_search_in_rotated_orderd_array

/*
	假设按照升序排序的数组在预先未知的某个点上进行了旋转。( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
	搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。你可以假设数组中不存在重复的元素。

	你的算法时间复杂度必须是 O(log n) 级别。

	Example 1:
	Input: nums = [4,5,6,7,0,1,2], target = 0
	Output: 4

	Example 2:
	Input: nums = [4,5,6,7,0,1,2], target = 3
	Output: -1
 */


func search(nums []int, target int) int {
	if len(nums) < 1 {
		return -1
	}
	if len(nums) == 1 {
		if target == nums[0] {
			return 0
		} else {
			return -1
		}
	}

	length := len(nums)
	head := nums[0]
	mid := nums[length/2]

	switch target {
	case nums[0]:
		return 0
	case nums[length/2]:
		return length/2
	case nums[length-1]:
		return length-1
	}

	if head < mid {
		if target > head && target < mid {
			return search(nums[:length/2], target)
		} else {
			ret := search(nums[length/2 + 1:], target)
			if ret == -1 {
				return -1
			}
			return length/2 + 1 + ret
		}
	} else {
		if target > mid && target < head {
			ret := search(nums[length/2 + 1:], target)
			if ret == -1 {
				return -1
			}
			return length/2 + 1 + ret
		} else {
			return search(nums[:length/2], target)
		}
	}
}

func SearchLoop(nums []int, target int) int {
	if len(nums) < 1 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right - left) / 2
		if nums[mid] == target {
			return mid
		}
		// 一定是小于等于
		if nums[left] <= nums[mid] {
			if target < nums[mid] && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
