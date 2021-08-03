package _94_diff_islands

import (
	"bytes"
	"strconv"
)

func numDistinctIslands(grid [][]int) int {
	paths := make(map[string]struct{})
	m, n := len(grid), len(grid[0])
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var dfs func(row, column int, originRow, originColumn int, buf *bytes.Buffer)
	dfs = func(row, column int, originRow, originColumn int, buf *bytes.Buffer) {
		grid[row][column] = 0
		relativeRow := row - originRow
		relativeColumn := column - originColumn
		buf.WriteString(strconv.Itoa(relativeRow))
		buf.WriteString(strconv.Itoa(relativeColumn))
		for _, dir := range directions {
			newRow := row + dir[0]
			newColumn := column + dir[1]
			if newRow >= 0 && newRow < m && newColumn >= 0 && newColumn < n && grid[newRow][newColumn] == 1 {
				dfs(newRow, newColumn, originRow, originColumn, buf)
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				buf := bytes.NewBuffer([]byte{})
				dfs(i, j, i, j, buf)
				paths[buf.String()] = struct{}{}
			}
		}
	}
	return len(paths)
}
