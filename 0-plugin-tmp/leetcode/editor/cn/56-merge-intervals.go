package cn

import "sort"

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
type Interval [][]int

func (i Interval) Len() int {
	return len(i)
}

func (i Interval) Less(x, y int) bool {
	return i[x][0] < i[y][0]
}

func (i Interval) Swap(x, y int) {
	i[x], i[y] = i[y], i[x]
}

func merge(intervals [][]int) [][]int {
	sort.Sort(Interval(intervals))
	ans := make([][]int, 0)
	ans = append(ans, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if canMerge(ans[len(ans)-1], intervals[i]) {
			ans[len(ans)-1] = mergeTwo(ans[len(ans)-1], intervals[i])
		} else {
			ans = append(ans, intervals[i])
		}
	}
	return ans
}

func canMerge(i1, i2 []int) bool {
	return i1[1] >= i2[0]
}

func mergeTwo(i1, i2 []int) []int {
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
