package _1_N_queens

import "bytes"

var ret [][]string
func solveNQueens(n int) [][]string {
	ret = make([][]string, 0)
	queens := make([]int, n)
	for k := range queens {
		queens[k] = -1
	}
	column := make(map[int]bool)
	diagonal1 := make(map[int]bool)
	diagonal2 := make(map[int]bool)
	backTrace(queens, n, 0, column, diagonal1, diagonal2)
	return ret
}

func backTrace(queens []int, n int, curQueen int, column map[int]bool, diagonal1, diagonal2 map[int]bool) {
	if n == curQueen {
		// 说明全部放进合适位置了
		oneSolution := generateMap(queens)
		ret = append(ret, oneSolution)
		return
	}

	// 从第一列开始遍历，判断是否可以放第curQueen个queen
	for i := 0; i < n; i++ {
		if column[i] || diagonal1[curQueen-i] || diagonal2[curQueen+i] {
			continue
		}
		queens[curQueen] = i
		column[i] = true
		diagonal1[curQueen-i] = true
		diagonal2[curQueen+i] = true
		backTrace(queens, n, curQueen+1, column, diagonal1, diagonal2)
		delete(column, i)
		delete(diagonal1, curQueen-i)
		delete(diagonal2, curQueen+i)
	}
}

func generateMap(queens []int) []string {
	n := len(queens)
	ret := make([]string, n)
	for k, v := range queens {
		for i := 0; i < n; i++ {
			if i == v {
				ret[k] = ret[k] + "Q"
			} else {
				ret[k] = ret[k] + "."
			}
		}
	}
	return ret
}

func solveNQueensII(n int) [][]string {
	ans := make([][]string, 0)
	if n < 1 {
		return ans
	}
	columnRecord := make([]bool, n)
	diagRecord, diagReverseRecord := make([]bool, 2*n-1), make([]bool, 2*n-1)
	queensColumn := make([]int, n)

	var dfs func(rowIdx int)
	dfs = func(rowIdx int) {
		if rowIdx == n {
			ans = append(ans, generate(queensColumn))
			return
		}
		for column := 0; column < n; column++ {
			if columnRecord[column] || diagRecord[rowIdx+column] || diagReverseRecord[rowIdx-column+n-1] {
				continue
			}
			columnRecord[column] = true
			diagRecord[rowIdx+column] = true
			diagReverseRecord[rowIdx-column+n-1] = true
			queensColumn[rowIdx] = column
			dfs(rowIdx+1)
			columnRecord[column] = false
			diagRecord[rowIdx+column] = false
			diagReverseRecord[rowIdx-column+n-1] = false
			queensColumn[rowIdx] = 0
		}
	}
	dfs(0)
	return ans
}

func generate(queens []int) []string {
	ans := make([]string, 0)
	for i := 0; i < len(queens); i++ {
		column := queens[i]
		buf := bytes.NewBuffer([]byte{})
		for j := 0; j < len(queens); j++ {
			if j == column {
				buf.WriteByte('Q')
			} else {
				buf.WriteByte('.')
			}
		}
		ans = append(ans, buf.String())
	}
	return ans
}
