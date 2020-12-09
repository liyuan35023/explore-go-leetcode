package _6_permutations

/*
	给定一个 没有重复 数字的序列，返回其所有可能的全排列。

	示例:

	输入: [1,2,3]
	输出:
	[
	  [1,2,3],
	  [1,3,2],
	  [2,1,3],
	  [2,3,1],
	  [3,1,2],
	  [3,2,1]
	]
 */

func permute(nums []int) [][]int {
	var ans [][]int
	n := len(nums)
	usedMap := make([]bool, n)
	solution := make([]int, 0)
	var dfs func(index int)

	dfs = func(index int) {
		if index > n {
			ans = append(ans, append([]int{}, solution...))
			return
		}
		for i := 0; i < n; i++ {
			if !usedMap[i] {
				solution = append(solution, nums[i])
				usedMap[i] = true
				dfs(index+1)
				solution = solution[:len(solution)-1]
				usedMap[i] = false
			}
		}
	}
	dfs(1)
	return ans
}

func PermuteII(nums []int) [][]int {
	ans := make([][]int, 0)
	solution := make([]int, 0)
	usedMap := make(map[int]bool)
	for k := range nums {
		backTrace(k, nums, solution, &ans, usedMap)
	}
	return ans
}

func backTrace(curIndex int, nums []int, solution []int, ans *[][]int, usedMap map[int]bool) {
	solution = append(solution, nums[curIndex])
	usedMap[curIndex] = true
	defer func() {
		solution = solution[:len(solution)-1]
		usedMap[curIndex] = false
	}()
	if len(solution) == len(nums) {
		tmp := make([]int, len(solution))
		copy(tmp, solution)
		*ans = append(*ans, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if usedMap[i] {
			continue
		}
		backTrace(i, nums, solution, ans, usedMap)
	}

}
