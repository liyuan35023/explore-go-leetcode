package _00_longest_increasing_subSequence

/*
	Example:

	Input: [10,9,2,5,3,7,101,18]
	Output: 4
	Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
	Note:

	There may be more than one LIS combination, it is only necessary for you to return the length.
	Your algorithm should run in O(n^2) complexity.
	Follow up: Could you improve it to O(n log n) time complexity?

	题目大意 #
	给定一个无序的整数数组，找到其中最长上升子序列的长度。

 */

func lengthOfLIS(nums []int) int {
	// dp O(n^2)
	// dp[i] 表示以 index为i的数结尾的最长子序列的长度
	n := len(nums)
	if n < 2 {
		return n
	}
	ans := 1
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		max := 0
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				max = maxNum(max, dp[j])
			}
		}
		dp[i] = max + 1
		ans = maxNum(ans, dp[i])
	}
	return ans
}

func maxNum(x, y int) int {
	if x > y {
		return x
	}
	return y
}


func lengthOfLISII(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	//ans := 1
	// length[i] 表示以i+1长的上升子序列的末尾数字中的最小的数
	length := make([]int, 1)
	length[0] = nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > length[len(length)-1] {
			length = append(length, nums[i])
		} else {
			idx := binarySearch(length, nums[i])
			length[idx] = nums[i]
		}
	}
	return len(length)
}

func binarySearch(nums []int, target int) int {
	// 找到target，如果存在，返回所在index， 如果不存在，返回该插入的index
	i, j := 0, len(nums) - 1
	for i <= j {
		mid := (i+j) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			j = mid - 1
		} else if nums[mid] < target {
			i = mid + 1
		}
	}
	return i
}

func findFirstBiggerIdx(nums []int, target int) int {
	left, right := 0, len(nums) - 1
	for left <= right {
		mid := left + (right - left) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}



















