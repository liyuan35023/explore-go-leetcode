package _9_gray_code

func grayCode(n int) []int {
	// dfs
	ans := make([]int, 0)
	ans = append(ans, 0)
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx > n {
			return
		}
		// copy
		size := len(ans)
		for i := size - 1; i >= 0; i-- {
			ans = append(ans, ans[i])
		}
		// add 0 for front part
		for i := 0; i < size; i++ {
			ans[i] *= 2
		}

		// add 1 for second half part
		for i := size; i < 2 * size; i++ {
			ans[i] = ans[i] * 2 + 1
		}
		dfs(idx+1)
	}
	dfs(1)
	return ans
}
