package _6_valid_sodoku

/*
	判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。

	数字 1-9 在每一行只能出现一次。
	数字 1-9 在每一列只能出现一次。
	数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。

 */
func isValidSudoku(board [][]byte) bool {
	rowRecord := make([][]bool, 9)
	columnRecord := make([][]bool, 9)
	blockRecord := make([][]bool, 9)
	for i := 0; i < 9; i++ {
		rowRecord[i] = make([]bool, 9)
		columnRecord[i] = make([]bool, 9)
		blockRecord[i] = make([]bool, 9)
	}
	
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			if board[row][column] != '.' {
				num := int(board[row][column] - '0') - 1
				blockIdx := (row / 3) * 3 + column / 3
				if rowRecord[row][num] || columnRecord[column][num] || blockRecord[blockIdx][num] {
					return false
				}
				rowRecord[row][num] = true
				columnRecord[column][num] = true
				blockRecord[blockIdx][num] = true
			}
		}
	}
	return true
}