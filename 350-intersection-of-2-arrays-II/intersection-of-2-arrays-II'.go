package _50_intersection_of_2_arrays_II

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	ans := make([]int, 0)
	numMap := make(map[int]int)
	for _, v := range nums1 {
		numMap[v]++
	}
	for _, v := range nums2 {
		if count, ok := numMap[v]; ok && count > 0 {
			ans = append(ans, v)
			numMap[v] = count - 1
		}
	}
	return ans
}

func intersect(nums1 []int, nums2 []int) []int {
	ans := make([]int, 0)
	sort.Ints(nums1)
	sort.Ints(nums2)
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			ans = append(ans, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return ans
}
