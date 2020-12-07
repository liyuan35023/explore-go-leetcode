package _0_combination_sum_II

import "sort"

/*
	给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

	candidates 中的每个数字在每个组合中只能使用一次。

	Example 1:
	Input: candidates = [10,1,2,7,6,1,5], target = 8,
	A solution set is:
	[
	  [1, 7],
	  [1, 2, 5],
	  [2, 6],
	  [1, 1, 6]
	]

	Example 2:
	Input: candidates = [2,5,2,1,2], target = 5,
	A solution set is:
	[
	  [1,2,2],
	  [5]
	]
 */

var ans [][]int

func CombinationSum2(candidates []int, target int) [][]int {
	if len(candidates) < 1 {
		return nil
	}
	sort.Ints(candidates)
	ans = make([][]int, 0)
	c := make([]int, 0)
	backTrace(candidates, target, 0, c)
	return ans
}

func backTrace(candidates []int, target int, index int, tmp []int) {
	if target == 0 {
		solution := make([]int, len(tmp))
		copy(solution, tmp)
		ans = append(ans, solution)
		return
	} else if target < 0 {
		return
	}
	
	for i := index; i < len(candidates); i++ {
		if candidates[i] > target {
			break
		}

		if i > index && candidates[i] == candidates[i-1] {
			continue
		}
		tmp = append(tmp, candidates[i])
		backTrace(candidates, target-candidates[i], i+1, tmp)
		tmp = tmp[:len(tmp)-1]
	}
}

func CombinationSum2II(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	if len(candidates) < 1 {
		return ans
	}
	solution := make([]int, 0)
	sort.Ints(candidates)
	var dfs func(idx int, target int)
	dfs = func(idx int, target int) {
		if target == 0 {
			ans = append(ans, append([]int{}, solution...))
			return
		}
		for i := idx; i < len(candidates); i++ {
			if candidates[i] > target {
				return
			}
			if i > idx && candidates[i] == candidates[i-1] {
				continue
			}
			solution = append(solution, candidates[i])
			dfs(i+1, target-candidates[i])
			solution = solution[:len(solution)-1]
		}
	}
	dfs(0, target)
	return ans
}

