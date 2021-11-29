package cn
//有 n 个花园，按从 1 到 n 标记。另有数组 paths ，其中 paths[i] = [xi, yi] 描述了花园 xi 到花园 yi 的双向路径。在
//每个花园中，你打算种下四种花之一。 
//
// 另外，所有花园 最多 有 3 条路径可以进入或离开. 
//
// 你需要为每个花园选择一种花，使得通过路径相连的任何两个花园中的花的种类互不相同。 
//
// 以数组形式返回 任一 可行的方案作为答案 answer，其中 answer[i] 为在第 (i+1) 个花园中种植的花的种类。花的种类用 1、2、3、4 
//表示。保证存在答案。 
//
//
// 示例 1： 
//
// 
//输入：n = 3, paths = [[1,2],[2,3],[3,1]]
//输出：[1,2,3]
//解释：
//花园 1 和 2 花的种类不同。
//花园 2 和 3 花的种类不同。
//花园 3 和 1 花的种类不同。
//因此，[1,2,3] 是一个满足题意的答案。其他满足题意的答案有 [1,2,4]、[1,4,2] 和 [3,2,1]
// 
//
// 示例 2： 
//
// 
//输入：n = 4, paths = [[1,2],[3,4]]
//输出：[1,2,1,2]
// 
//
// 示例 3： 
//
// 
//输入：n = 4, paths = [[1,2],[2,3],[3,4],[4,1],[1,3],[2,4]]
//输出：[1,2,3,4]
//
// 提示： 
//
// 
// 1 <= n <= 104 
// 0 <= paths.length <= 2 * 104 
// paths[i].length == 2 
// 1 <= xi, yi <= n 
// xi != yi 
// 每个花园 最多 有 3 条路径可以进入或离开 
// 


//leetcode submit region begin(Prohibit modification and deletion)
func gardenNoAdj(n int, paths [][]int) []int {
	//ans := make([]int, n)
	//exclude := make([]map[int]struct{}, n)
	//for i := 0; i < n; i++ {
	//	exclude[i] = make(map[int]struct{})
	//}
	//
	//for _, p := range paths {
	//	exclude[p[0]-1][p[1]-1] = struct{}{}
	//	exclude[p[1]-1][p[0]-1] = struct{}{}
	//}
	//excludeType := make(map[int]int)
	//for i := 0; i < n; i++ {
	//	for excludeIdx := range exclude[i] {
	//		excludeType[ans[excludeIdx]] = 1
	//	}
	//	for j := 1; j <= 4; j++ {
	//		if _, ok := excludeType[j]; !ok {
	//			ans[i] = j
	//		}
	//	}
	//}
	//
	//return ans

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
//leetcode submit region end(Prohibit modification and deletion)
