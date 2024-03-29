package cn

import (
	"sort"
)

//给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
// candidates 中的每个数字在每个组合中只能使用一次。 
//
// 说明： 
//
// 
// 所有数字（包括目标数）都是正整数。 
// 解集不能包含重复的组合。 
// 
//
// 示例 1: 
//
// 输入: candidates = [10,1,2,7,6,1,5], target = 8,
//所求解集为:
//[
//  [1, 7],
//  [1, 2, 5],
//  [2, 6],
//  [1, 1, 6]
//]
// 
//
// 示例 2: 
//
// 输入: candidates = [2,5,2,1,2], target = 5,
//所求解集为:
//[
//  [1,2,2],
//  [5]
//] 

//leetcode submit region begin(Prohibit modification and deletion)
func combinationSum2(candidates []int, target int) [][]int {


}
//func combinationSum2(candidates []int, target int) [][]int {
//	sort.Ints(candidates)
//	ans := make([][]int, 0)
//	var dfs func(idx int, solution []int, sum int)
//	dfs = func(idx int, solution []int, sum int) {
//		if sum == target {
//			ans = append(ans, append([]int{}, solution...))
//			return
//		}
//		if idx >= len(candidates) || target - sum < candidates[idx] {
//			return
//		}
//		for i := idx; i < len(candidates); i++ {
//			if i > idx && candidates[i-1] == candidates[i] {
//				continue
//			}
//			dfs(i+1, append(solution, candidates[i]), sum+candidates[i])
//		}
//	}
//	dfs(0, []int{}, 0)
//	return ans
//}
//leetcode submit region end(Prohibit modification and deletion)
