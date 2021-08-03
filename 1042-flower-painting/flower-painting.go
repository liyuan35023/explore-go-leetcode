package _042_flower_painting

func gardenNoAdj(n int, paths [][]int) []int {
	ans := make([]int, n)
	// 建立邻接图
	graph := make([]map[int]struct{}, n)
	for i := 0; i < n; i++ {
		graph[i] = make(map[int]struct{})
	}
	for _, p := range paths {
		graph[p[0]-1][p[1]-1] = struct{}{}
		graph[p[1]-1][p[0]-1] = struct{}{}
	}

	for i := 0; i < n; i++ {
		exceptType := make(map[int]struct{})
		for k := range graph[i] {
			if ans[k] != 0 {
				exceptType[ans[k]] = struct{}{}
			}
		}
		for j := 1; j <= 4; j++ {
			if _, ok := exceptType[j]; !ok {
				ans[i] = j
				break
			}
		}
	}
	return ans
}

