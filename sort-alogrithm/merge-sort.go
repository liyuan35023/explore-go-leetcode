package sort_alogrithm

// 稳定排序

func MergeSort(nums []int) []int {
	result := make([]int, len(nums))
	mergeTwo(nums, result, 0, len(nums)-1)
	return result
}

func mergeTwo(nums []int, result []int, left, right int) {
	if left == right {
		return
	}
	mid := (right-left)/2 + left
	mergeTwo(nums, result, left, mid)
	mergeTwo(nums, result, mid+1, right)
	k := left
	i, j := left, mid+1
	for ; i <= mid && j <= right; k++ {
		if nums[i] < nums[j] {
			result[k] = nums[i]
			i++
		} else {
			result[k] = nums[j]
			j++
		}
	}
	for ; i <= mid; i, k = i+1, k+1 {
		result[k] = nums[i]
	}
	for ; j <= right; j, k = j+1, k+1 {
		result[k] = nums[j]
	}
	for n := left; n <= right; n++ {
		nums[n] = result[n]
	}
}

func MergeSortLoop(nums []int) []int {
	result := make([]int, len(nums))
	n := len(nums)

	for block := 1; block < 2*n; block*=2 {
		for i := 0; i < n; i += 2*block {
			start := i
			end := i + 2*block - 1
			if end >= n {
				end = n-1
			}
			mid := start + (end - start) / 2
			i ,j := start, mid + 1
			k := start
			for ; i <= mid && j <= end; k++ {
				if nums[i] < nums[j] {
					result[k] = nums[i]
					i++
				} else {
					result[k] = nums[j]
					j++
				}
			}
			for ; i <= mid; i, k = i+1, k+1 {
				result[k] = nums[i]
			}
			for ; j <= end; j, k = j+1, k+1 {
				result[k] = nums[j]
			}
			for n := start; n <= end; n++ {
				nums[n] = result[n]
			}
		}
	}
	return result
}