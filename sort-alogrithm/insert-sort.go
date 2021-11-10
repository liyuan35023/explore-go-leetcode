package sort_alogrithm

func InsertSort(nums []int) {
	idx := 1
	for ; idx < len(nums); idx++ {
		tmp := nums[idx]
		r := idx -1
		for ; r >= 0 && tmp < nums[r]; r-- {
			nums[r+1] = nums[r]
		}
		nums[r+1] = tmp
	}
}
