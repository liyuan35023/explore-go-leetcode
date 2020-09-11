package _8_four_sum

import "sort"

/*
	Given an array nums of n integers and an integer target, are there elements a, b, c, and d in nums such that a + b + c + d = target?
	Find all unique quadruplets in the array which gives the sum of target.

	Note:

	The solution set must not contain duplicate quadruplets.
 */

func FourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return [][]int{}
	}

	sort.Ints(nums)
	length := len(nums)
	ret := make([][]int, 0)

	for first := 0; first < length-3; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		for second := first+1; second < length-2; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			tmp := target - nums[first] - nums[second]
			i, j := second+1, length-1
			for i < j {
				sum := nums[i] + nums[j]
				if sum == tmp {
					ret = append(ret, []int{nums[first], nums[second], nums[i], nums[j]})
					for i < j && nums[i] == nums[i+1] && nums[j] == nums[j-1] {
						j--
						i++
					}
					j--
					i++
				} else if sum > tmp {
					j--
				} else {
					i++
				}
			}
		}
	}
	return ret
}