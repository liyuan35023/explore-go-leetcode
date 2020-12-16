package _69_majority_element

/*

	Example 1:

	Input: [3,2,3]
	Output: 3

	Example 2:
	Input: [2,2,1,1,1,2,2]
	Output: 2
	题目大意 #
	给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。你可以假设数组是非空的，并且给定的数组总是存在众数。
 */

func majorityElement(nums []int) int {
	ans := nums[0]
	validLen := 1
	for i := 1; i < len(nums); i++ {
		if ans == nums[i] {
			validLen++
		} else {
			if validLen > 0 {
				validLen--
			} else {
				ans = nums[i]
				validLen = 1
			}
		}
	}
	return ans
}
