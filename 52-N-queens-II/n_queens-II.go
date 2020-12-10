package _2_N_queens_II

/*
	Example:
	Input: 4
	Output: 2
	Explanation: There are two distinct solutions to the 4-queens puzzle as shown below.
	[
	 [".Q..",  // Solution 1
	  "...Q",
	  "Q...",
	  "..Q."],

	 ["..Q.",  // Solution 2
	  "Q...",
	  "...Q",
	  ".Q.."]
	]
	题目大意 #
	给定一个整数 n，返回 n 皇后不同的解决方案的数量

 */

func totalNQueens1(n int) int {
	ans := 0
	if n < 2 {
		return n
	}
	columnRecord := make([]bool, n)
	diagnalRecord := make([]bool, 2*n-1)
	diagnalReverseRecord := make([]bool, 2*n-1)
	var dfs func(row int)
	dfs = func(row int) {
		if row == n {
			ans++
			return
		}
		for column := 0; column < n; column++ {
			if columnRecord[column] || diagnalRecord[row+column] || diagnalReverseRecord[row-column+n-1] {
				continue
			}
			columnRecord[column] = true
			diagnalRecord[row+column] = true
			diagnalReverseRecord[row-column+n-1] = true
			dfs(row+1)
			columnRecord[column] = false
			diagnalRecord[row+column] = false
			diagnalReverseRecord[row-column+n-1] = false
		}
	}
	dfs(0)
	return ans
}
