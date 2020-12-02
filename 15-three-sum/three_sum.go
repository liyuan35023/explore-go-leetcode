package _5_three_sum

import "sort"

/*
	Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0?
	Find all unique triplets in the array which gives the sum of zero.
	The solution set must not contain duplicate triplets.

	Given array nums = [-1, 0, 1, 2, -1, -4],

	A solution set is:
	[
	  [-1, 0, 1],
	  [-1, -1, 2]
	]
 */

func ThreeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	var ret [][]int
	sort.Ints(nums)

	for head := 0; head < len(nums); head++ {
		if head > 0 && nums[head] == nums[head-1] {
			continue
		}
		target := -nums[head]
		tail := len(nums)-1
		for mid := head+1; mid < len(nums); mid++ {
			if mid > head+1 && nums[mid] == nums[mid-1] {
				continue
			}
			for mid < tail && nums[mid] + nums[tail] > target {
				tail--
			}
			if mid == tail {
				 break
			}
			if nums[mid] + nums[tail] == target {
				ret = append(ret, []int{nums[head], nums[mid], nums[tail]})
			}
		}
	}
	return ret
}

func ThreeSumII(nums []int) [][]int {
	ans := make([][]int, 0)
	if len(nums) < 3 {
		return ans
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				left++
			} else if sum > target {
				right--
			} else if sum < target {
				left++
			}
		}
	}
	return ans
}
