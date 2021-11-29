package cn
//给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。 
//
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。 
//
// 此外，你可以假设该网格的四条边均被水包围。 
//
// 示例 1： 
//
//输入：grid = [
//  ["1","1","1","1","0"],
//  ["1","1","0","1","0"],
//  ["1","1","0","0","0"],
//  ["0","0","0","0","0"]
//]
//输出：1
// 
//
// 示例 2： 
//
// 
//输入：grid = [
//  ["1","1","0","0","0"],
//  ["1","1","0","0","0"],
//  ["0","0","1","0","0"],
//  ["0","0","0","1","1"]
//]
//输出：3
//
// 提示： 
// m == grid.length
// n == grid[i].length 
// 1 <= m, n <= 300 
// grid[i][j] 的值为 '0' 或 '1' 

//leetcode submit region begin(Prohibit modification and deletion)
//func numIslands(grid [][]byte) int {
//	// bfs
//	ans := 0
//	m, n := len(grid), len(grid[0])
//	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
//	for i := 0; i < m; i++ {
//		for j := 0; j < n; j++ {
//			if grid[i][j] == '1' {
//				ans++
//				grid[i][j] = '0'
//				queue := make([][]int, 0)
//				queue = append(queue, []int{i, j})
//				for len(queue) != 0 {
//					row, column := queue[0][0], queue[0][1]
//					queue = queue[1:]
//					for _, dir := range directions {
//						newRow, newColumn := row + dir[0], column + dir[1]
//						if newRow >= 0 && newRow < m && newColumn >= 0 && newColumn < n && grid[newRow][newColumn] == '1' {
//							grid[newRow][newColumn] = '0'
//							queue = append(queue, []int{newRow, newColumn})
//						}
//					}
//
//				}
//			}
//		}
//	}
//	return ans
//}

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	ans := 0
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var dfs func(row, column int) int
	dfs = func(row, column int) int {
	grid[row][column] = '0'
	area := 1
	for _, dir := range directions {
	newRow, newColumn := row + dir[0], column + dir[1]
	if newRow >= 0 && newRow < m && newColumn >= 0 && newColumn < n && grid[newRow][newColumn] == '1' {
		area += dfs(newRow, newColumn)
	}
	}
	return area
	}

	for i := 0; i < m; i++ {
	for j := 0; j < n; j++ {
	if grid[i][j] == '1' {
	ans++
	dfs(i, j)
	}
	}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
