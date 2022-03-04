package cn
//给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。 
//
// 示例 1：
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// 
// 示例 2：
//
//输入：nums = [0,1]
//输出：[[0,1],[1,0]]
//
// 示例 3： 
//
//输入：nums = [1]
//输出：[[1]]
//
// 提示： 
//
// 1 <= nums.length <= 6 
// -10 <= nums[i] <= 10 
// nums 中的所有整数 互不相同 
// 

//leetcode submit region begin(Prohibit modification and deletion)
func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	var dfs



}
//func permute(nums []int) [][]int {
//	n := len(nums)
//	ans := make([][]int, 0)
//	used := make(map[int]bool)
//	var dfs func(solve []int)
//	dfs = func(solve []int) {
//		if len(solve) == n {
//			ans = append(ans, append([]int{}, solve...))
//			return
//		}
//		for i := 0; i < n; i++ {
//			if used[i] {
//				continue
//			}
//			used[i] = true
//			solve = append(solve, nums[i])
//			dfs(solve)
//			solve = solve[:len(solve)-1]
//			used[i] = false
//		}
//	}
//	dfs([]int{})
//	return ans
//}
//leetcode submit region end(Prohibit modification and deletion)
