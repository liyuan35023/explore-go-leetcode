package cn
//给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。 
//
// 示例 1:
//
//输入: nums = [1,1,1,2,2,3], k = 2
//输出: [1,2]
//
// 示例 2: 
//
//输入: nums = [1], k = 1
//输出: [1] 
//
// 提示： 
//
// 1 <= nums.length <= 105
// k 的取值范围是 [1, 数组中不相同的元素的个数] 
// 题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的 
// 
// 进阶：你所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小。


//leetcode submit region begin(Prohibit modification and deletion)
func topKFrequent(nums []int, k int) []int {
	heap := make([]int, 0)
	// 堆排序
	// 小顶堆， 删除堆顶的元素，直至剩下k个
	numMap := make(map[int]int)
	for _, v := range nums {
		numMap[v]++
	}

	for num, count := range numMap {
		if len(heap) < k {
			heap = append(heap, num)
			headUpAdjust(heap, numMap, len(heap)-1)
		} else if numMap[heap[0]] < count {
			heap[0] = num
			headDownAdjust(heap, numMap, 0, k)
		}
	}
	return heap
}

func constructHeap(nums []int, numMap map[int]int, length int) {
	lastNotLeaf := length / 2 - 1
	for i := lastNotLeaf; i >= 0; i-- {
		headDownAdjust(nums, numMap, i, length)
	}
}

func headDownAdjust(nums []int, numMap map[int]int, parent int, length int) {
	for parent < length / 2 {
		minChildIdx := getMinChild(nums, numMap, parent, length)
		if numMap[nums[parent]] > numMap[nums[minChildIdx]] {
			nums[parent], nums[minChildIdx] = nums[minChildIdx], nums[parent]
			parent = minChildIdx
		} else {
			break
		}
	}
}

func getMinChild(nums []int, numMap map[int]int, parent int, length int) int {
	child1 := parent * 2 + 1
	child2 := parent * 2 + 2
	if child2 >= length {
		return child1
	} else {
		if numMap[nums[child1]] < numMap[nums[child2]] {
			return child1
		} else {
			return child2
		}
	}
}

func headUpAdjust(nums []int, numMap map[int]int, idx int) {
	for idx > 0 {
		parent := (idx - 1) / 2
		if parent >= 0 && numMap[nums[parent]] > numMap[nums[idx]] {
			nums[parent], nums[idx] = nums[parent], nums[idx]
			idx = parent
		} else {
			break
		}
	}
}




//leetcode submit region end(Prohibit modification and deletion)
