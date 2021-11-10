package sort_alogrithm

import "math/rand"

//https://zhuanlan.zhihu.com/p/341201904
// 快排时间复杂度
// 快排的优化：随机化，非递归版本的快排
func QuickSort(nums []int, left int, right int) {
	if left < right {
		mid := partition(nums, left, right)
		QuickSort(nums, left, mid-1)
		QuickSort(nums, mid+1, right)

		//p0, p2 := partition3(nums, left, right)
		//QuickSort(nums, left, p0-1)
		//QuickSort(nums, p2+1, right)
	}
}

func partition(nums []int, left int, right int) int {
	// random
	randIdx := rand.Intn(right-left+1) + left
	nums[left], nums[randIdx] = nums[randIdx], nums[left]

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

func partition3(nums []int, left int, right int) (int, int) {
	pivot := nums[left]
	p0 := left   // [left, p0)  < pivot
	i := left+1    // [p0, i) == pivot
	p2 := right  // (p2, right] > pivot

	for i <= p2 {
		if nums[i] == pivot {
			i++
		} else if nums[i] > pivot {
			nums[i], nums[p2] = nums[p2], nums[i]
			p2--
		} else if nums[i] < pivot {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
			i++
		}
	}
	return p0, p2
}

