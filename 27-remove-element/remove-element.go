package _7_remove_element

/*
	给定一个数组 nums 和一个数值 val，将数组中所有等于 val 的元素删除，并返回剩余的元素个数。

	Example 1:
	Given nums = [3,2,2,3], val = 3,

	Your function should return length = 2, with the first two elements of nums being 2.

	It doesn't matter what you leave beyond the returned length.

	Example 2:
	Given nums = [0,1,2,2,3,0,4,2], val = 2,

	Your function should return length = 5, with the first five elements of nums containing 0, 1, 3, 0, and 4.

	Note that the order of those five elements can be arbitrary.
	It doesn't matter what values are set beyond the returned length.
 */

func removeElement(nums []int, val int) int {
	if len(nums) < 1 {
		return 0
	}
	ans := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[ans] = nums[i]
			ans++
		}
	}
	return ans
}