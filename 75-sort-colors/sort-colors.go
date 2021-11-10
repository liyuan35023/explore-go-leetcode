package _5_sort_colors

func sortColors(nums []int) {
	// 三撸快排的思路
	pivot := 1
	p0, p2 := 0, len(nums) - 1
	i := 0
	for i <= p2 {
		if nums[i] == pivot {
			i++
		} else if nums[i] > pivot {
			nums[i], nums[p2] = nums[p2], nums[i]
			p2--
		} else {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
			i++
		}
	}
}

