package _6_remove_duplicates_sorted_array

/*
	给定一个有序数组 nums，对数组中的元素进行去重，使得原数组中的每个元素只有一个。最后返回去重以后数组的长度值。

	Example 1:
	Given nums = [1,1,2],

	Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively.

	It doesn't matter what you leave beyond the returned length.

	Example 2:
	Given nums = [0,0,1,1,1,2,2,3,3,4],

	Your function should return length = 5, with the first five elements of nums being modified to 0, 1, 2, 3, and 4 respectively.

	It doesn't matter what values are set beyond the returned length.
 */

func removeDuplicates(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	ans := 1
	curNum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != curNum {
			nums[ans] = nums[i]
			ans++
			curNum = nums[i]
		}
	}
	return ans
}