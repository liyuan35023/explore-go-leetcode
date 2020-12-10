package _5_jump_game

/*
	Example 1:

	Input: [2,3,1,1,4]
	Output: true
	Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
	Example 2:

	Input: [3,2,1,0,4]
	Output: false
	Explanation: You will always arrive at index 3 no matter what. Its maximum
				 jump length is 0, which makes it impossible to reach the last index.
	题目大意 #
	给定一个非负整数数组，最初位于数组的第一个位置。数组中的每个元素代表在该位置可以跳跃的最大长度。判断是否能够到达最后一个位置。

*/

func canJump(nums []int) bool {
	if len(nums) < 1 {
		return true
	}
	maxLocation := 0
	for i := 0; i < len(nums)-1 && i <= maxLocation; i++ {
		maxLocation = max(maxLocation, i+nums[i])
		if maxLocation >= len(nums) - 1 {
			 return true
		}
	}
	return maxLocation >= len(nums) - 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
