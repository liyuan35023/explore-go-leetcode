package _36_single_number

/*
	Example 1:

	Input: [2,2,1]
	Output: 1

	Example 2:
	Input: [4,1,2,1,2]
	Output: 4

	题目大意 #
	给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。要求算法时间复杂度是线性的，并且不使用额外的辅助空间。

 */

func singleNumber(nums []int) int {
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		ans ^= nums[i]
	}
	return ans
}
