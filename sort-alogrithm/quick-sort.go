package sort_alogrithm

func QuickSort(nums []int, left int, right int) {
	if left < right {
		mid := partition(nums, left, right)
		QuickSort(nums, left, mid-1)
		QuickSort(nums, mid+1, right)
	}
}

func partition(nums []int, left int, right int) int {
	pivot := nums[left]
	for left < right {
		// 从右往左，找到第一个小于pivot的index
		for left < right && nums[right] >= pivot {
			right--
		}
		nums[left] = nums[right]
		// 从左往右，找到第一个大于pivot的index
		for left < right && nums[left] <= pivot {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = pivot
	return left
}

