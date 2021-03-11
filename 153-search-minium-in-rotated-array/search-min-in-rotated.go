package _53_search_minium_in_rotated_array
/*

	假设按照升序排序的数组在预先未知的某个点上进行了旋转。例如，数组[0,1,2,4,5,6,7] 可能变为[4,5,6,7,0,1,2] 。

	请找出其中最小的元素。

	示例 1：
	输入：nums = [3,4,5,1,2]
	输出：1
	示例 2：

	输入：nums = [4,5,6,7,0,1,2]
	输出：0
	示例 3：

	输入：nums = [1]
	输出：1

	提示：
	1 <= nums.length <= 5000
	-5000 <= nums[i] <= 5000
	nums 中的所有整数都是 唯一 的
	nums 原来是一个升序排序的数组，但在预先未知的某个点上进行了旋转

 */

func findMinII(nums []int) int {
	left, right := 0, len(nums) - 1
	if nums[left] < nums[right] {
		return nums[left]
	}

	for left < right {
		mid := left + (right-left) / 2
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}
		if nums[mid] < nums[mid-1] {
			return nums[mid]
		}
		if nums[mid] > nums[0] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return nums[left]
}














func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left] < nums[right] {
			return nums[left]
		}
		mid := left + (right - left) / 2
		if mid < len(nums)-1 && nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}
		if mid > 0 && nums[mid-1] > nums[mid] {
			return nums[mid]
		}
		if nums[left] <= nums[mid] {
			left = mid + 1
		} else if nums[left] > nums[mid] {
			right = mid - 1
		}
	}
	return nums[left]
}
