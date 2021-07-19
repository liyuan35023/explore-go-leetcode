package cn
//给你一个 无重叠的 ，按照区间起始端点排序的区间列表。 
//
// 在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。 
//
// 示例 1： 
//
//输入：intervals = [[1,3],[6,9]], newInterval = [2,5]
//输出：[[1,5],[6,9]]
// 
// 示例 2：
//
//输入：intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
//输出：[[1,2],[3,10],[12,16]]
//解释：这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10]重叠。
//
// 示例 3： 
//
//输入：intervals = [], newInterval = [5,7]
//输出：[[5,7]]
//
// 示例 4： 
//
//输入：intervals = [[1,5]], newInterval = [2,3]
//输出：[[1,5]]
//
// 示例 5： 
//
//输入：intervals = [[1,5]], newInterval = [2,7]
//输出：[[1,7]]
// 
// 提示：
//
// 0 <= intervals.length <= 104 
// intervals[i].length == 2 
// 0 <= intervals[i][0] <= intervals[i][1] <= 105 
// intervals 根据 intervals[i][0] 按 升序 排列 
// newInterval.length == 2 
// 0 <= newInterval[0] <= newInterval[1] <= 105 


//leetcode submit region begin(Prohibit modification and deletion)
func insert(intervals [][]int, newInterval []int) [][]int {
	ans := make([][]int, 0)
	if len(intervals) < 1 {
		ans = append(ans, newInterval)
		return ans
	}
	inserted := false
	for _, interval := range intervals {
		if !inserted {
			if newInterval[0] < interval[0] && newInterval[1] < interval[0] {
				ans = append(ans, newInterval)
				ans = append(ans, interval)
				inserted = true
			} else if canMerge(interval, newInterval) {
				m := mergeTwooo(interval, newInterval)
				ans = append(ans, m)
				inserted = true
			} else {
				ans = append(ans, interval)
			}
		} else {
			if ans[len(ans)-1][1] >= interval[0] {
				m := mergeTwooo(ans[len(ans)-1], interval)
				ans[len(ans)-1] = m
			} else {
				ans = append(ans, interval)
			}
		}
	}
	if !inserted {
		ans = append(ans, newInterval)
	}
	return ans
}

func canMerge(i1 []int, i2 []int) bool {
	if i1[0] < i2[0] {
		return i1[1] >= i2[0]
	} else {
		return i2[1] >= i1[0]
	}
}

func mergeTwooo(i1 []int, i2 []int) []int {
	return []int{min(i1[0], i2[0]), max(i1[1], i2[1])}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}


//leetcode submit region end(Prohibit modification and deletion)
