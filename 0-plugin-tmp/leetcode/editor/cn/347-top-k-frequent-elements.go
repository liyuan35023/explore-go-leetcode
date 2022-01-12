package cn

import "math/rand"

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
	countMap := make(map[int]int)
	for _, v := range nums {
		countMap[v]++
	}
	seq := make([]int, 0)
	for value := range countMap {
		seq = append(seq, value)
	}
	qsort(seq, 0, len(seq)-1, k, countMap)
	return seq[:k]
}

func qsort(nums []int, left, right int, k int, countMap map[int]int) {
	for left < right {
		mid := partition(nums, left, right, countMap)
		if mid == k - 1 {
			return
		} else if mid > k - 1 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
}

func partition(nums []int, left, right int, countMap map[int]int) int {
	randomIdx := rand.Intn(right - left + 1) + left
	nums[left], nums[randomIdx] = nums[randomIdx], nums[left]
	pivot := nums[left]
	for left < right {
		for left < right && countMap[nums[right]] <= countMap[pivot] {
			right--
		}
		nums[left] = nums[right]
		for left < right && countMap[nums[left]] >= countMap[pivot] {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = pivot
	return left
}


//leetcode submit region end(Prohibit modification and deletion)
