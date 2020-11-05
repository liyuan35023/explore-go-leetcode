package _00_num_of_island

/*
	给定一个由 ‘1’（陆地）和 ‘0’（水）组成的的二维网格，计算岛屿的数量。
	一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。

	Example 1:

	Input:
	11110
	11010
	11000
	00000

	Output: 1

	Example 2:

	Input:
	11000
	11000
	00100
	00011

	Output: 3
 */

func numIslands(grid [][]byte) int {
	if len(grid) < 1 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])

	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				ans++
				// find island ++
				recordIsland(i, j, grid)
			}
		}
	}
	return ans
}

func recordIsland(i, j int, grid [][]byte) {
	m := len(grid)
	n := len(grid[0])
	if i < 0 || i >= m || j < 0 || j >= n {
		return
	}
	if grid[i][j] == '1' {
		grid[i][j] = '0'
		recordIsland(i-1, j, grid)
		recordIsland(i+1, j, grid)
		recordIsland(i, j-1, grid)
		recordIsland(i, j+1, grid)
	}
}

