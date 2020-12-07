package _9_combination_sum

import "sort"

/*
	给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

	candidates 中的数字可以无限制重复被选取。

	Example 1:
	Input: candidates = [2,3,6,7], target = 7,
	A solution set is:
	[
	  [7],
	  [2,2,3]
	]

	Example 2:
	Input: candidates = [2,3,5], target = 8,
	A solution set is:
	[
	  [2,2,2,2],
	  [2,3,3],
	  [3,5]
	]

	Note:
	1 All numbers (including target) will be positive integers.
	2 The solution set must not contain duplicate combinations.

 */

var ans [][]int
func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) < 1 {
		return nil
	}
	sort.Ints(candidates)

	c := make([]int, 0)
	ans = make([][]int, 0)
	backTrace(candidates, target, 0, c)
	return ans
}

func backTrace(candidates []int, target int, start int, tmp []int) {
	if target == 0 {
		solution := make([]int, len(tmp))
		copy(solution, tmp)
		ans = append(ans, solution)
		return
	}
	if target < 0 {
		return
	}
	for i := start; i < len(candidates); i++ {
		if candidates[i] > target {
			break
		}
		tmp = append(tmp, candidates[i])
		backTrace(candidates, target - candidates[i], i, tmp)
		tmp = tmp[:len(tmp)-1]
	}
	return
}

func combinationSumII(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	if len(candidates) < 1 {
		return ans
	}
	sort.Ints(candidates)
	solution := make([]int, 0)
	var dfs func(idx int, target int)
	dfs = func(idx int, target int) {
		if target == 0 {
			ans = append(ans, append([]int{}, solution...))
			return
		}
		for i := idx; i < len(candidates); i++ {
			if candidates[i] > target {
				break
			}
			solution = append(solution, candidates[i])
			dfs(i, target-candidates[i])
			solution = solution[:len(solution)-1]
		}

	}
	dfs(0, target)
	return ans
}
