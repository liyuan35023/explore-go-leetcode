package cn

import "sort"

//给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
// candidates 中的数字可以无限制重复被选取。 
//
// 说明：
// 
// 所有数字（包括 target）都是正整数。 
// 解集不能包含重复的组合。 
//
// 示例 1： 
//
// 输入：candidates = [2,3,6,7], target = 7,
//所求解集为：
//[
//  [7],
//  [2,2,3]
//]
//
// 示例 2： 
//
// 输入：candidates = [2,3,5], target = 8,
//所求解集为：
//[
// [2,2,2,2],
// [2,3,3],
// [3,5]
//]
// 提示： 
//
// 
// 1 <= candidates.length <= 30 
// 1 <= candidates[i] <= 200 
// candidate 中的每个元素都是独一无二的。 
// 1 <= target <= 500 
//

//leetcode submit region begin(Prohibit modification and deletion)
func combinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(candidates)
	var dfs func(idx int, solve []int, remain int)
	dfs = func(idx int, solve []int, remain int) {
		if remain == 0 {
			ans = append(ans, append([]int{}, solve...))
			return
		}

		
	}






}
//func combinationSum(candidates []int, target int) [][]int {
//	sort.Ints(candidates)
//	ans := make([][]int, 0)
//
//	var dfs func(idx int, curSum int, solution []int)
//	dfs = func(idx int, curSum int, solution []int) {
//		if curSum == target {
//			ans = append(ans, append([]int{}, solution...))
//			return
//		}
//		if idx >= len(candidates) || target - curSum < candidates[idx] {
//			return
//		}
//		dfs(idx, curSum+candidates[idx], append(solution, candidates[idx]))
//		dfs(idx+1, curSum, solution)
//	}
//	dfs(0, 0, []int{})
//	return ans
//}
//leetcode submit region end(Prohibit modification and deletion)
