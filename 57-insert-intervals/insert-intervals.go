package _7_insert_intervals

/*
	Example 1:

	Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
	Output: [[1,5],[6,9]]

	Example 2:


	Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
	Output: [[1,2],[3,10],[12,16]]
	Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].

	题目大意 #
	给出一个无重叠，按照区间起始点排序的区间列表
	这一题是第 56 题的加强版。给出多个没有重叠的区间，然后再给一个区间，要求把如果有重叠的区间进行合并。

 */
func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) < 1 {
		return [][]int{newInterval}
	}
	ans := make([][]int, 0)
	mergedNew := false
	for i := 0; i < len(intervals); i++ {
		if intervals[i][1] < newInterval[0] {
			ans = append(ans, intervals[i])
			continue
		}
		if intervals[i][0] > newInterval[1] {
			if !mergedNew {
				ans = append(ans, newInterval)
				mergedNew = true
			}
			ans = append(ans, intervals[i])
			continue
		}
		if needMerge(intervals[i], newInterval) {
			if !mergedNew {
				ans = append(ans, mergeTwo(intervals[i], newInterval))
				mergedNew = true
			} else {
				ans[len(ans)-1] = mergeTwo(ans[len(ans)-1], intervals[i])
			}
		}
	}
	if !mergedNew {
		ans = append(ans, newInterval)
	}
	return ans
}

func mergeTwo(i1 []int, i2 []int) []int {
	return []int{min(i1[0], i2[0]), max(i1[1], i2[1])}
}

func needMerge(i1 []int, i2 []int) bool {
	return i1[0] <= i2[0] && i1[1] >= i2[0] || i2[0] <= i1[0] && i2[1] >= i1[0]
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
