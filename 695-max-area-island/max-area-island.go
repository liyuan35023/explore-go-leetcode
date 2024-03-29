package _95_max_area_island

/*
	Example 1:
	[[0,0,1,0,0,0,0,1,0,0,0,0,0],
	 [0,0,0,0,0,0,0,1,1,1,0,0,0],
	 [0,1,1,0,1,0,0,0,0,0,0,0,0],
	 [0,1,0,0,1,1,0,0,1,0,1,0,0],
	 [0,1,0,0,1,1,0,0,1,1,1,0,0],
	 [0,0,0,0,0,0,0,0,0,0,1,0,0],
	 [0,0,0,0,0,0,0,1,1,1,0,0,0],
	 [0,0,0,0,0,0,0,1,1,0,0,0,0]]
	Given the above grid, return 6. Note the answer is not 11, because the island must be connected 4-directionally.

	Example 2:

	[[0,0,0,0,0,0,0,0]]
	Given the above grid, return0.

	Note: The length of each dimension in the given grid does not exceed 50.

	题目大意 #
	给定一个包含了一些 0 和 1 的非空二维数组 grid 。一个 岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在水平或者竖直方向上相邻。
	你可以假设 grid 的四个边缘都被 0（代表水）包围着。找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为 0 。)

 */

func maxAreaOfIsland(grid [][]int) int {
	ans := 0
	m, n := len(grid), len(grid[0])
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var dfs func(row, column int) int
	dfs = func(row, column int) int {
		if row < 0 || row >= m || column < 0 || column >= n || grid[row][column] == 0 {
			return 0
		}
		grid[row][column] = 0
		area := 0
		for _, dir := range directions {
			newRow := row + dir[0]
			newColumn := column + dir[1]
			area += dfs(newRow, newColumn)
		}
		return area
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans = max(ans, dfs(i, j))
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
