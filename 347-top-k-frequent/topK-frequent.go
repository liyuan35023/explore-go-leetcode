package _47_top_k_frequent

/*
	Example 1:


	Input: nums = [1,1,1,2,2,3], k = 2
	Output: [1,2]


	Example 2:
	Input: nums = [1], k = 1
	Output: [1]

	Note:
	You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
	Your algorithm’s time complexity must be better than O(n log n), where n is the array’s size.
	题目大意 #
	给一个非空的数组，输出前 K 个频率最高的元素。
 */

func topKFrequent(nums []int, k int) []int {
	ans := make([]int, 0)                   // k num
	numMap := make(map[int]int)             // num -> count
	for _, v := range nums {
		if count, ok := numMap[v]; ok {
			numMap[v] = count+1
		} else {
			numMap[v] = 1
		}
	}
	for num, count := range numMap {
		if len(ans) < k {
			ans = append(ans, num)
			upAdjust(ans, numMap, len(ans))
		} else {
			if numMap[ans[0]] < count {
				ans[0] = num
				downAdjust(ans, numMap, 0, len(ans))
			}
		}
	}
	return ans
}

func upAdjust(nums []int, numMap map[int]int, length int) {
	// 构建小顶堆
	// 插入操作
	idx := length - 1
	parent := (idx - 1) / 2
	for parent >= 0 && numMap[nums[parent]] > numMap[nums[idx]] {
		nums[parent], nums[idx] = nums[idx], nums[parent]
		idx = parent
		parent = (idx - 1) / 2
	}
}

func downAdjust(nums []int, numMap map[int]int, parent int, length int) {
	// 小顶堆
	// 用于替换堆顶后进行下沉
	for parent < length/2 {
		minChild := getMinChildIdx(nums, numMap, parent, length)
		if numMap[nums[parent]] > numMap[nums[minChild]] {
			nums[parent], nums[minChild] = nums[minChild], nums[parent]
			parent = minChild
		} else {
			break
		}
	}
}

func getMinChildIdx(nums []int, numMap map[int]int, parent int, length int) int {
	child1 := 2 * parent + 1
	child2 := 2 * parent + 2
	ans := child1
	if child2 < length && numMap[nums[child1]] > numMap[nums[child2]] {
		ans = child2
	}
	return ans
}

func topKFrequent22(nums []int, k int) []int {
	ans := make([]int, 0)
	countMap := make(map[int]int)
	for _, v := range nums {
		countMap[v]++
	}
	for value, count := range countMap {
		if len(ans) < k {
			ans = append(ans, value)
			upAdjust22(ans, countMap, len(ans))
		} else {
			if count > countMap[ans[0]] {
				ans[0] = value
				downAdjust22(ans, countMap, 0, len(ans)-1)
			}
		}
	}
	return ans
}

func upAdjust22(nums []int, countMap map[int]int, length int) {
	idx := length - 1
	for idx >= 1 {
		parent := (idx - 1) / 2
		if countMap[nums[parent]] <= countMap[nums[idx]] {
			break
		}
		nums[idx], nums[parent] = nums[parent], nums[idx]
		idx = parent
	}
}

func downAdjust22(nums []int, countMap map[int]int, start, end int) {
	idx := start
	for idx < (end - start + 1) / 2 {
		minChild := findMinChild(nums, countMap, idx, end-start+1)
		if countMap[nums[minChild]] >=  countMap[nums[idx]] {
			return
		}
		nums[idx], nums[minChild] = nums[minChild], nums[idx]
		idx = minChild
	}
}

func findMinChild(nums []int, countMap map[int]int, parent int, length int) int {
	child1 := 2 * parent + 1
	child2 := 2 * parent + 2
	if child2 > length - 1 {
		return child1
	}

	if countMap[nums[child1]] < countMap[nums[child2]] {
		return child1
	}
	return child2

}