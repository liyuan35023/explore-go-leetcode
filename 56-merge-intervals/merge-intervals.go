package _6_merge_intervals

import (
	"sort"
)

/*
	给出一个区间的集合，请合并所有重叠的区间。

	示例 1:
	输入: intervals = [[1,3],[2,6],[8,10],[15,18]]
	输出: [[1,6],[8,10],[15,18]]
	解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
	示例2:

	输入: intervals = [[1,4],[4,5]]
	输出: [[1,5]]
	解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
	注意：输入类型已于2019年4月15日更改。 请重置默认代码定义以获取新方法签名。

	提示：
	intervals[i][0] <= intervals[i][1]

 */

type Intervals [][]int

func (i Intervals) Len() int {
	return len(i)
}

func (i Intervals) Less(m, n int) bool {
	if i[m][0] < i[n][0] {
		return true
	}
	return false
}

func (i Intervals) Swap(m, n int) {
	i[m], i[n] = i[n], i[m]
}

func merge(intervals [][]int) [][]int {
	if len(intervals) < 1 {
		return nil
	}
	// 排序
	sort.Sort(Intervals(intervals))

	ans := make([][]int, 0)
	preInterval := intervals[0]
	for i := 1; i < len(intervals); i++ {
		tmp := mergeTwo(preInterval, intervals[i])
		if len(tmp) == 1 {
			preInterval = tmp[0]
		} else {
			preInterval = intervals[i]
			ans = append(ans, tmp[0])
		}
	}
	ans = append(ans, preInterval)
	return ans
}

func mergeTwo(interval1 []int, interval2[]int) [][]int {
	if interval1[1] >= interval2[0] {
		// need merge
		return [][]int{[]int{interval1[0], max(interval1[1], interval2[1])}}
	}
	return [][]int{interval1, interval2}
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
