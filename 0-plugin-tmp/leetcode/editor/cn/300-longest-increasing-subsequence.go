package cn


//给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
//
// 子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序
//列。 
//
// 示例 1： 
//
//输入：nums = [10,9,2,5,3,7,101,18]
//输出：4
//解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
//
// 示例 2： 
//
//输入：nums = [0,1,0,3,2,3]
//输出：4
//
// 示例 3： 
//
//输入：nums = [7,7,7,7,7,7,7]
//输出：1
//
// 提示： 
//
// 1 <= nums.length <= 2500 
// -104 <= nums[i] <= 104 
//
// 进阶： 
//
// 你可以设计时间复杂度为 O(n2) 的解决方案吗？ 
// 你能将算法的时间复杂度降低到 O(n log(n)) 吗? 

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLIS(nums []int) int {
	seq := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if len(seq) == 0 || nums[i] > seq[len(seq)-1] {
			seq = append(seq, nums[i])
		} else {
			idx := findFirstBigger(seq, nums[i])
			seq[idx] = nums[i]
		}
	}
	return len(seq)
}

func findFirstBigger(nums []int, target int) int {
	left, right := 0, len(nums) - 1
	for left <= right {
		mid := left + (right - left) / 2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}





//func lengthOfLIS(nums []int) int {
//	lastNum := make([]int, 0)
//	lastNum = append(lastNum, nums[0])
//	findIdx := func(nums []int, target int) int {
//		left, right := 0, len(nums) - 1
//		for left <= right {
//			mid := left + (right - left) / 2
//			if nums[mid] == target {
//				return mid
//			} else if nums[mid] > target {
//				//if mid == 0 || mid > 0 && nums[mid-1] < target {
//				//	return mid
//				//}
//				right = mid - 1
//			} else {
//				//if mid < len(nums) - 1 && nums[mid+1] > target {
//				//	return mid + 1
//				//}
//				left = mid + 1
//			}
//		}
//		return left
//	}
//	for i := 1; i < len(nums); i++ {
//		if nums[i] > lastNum[len(lastNum)-1] {
//			lastNum = append(lastNum, nums[i])
//		} else {
//			idx := findIdx(lastNum, nums[i])
//			lastNum[idx] = nums[i]
//		}
//	}
//	return len(lastNum)
//
//
//}
//func lengthOfLIS(nums []int) int {
//	seq := make([]int, 0)
//	seq = append(seq, nums[0])
//	for i := 1; i < len(nums); i++ {
//		if nums[i] < seq[0] {
//			seq[0] = nums[i]
//		} else if nums[i] > seq[len(seq)-1] {
//			seq = append(seq, nums[i])
//		} else {
//			SearchAndReplace(seq, nums[i])
//		}
//	}
//	return len(seq)
//}
//
//func SearchAndReplace(nums []int, target int) {
//	left, right := 0, len(nums) - 1
//	for left <= right {
//		mid := left + (right - left) / 2
//		if nums[mid] == target {
//			return
//		} else if nums[mid] > target {
//			right = mid - 1
//		} else {
//			left = mid + 1
//		}
//	}
//	nums[left] = target
//}

//leetcode submit region end(Prohibit modification and deletion)
