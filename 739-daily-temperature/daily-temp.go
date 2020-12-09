package _39_daily_temperature
/*
	For example, given the list of temperatures T = [73, 74, 75, 71, 69, 72, 76, 73], your output should be [1, 1, 4, 2, 1, 1, 0, 0].

	Note: The length of temperatures will be in the range [1, 30000]. Each temperature will be an integer in the range [30, 100].

	题目大意 #
	给出一个温度数组，要求输出比当天温度高的在未来的哪一天，输出未来第几天的天数。例如比 73 度高的在未来第 1 天出现，比 75 度高的在未来第 4 天出现。

 */

func dailyTemperatures(T []int) []int {
	// 单调栈
	// 单调不增栈
	ans := make([]int, len(T))
	stack := make([]int, 0)
	for i := 0; i < len(T); i++ {
		for len(stack) > 0 && T[i] > T[stack[len(stack)-1]] {
			ans[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}
