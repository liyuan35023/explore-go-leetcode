package cn

import "sort"

//给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
//
// 示例 1： 
//
//输入：nums = [1,1,2]
//输出：
//[[1,1,2],
// [1,2,1],
// [2,1,1]]
//
// 示例 2： 
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// 
// 提示：
//
// 1 <= nums.length <= 8
// -10 <= nums[i] <= 10 
// 

//leetcode submit region begin(Prohibit modification and deletion)
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	used := make([]bool, len(nums))
	ans := make([][]int, 0)
	var dfs func(idx int, solve []int)
	dfs = func(idx int, solve []int) {
		if idx == len(nums) {
			ans = append(ans, append([]int{}, solve...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] || i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			dfs(idx+1, append(solve, nums[i]))
			used[i] = false
		}
	}
	dfs(0, []int{})
	return ans
}
//func permuteUnique(nums []int) [][]int {
//	ans := make([][]int, 0)
//	sort.Ints(nums)
//	used := make(map[int]bool)
//	var dfs func(solve []int)
//	dfs = func(solve []int) {
//		if len(solve) == len(nums) {
//			ans = append(ans, append([]int{}, solve...))
//			return
//		}
//		for i := 0; i < len(nums); i++ {
//			if used[i] || i > 0 && !used[i-1] && nums[i] == nums[i-1] {
//				continue
//			}
//			solve = append(solve, nums[i])
//			used[i] = true
//			dfs(solve)
//			used[i] = false
//			solve = solve[:len(solve)-1]
//		}
//	}
//	dfs([]int{})
//	return ans
//}
//leetcode submit region end(Prohibit modification and deletion)
