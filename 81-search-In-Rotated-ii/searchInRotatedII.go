package _1_search_In_Rotated

/*
	Example 1:

	Input: nums = [2,5,6,0,0,1,2], target = 0
	Output: true
	Example 2:

	Input: nums = [2,5,6,0,0,1,2], target = 3
	Output: false
	Follow up:

	This is a follow up problem to  Search in Rotated Sorted Array, where nums may contain duplicates.
	Would this affect the run-time complexity? How and why?
	题目大意 #
	假设按照升序排序的数组在预先未知的某个点上进行了旋转。( 例如，数组 [0,0,1,2,2,5,6] 可能变为 [2,5,6,0,0,1,2] )。

	编写一个函数来判断给定的目标值是否存在于数组中。若存在返回 true，否则返回 false。

	进阶:

	这是搜索旋转排序数组 的延伸题目，本题中的 nums 可能包含重复元素。
	这会影响到程序的时间复杂度吗？会有怎样的影响，为什么？
 */

func search(nums []int, target int) bool {
	if len(nums) < 1 {
		return false
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right - left) / 2
		if nums[mid] == target {
			return true
		}

		if nums[left] < nums[mid] {
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[left] > nums[mid] {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			// ==
			left++
		}

	}
	return false
}

