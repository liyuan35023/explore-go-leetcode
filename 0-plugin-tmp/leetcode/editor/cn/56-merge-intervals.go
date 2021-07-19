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
type interval [][]int

func (i interval) Len() int {
	return len(i)
}

func (i interval) Less(x, y int) bool {
	return i[x][0] < i[y][0]
}

func (i interval) Swap(x, y int) {
	i[x], i[y] = i[y], i[x]
}

func merge(intervals [][]int) [][]int {
	sort.Sort(interval(intervals))
	ans := make([][]int, 0)
	ans = append(ans, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if ans[len(ans)-1][1] >= intervals[i][0] {
			ans[len(ans)-1] = merge2Interval(ans[len(ans)-1], intervals[i])

		} else {
			ans = append(ans, intervals[i])
		}
	}
	return ans
}

func merge2Interval(i1 []int, i2 []int) []int {
	return []int{i1[0], max(i1[1], i2[1])}
}



//leetcode submit region end(Prohibit modification and deletion)
