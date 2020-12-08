package _1_first_positive_missing

/*
Example 1:
Input: [1,2,0]
Output: 3

Example 2:
Input: [3,4,-1,1]
Output: 2

Example 3:
Input: [7,8,9,11,12]
Output: 1

Note:
Your algorithm should run in O(n) time and uses constant extra space.

题目大意 #
找到缺失的第一个正整数。

 */

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 1; i <= n; i++ {
		idx := i - 1
		for nums[idx] != i {
			if nums[idx] >= 1 && nums[idx] <= n && nums[nums[idx]-1] != nums[idx] {
				nums[idx], nums[nums[idx]-1] = nums[nums[idx]-1], nums[idx]
			} else {
				break
			}
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i+1
		}
	}
	return n+1
}
