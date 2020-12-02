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
	这一题是第 56 题的加强版。给出多个没有重叠的区间，然后再给一个区间，要求把如果有重叠的区间进行合并。

 */
func insert(intervals [][]int, newInterval []int) [][]int {
	ans := make([][]int, 0)
	for k, interval := range intervals {
		if interval[0] < newInterval[0] && interval[1] < newInterval[0] {
			ans = append(ans, interval)
			continue
		}
		if !needMerge(interval, newInterval) {
			ans = append(ans, newInterval)
			ans = append(ans, intervals[k:]...)
			return ans
		} else {
			merged := merge(interval, newInterval)
			ans = append(ans, merged)
			for i := k+1; i < len(intervals); i++ {
				if needMerge(ans[len(ans)-1], intervals[i]) {
					merged := merge(ans[len(ans)-1], intervals[i])
					ans[len(ans)-1] = merged
				} else {
					ans = append(ans, intervals[i:]...)
					return ans
				}
			}
			return ans
		}
	}
	ans = append(ans, newInterval)
	return ans
}

func needMerge(interval1 []int, interval2 []int) bool {
	return interval1[0] <= interval2[0] && interval1[1] >= interval2[0] || interval1[0] >= interval2[0] && interval1[0] <= interval2[1]
}

func merge(interval1 []int, interval2 []int) []int {
	return []int{min(interval1[0], interval2[0]), max(interval1[1], interval2[1])}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}