package _5_jump_game_II

/*
	给定一个非负整数数组，最初位于数组的第一个位置
	数组中每个元素代表，在该位置可以跳跃的最大长度
	目标是使用最少的步数， 到达数组的最后一个位置

	example：

	input: [2,3,1,1,4]
	output: 2
 */

func Jump(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	step := 0
	n := len(nums)
	maxLocation := 0
	curStepMax := 0
	for i := 0; i < len(nums) && maxLocation < n-1; i++ {
		curStepMax = max(curStepMax, i+nums[i])
		if i == maxLocation {
			step++
			maxLocation = curStepMax
		}
	}
	return step
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
