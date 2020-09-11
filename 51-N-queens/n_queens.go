package _1_N_queens


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

