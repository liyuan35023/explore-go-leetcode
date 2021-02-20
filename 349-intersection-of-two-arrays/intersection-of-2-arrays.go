package _49_intersection_of_two_arrays

import "sort"

/*
	Example 1:


Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]

Example 2:


Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]

Note:

Each element in the result must be unique.
The result can be in any order.
题目大意 #
找到两个数组的交集元素，如果交集元素同一个数字出现了多次，只输出一次。

 */

func intersection(nums1 []int, nums2 []int) []int {
	ans := make([]int, 0)
	num1Map := make(map[int]struct{})
	num2Map := make(map[int]struct{})
	for _, v := range nums1 {
		num1Map[v] = struct{}{}
	}

	for _, v := range nums2 {
		num2Map[v] = struct{}{}
	}

	for k, _ := range num1Map {
		if _, ok := num2Map[k]; ok {
			ans = append(ans, k)
		}

	}
	return ans
}

func intersectionII(nums1 []int, nums2 []int) []int {
	ans := make([]int, 0)
	sort.Ints(nums1)
	sort.Ints(nums2)
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			ans = append(ans, nums2[j])
			i++
			for i < len(nums1) && nums1[i] == nums1[i-1] {
				i++
			}
		} else {
			if nums1[i] < nums2[j] {
				i++
			} else {
				j++
			}
		}
	}
	return ans
}






