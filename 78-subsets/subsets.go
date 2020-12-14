package _8_subsets

/*
	Example:

	Input: nums = [1,2,3]
	Output:
	[
	  [3],
	  [1],
	  [2],
	  [1,2,3],
	  [1,3],
	  [2,3],
	  [1,2],
	  []
	]
	题目大意 #
	给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。说明：解集不能包含重复的子集。
 */

func subsets(nums []int) [][]int {
	ans := [][]int{[]int{}}
	if len(nums) < 1 {
		return ans
	}
	solution := make([]int, 0)
	backTrace(nums, 0, solution, &ans)
	return ans
}

func backTrace(nums []int, idx int, solution []int, ans *[][]int) {
	if idx == len(nums) {
		return
	}

	// 不用当前的idx的值
	backTrace(nums, idx+1, solution, ans)

	solution = append(solution, nums[idx])
	tmp := make([]int, len(solution))
	copy(tmp, solution)
	*ans = append(*ans, tmp)
	backTrace(nums, idx+1, solution, ans)
}

func subsetsII(nums []int) [][]int {
	ans := [][]int{[]int{}}

	for _, v := range nums {
		n := len(ans)
		for i := 0; i < n; i++ {
			ans = append(ans, append([]int{v}, ans[i]...))
		}
	}
	return ans
}
