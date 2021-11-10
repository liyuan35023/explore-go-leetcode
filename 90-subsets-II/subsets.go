package _0_subsets_II

import "sort"

func subsetsWithDup(nums []int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(nums)
	used := make(map[int]bool)
	var dfs func(idx int, solve []int)
	dfs = func(idx int, solve []int) {
		if idx >= len(nums) {
			ans = append(ans, append([]int{}, solve...))
			return
		}
		if idx > 0 && nums[idx-1] == nums[idx] && !used[idx-1] {
			dfs(idx+1, solve)
		} else {
			dfs(idx+1, solve)
			used[idx] = true
			dfs(idx+1, append(solve, nums[idx]))
			used[idx] = false
		}
	}
	dfs(0, []int{})
	return ans
}
