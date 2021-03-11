package _37_single_number_II

/*
	Example 1:

	Input: [2,2,3,2]
	Output: 3
	Example 2:

	Input: [0,1,0,1,0,1,99]
	Output: 99
	题目大意 #
	给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。找出那个只出现了一次的元素。
	要求算法时间复杂度是线性的，并且不使用额外的辅助空间。
 */

func singleNumberIII(nums []int) int {
	once, twice := 0, 0
	for i := 0; i < len(nums); i++ {
		once = ^twice & (once ^ nums[i])
		twice = ^once & (twice ^ nums[i])
	}
	return once
}

