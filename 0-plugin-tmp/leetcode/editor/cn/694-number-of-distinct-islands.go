package cn

import (
	"strconv"
)

//给定一个非空 01 二维数组表示的网格，一个岛屿由四连通（上、下、左、右四个方向）的 1 组成，你可以认为网格的四周被海水包围。
//
// 请你计算这个网格中共有多少个形状不同的岛屿。两个岛屿被认为是相同的，当且仅当一个岛屿可以通过平移变换（不可以旋转、翻转）和另一个岛屿重合。 
//
//
// 示例 1： 
//
//11000
//11000
//00011
//00011
// 
//
// 给定上图，返回结果 1 。 
//
// 示例 2： 
//
//11011
//10000
//00001
//11011 
//
// 给定上图，返回结果 3 。 
// 
//注意： 
//
// 11
//1
//
// 和 
//
//  1
//11
//
// 是不同的岛屿，因为我们不考虑旋转、翻转操作。 
//
// 
//
// 提示：二维数组每维的大小都不会超过 50 。

//leetcode submit region begin(Prohibit modification and deletion)
func numDistinctIslands(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	island := make(map[string]int)
	var dfs func(row, column int, startRow, startColumn int) string
	dfs = func(row, column int, startRow, startColumn int) string {
		grid[row][column] = 0
		str := strconv.Itoa(row-startRow) + strconv.Itoa(column-startColumn)
		for _, dir := range directions {
			newRow, newColumn := row + dir[0], column + dir[1]
			if newRow >= 0 && newRow < m && newColumn >= 0 && newColumn < n && grid[newRow][newColumn] == 1 {
				str += dfs(newRow, newColumn, startRow, startColumn)
			}
		}
		return str
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				island[dfs(i, j, i, j)] = 1
			}
		}
	}
	return len(island)

}
//leetcode submit region end(Prohibit modification and deletion)
