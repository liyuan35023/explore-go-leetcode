package cn
//给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。 
//
// 示例 1： 
//
//输入：n = 3
//输出：[[1,2,3],[8,9,4],[7,6,5]]
//
// 示例 2： 
//
//输入：n = 1
//输出：[[1]]
//
// 提示：
//
// 1 <= n <= 20
// 

//leetcode submit region begin(Prohibit modification and deletion)
func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}
	top, bottom := 0, n-1
	left, right := 0, n-1
	cur := 1
	for top < bottom && left < right {
		for i := left; i < right; i++ {
			ans[top][i] = cur
			cur++
		}
		for i := top; i < bottom; i++ {
			ans[i][right] = cur
			cur++
		}
		for i := right; i > left; i-- {
			ans[bottom][i] = cur
			cur++
		}
		for i := bottom; i > top; i-- {
			ans[i][left] = cur
			cur++
		}
		top, bottom = top+1, bottom-1
		left, right = left+1, right-1
	}
	if top == bottom && left == right {
		ans[top][left] = cur
	}
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
