package _4_find_firstandlast_position

/*
	给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
	你的算法时间复杂度必须是 O(log n) 级别。如果数组中不存在目标值，返回 [-1, -1]。

	Example 1:
	Input: nums = [5,7,7,8,8,10], target = 8
	Output: [3,4]

	Example 2:
	Input: nums = [5,7,7,8,8,10], target = 6
	Output: [-1,-1]

	这一题是经典的二分搜索变种题。二分搜索有 4 大基础变种题：

	1 查找第一个值等于给定值的元素
	2 查找最后一个值等于给定值的元素
	3 查找第一个大于等于给定值的元素
	4 查找最后一个小于等于给定值的元素
*/
func SearchRange(nums []int, target int) []int {
	return []int{SearchFirst(nums, target), SearchLast(nums, target)}
}

func NormalBinarySearch(nums []int, target int) int {
	if len(nums) < 1 {
		return -1
	}
	i, j := 0, len(nums)-1

	for i <= j {
		mid := (j + i) / 2
		switch {
		case target < nums[mid]:
			j = mid - 1
		case target > nums[mid]:
			i = mid + 1
		default:
			return mid
		}
	}
	return -1
}

func SearchFirst(nums []int, target int) int {
	if len(nums) < 1 {
		return -1
	}
	i, j := 0, len(nums)-1

	for i <= j {
		mid := (i + j) / 2
		if target == nums[mid] {
			if mid == 0 || nums[mid-1] != target {
				return mid
			} else {
				j = mid - 1
			}
		} else if target < nums[mid] {
			j = mid - 1
		} else if target > nums[mid] {
			i = mid + 1
		}
	}
	return -1
}

func SearchLast(nums []int, target int) int {
	if len(nums) < 1 {
		return -1
	}
	i, j := 0, len(nums)-1
	for i <= j {
		mid := (i + j) / 2
		if target == nums[mid] {
			if mid == len(nums)-1 || target != nums[mid+1] {
				return mid
			}
			i = mid + 1
		} else if target < nums[mid] {
			j = mid - 1
		} else if target > nums[mid] {
			i = mid + 1
		}
	}
	return -1
}

func SearchFirstBiggerEqual(nums []int, target int) int {
	if len(nums) < 1 {
		return -1
	}
	i, j := 0, len(nums)-1
	for i <= j {
		mid := (i + j) / 2
		if target <= nums[mid] {
			if mid == 0 || target > nums[mid-1] {
				return mid
			}
			j = mid - 1
		} else if target > nums[mid] {
			i = mid + 1
		}
	}
	return -1
}

func SearchLastLesserEqual(nums []int, target int) int {
	if len(nums) < 1 {
		return -1
	}
	i, j := 0, len(nums)-1
	for i < j {
		mid := (i + j) / 2
		if target >= nums[mid] {
			if mid == len(nums)-1 || target < nums[mid+1] {
				return mid
			}
			i = mid + 1
		} else if target < nums[mid] {
			j = mid - 1
		}
	}
	return -1
}
