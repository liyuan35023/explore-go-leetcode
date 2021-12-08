package cn

//以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返
//回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
//
// 示例 1：
// 
//输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
//输出：[[1,6],[8,10],[15,18]]
//解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
//
// 示例 2： 
//
//输入：intervals = [[1,4],[4,5]]
//输出：[[1,5]]
//解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
// 
//
// 提示： 
//
// 1 <= intervals.length <= 104 
// intervals[i].length == 2 
// 0 <= starti <= endi <= 104


//leetcode submit region begin(Prohibit modification and deletion)
func merge(intervals [][]int) [][]int {
}

func canMerge(i1, i2 []int) bool {
	return i1[1] >= i2[0]
}
func mergeTwoInterval(i1, i2 []int) []int {
	return []int{i1[0], max(i1[1], i2[1])}
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y

}

func sortIntervals(intervals [][]int, l, r int) {
	// 三路快排
	if l >= r {
		return
	}
	p0, p2 := partition3(intervals, l, r)
	sortIntervals(intervals, l, p0 - 1)
	sortIntervals(intervals, p2 + 1, r)
}

func partition3(intervals [][]int, left, right int) (int, int) {
	pivot := intervals[left][0]
	p0 := left
	i := left + 1
	p2 := right
	for i <= p2 {
		if intervals[i][0] == pivot {
			i++
		} else if intervals[i][0] < pivot {
			intervals[i], intervals[p0] = intervals[p0], intervals[i]
			p0++
			i++
		} else {
			intervals[i], intervals[p2] = intervals[p2], intervals[i]
			p2--
		}
	}
	return p0, p2
}


//leetcode submit region end(Prohibit modification and deletion)
