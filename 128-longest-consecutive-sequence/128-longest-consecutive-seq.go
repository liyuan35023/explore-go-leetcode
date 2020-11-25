package _28_longest_consecutive_sequence

/*
	Example:

	Input: [100, 4, 200, 1, 3, 2]
	Output: 4
	Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
	题目大意 #
	给定一个未排序的整数数组，找出最长连续序列的长度。要求算法的时间复杂度为 O(n)。

 */

func longestConsecutive(nums []int) int {
	// 并查集 借助map
	if len(nums) < 2 {
		return len(nums)
	}
	numMap := make(map[int]int)
	seqList := make([][]int, 0)
	ans := 0
	for _, v := range nums {
		if _, ok := numMap[v]; ok {
			continue
		}
		idxLess, okLess := numMap[v-1]
		idxBigger, okMore := numMap[v+1]
		if okLess && okMore {
			// union two seq
			for _, bv := range seqList[idxBigger] {
				numMap[bv] = idxLess
				seqList[idxLess] = append(seqList[idxLess], bv)
			}
			seqList[idxLess] = append(seqList[idxLess], v)
			numMap[v] = idxLess
			ans = max(ans, len(seqList[idxLess]))
		} else if okLess && !okMore {
			seqList[idxLess] = append(seqList[idxLess], v)
			numMap[v] = idxLess
			ans = max(ans, len(seqList[idxLess]))
		} else if !okLess && okMore {
			seqList[idxBigger] = append(seqList[idxBigger], v)
			numMap[v] = idxBigger
			ans = max(ans, len(seqList[idxBigger]))
		} else {
			seqList = append(seqList, []int{v})
			numMap[v] = len(seqList) - 1
			ans = max(ans, len(seqList[numMap[v]]))
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func longestConsecutiveII(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	seqMap := make(map[int][]int)
	ans := 0
	for _, v := range nums {
		if _, ok := seqMap[v]; ok {
			continue
		}
		seqLess, okLess := seqMap[v-1]
		seqMore, okMore := seqMap[v+1]
		if okLess && okMore {
			seqMap[seqLess[0]][1] = seqMore[1]
			seqMap[seqMore[1]][0] = seqLess[0]
			seqMap[v] = []int{seqLess[0], seqMore[1]}
			ans = max(ans, seqMore[1] - seqLess[0] + 1)
		} else if okLess && !okMore {
			seqMap[v] = []int{seqLess[0], v}
			seqMap[seqLess[0]] = []int{seqLess[0], v}
			ans = max(ans, v - seqLess[0] + 1)
		} else if !okLess && okMore {
			seqMap[v] = []int{v, seqMore[1]}
			seqMap[seqMore[1]] = []int{v, seqMore[1]}
			ans = max(ans, seqMore[1] - v + 1)
		} else {
			seqMap[v] = []int{v, v}
			ans = max(ans, 1)
		}
	}
	return ans
}

func longestConsecutiveIII(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	ans := 0
	numMap := make(map[int]int)
	for _, v := range nums {
		if numMap[v] == 0 {
			left := numMap[v-1]
			right := numMap[v+1]
			length := left + right + 1
			numMap[v] = length
			numMap[v-left] = length
			numMap[v+right] = length
			ans = max(ans, length)
		}
	}
	return ans
}
