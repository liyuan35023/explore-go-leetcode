package _6_three_sum_closest

import (
	"math"
	"sort"
)

/*
	Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target.
	Return the sum of the three integers. You may assume that each input would have exactly one solution.

	Example:

	Given array nums = [-1, 2, 1, -4], and target = 1.

	The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
 */

func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	length := len(nums)
	closet := int(math.MinInt32)

	update := func(cur int) {
		if abs(cur - target) < abs(closet - target) {
			closet = cur
		}
	}

	for head := 0; head < length-2; head++ {
		if head > 0 && nums[head] == nums[head-1] {
			continue
		}
		mid, tail := head+1, length-1
		for mid < tail {
			sum := nums[head] + nums[mid] + nums[tail]
			if sum == target {
				return sum
			}
			update(sum)
			if sum > target {
				for mid < tail && nums[tail] == nums[tail-1] {
					tail--
				}
				tail--
			} else {
				for mid < tail && nums[mid] == nums[mid+1] {
					mid++
				}
				mid++
			}
		}
	}
	return closet
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func threeSumClosest(nums []int, target int) int {
	// 排序
	sort.Ints(nums)
	length := len(nums)

	ret := 100000
	update := func(cur int) {
		if abs(cur-target) < abs(ret-target) {
			ret = cur
		}
	}

	// 双指针
	for head := 0; head < length-2; head++ {
		if head > 0 && nums[head] == nums[head-1] {
			continue
		}
		i, j := head+1, length-1
		for i < j {
			sum := nums[head] + nums[i] + nums[j]
			if sum == target {
				return sum
			}
			update(sum)
			if sum > target {
				for i < j && nums[j-1] == nums[j] {
					j--
				}
				j--
			} else {
				for i < j && nums[i+1] == nums[i] {
					i++
				}
				i++
			}
		}
	}
	return ret
}
