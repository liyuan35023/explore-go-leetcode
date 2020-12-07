package _7_sudoku_solver

import "strconv"

/*
	题目大意 #
	编写一个程序，通过已填充的空格来解决数独问题。一个数独的解法需遵循如下规则：

	数字 1-9 在每一行只能出现一次。
	数字 1-9 在每一列只能出现一次。
	数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
	空白格用 ‘.’ 表示。
 */

func solveSudoku(board [][]byte) {
	rowRecord, columnRecord, blockRecord := generateRecord(board)
	var dfs func(row int, column int) bool
	dfs = func(row int, column int) bool {
		if row >= 9 {
			return true
		}
		var ret bool
		if board[row][column] != '.' {
			if column == 8 {
				ret = dfs(row+1, 0)
			} else {
				ret = dfs(row, column+1)
			}
			return ret
		}
		blockIdx := (row / 3) * 3 + column / 3
		for num := 1; num <= 9; num++ {
			numIdx := num - 1
			if rowRecord[row][numIdx] || columnRecord[column][numIdx] || blockRecord[blockIdx][numIdx] {
				continue
			}
			byteNum := []byte(strconv.Itoa(num))
			board[row][column] = byteNum[0]
			rowRecord[row][numIdx] = true
			columnRecord[column][numIdx] = true
			blockRecord[blockIdx][numIdx] = true

			if column == 8 {
				ret = dfs(row+1, 0)
			} else {
				ret = dfs(row, column+1)
			}
			if ret {
				return true
			}
			board[row][column] = '.'
			rowRecord[row][numIdx] = false
			columnRecord[column][numIdx] = false
			blockRecord[blockIdx][numIdx] = false
		}
		return false
	}
	dfs(0, 0)
}

func generateRecord(board [][]byte) ([][]bool, [][]bool, [][]bool) {
	rowRecord, columnRecord, blockRecord := make([][]bool, 9), make([][]bool, 9), make([][]bool, 9)
	for i := 0; i < 9; i++ {
		rowRecord[i] = make([]bool, 9)
		columnRecord[i] = make([]bool, 9)
		blockRecord[i] = make([]bool, 9)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := int(board[i][j] - '0') - 1
				blockIdx := (i / 3) * 3 + j / 3
				rowRecord[i][num] = true
				columnRecord[j][num] = true
				blockRecord[blockIdx][num] = true
			}
		}
	}
	return rowRecord, columnRecord, blockRecord
}
