package cn

import "strconv"

//编写一个程序，通过填充空格来解决数独问题。
//
// 数独的解法需 遵循如下规则： 
//
// 
// 数字 1-9 在每一行只能出现一次。 
// 数字 1-9 在每一列只能出现一次。 
// 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图） 
// 
//
// 数独部分空格内已填入了数字，空白格用 '.' 表示。 
//
// 示例：
//
//输入：board = [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5","."
//,".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".","."
//,"3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"
//],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],["
//.",".",".",".","8",".",".","7","9"]]
//输出：[["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8"
//],["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],["
//4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],["9","
//6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4","
//5","2","8","6","1","7","9"]]
//解释：输入的数独如上图所示，唯一有效的解决方案如下所示：
//
// 提示： 
//
// 
// board.length == 9 
// board[i].length == 9 
// board[i][j] 是一位数字或者 '.' 
// 题目数据 保证 输入数独仅有一个解 


//leetcode submit region begin(Prohibit modification and deletion)
func solveSudoku(board [][]byte) {
	rowMap, columnMap, blockMap, spaces := generateMapAndSpace(board)
	var dfs func(idx int) bool
	dfs = func(idx int) bool {
		if idx == len(spaces) {
			return true
		}
		row, column := spaces[idx][0], spaces[idx][1]
		blockId := row / 3 * 3 + column / 3
		for i := 1; i <= 9; i++ {
			if rowMap[row][i] || columnMap[column][i] || blockMap[blockId][i] {
				continue
			}
			byteNum := strconv.Itoa(i)
			board[row][column] = byteNum[0]
			rowMap[row][i] = true
			columnMap[column][i] = true
			blockMap[blockId][i] = true
			if dfs(idx+1) {
				return true
			}
			board[row][column] = '.'
			rowMap[row][i] = false
			columnMap[column][i] = false
			blockMap[blockId][i] = false
		}
		return false
	}
	dfs(0)
}

func generateMapAndSpace(board [][]byte) (rowMap []map[int]bool, columnMap []map[int]bool, blockMap []map[int]bool,
	spaces [][2]int) {
	rowMap = make([]map[int]bool, 9)
	columnMap = make([]map[int]bool, 9)
	blockMap = make([]map[int]bool, 9)
	for i := 0; i < 9; i++ {
		rowMap[i] = make(map[int]bool)
		columnMap[i] = make(map[int]bool)
		blockMap[i] = make(map[int]bool)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				spaces = append(spaces, [2]int{i, j})
			} else {
				num := int(board[i][j] - '0')
				blockId := i / 3 * 3 + j / 3
				rowMap[i][num] = true
				columnMap[j][num] = true
				blockMap[blockId][num] = true

			}
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
