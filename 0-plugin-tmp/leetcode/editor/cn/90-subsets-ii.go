package cn

import "sort"

//给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
//
// 解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。 
//
// 示例 1：
//
//输入：nums = [1,2,2]
//输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
//
// 示例 2： 
//
//输入：nums = [0]
//输出：[[],[0]]
//
// 提示： 
//
// 1 <= nums.length <= 10 
// -10 <= nums[i] <= 10 


//leetcode submit region begin(Prohibit modification and deletion)
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	used := make([]bool, len(nums))
	ans := make([][]int, 0)
	var dfs func(idx int, solution []int)
	dfs = func(idx int, solution []int) {
		if idx >= len(nums) {
			ans = append(ans, append([]int{}, solution...))
			return
		}
		dfs(idx+1, solution)
		if idx > 0 && nums[idx-1] == nums[idx] && !used[idx-1] {
			return
		}
		used[idx] = true
		dfs(idx+1, append(solution, nums[idx]))
		used[idx] = false
	}
	dfs(0, []int{})
	return ans

}

//func subsetsWithDup(nums []int) [][]int {
//	ans := make([][]int, 0)
//	sort.Ints(nums)
//	var dfs func(idx int, solve []int)
//	used := make([]bool, len(nums))
//	dfs = func(idx int, solve []int) {
//		if idx == len(nums) {
//			ans = append(ans, append([]int{}, solve...))
//			return
//		}
//		dfs(idx+1, solve)
//		if idx > 0 && nums[idx] == nums[idx-1] && !used[idx-1] {
//			return
//		}
//		used[idx] = true
//		dfs(idx+1, append(solve, nums[idx]))
//		used[idx] = false
//	}
//	dfs(0, []int{})
//	return ans
//}
//leetcode submit region end(Prohibit modification and deletion)
