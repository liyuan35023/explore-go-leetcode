package _7_permutations_II

import "sort"

/*
	给定一个可包含重复数字的序列，返回所有不重复的全排列。

	示例:

	输入: [1,1,2]
	输出:
	[
	  [1,1,2],
	  [1,2,1],
	  [2,1,1]
	]

 */
var ans [][]int
func PermuteUnique(nums []int) [][]int {
	if len(nums) < 1 {
		return nil
	}

	sort.Ints(nums)
	ans = make([][]int, 0)
	used := make([]bool, len(nums))
	solution := make([]int, 0)
	backTrace(nums, used, 0, solution)
	return ans
}

func backTrace(nums []int, used []bool, index int, solution []int) {
	if index == len(nums) {
		tmp := make([]int, len(solution))
		copy(tmp, solution)
		ans = append(ans, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}
		if used[i] {
			continue
		}
		solution = append(solution, nums[i])
		used[i] = true
		backTrace(nums, used, index+1, solution)
		solution = solution[:len(solution)-1]
		used[i] = false
	}
}

func PermuteUniqueII(nums []int) [][]int {
	ans := make([][]int, 0)
	if len(nums) < 1 {
		return ans
	}
	n := len(nums)
	sort.Ints(nums)
	solution := make([]int, 0)
	used := make([]bool, n)
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx > n {
			ans = append(ans, append([]int{}, solution...))
			return
		}
		for i := 0; i < n; i++ {
			if used[i] || i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			solution = append(solution, nums[i])
			used[i] = true
			dfs(idx+1)
			used[i] = false
			solution = solution[:len(solution)-1]
		}
	}
	dfs(1)
	return ans
}

