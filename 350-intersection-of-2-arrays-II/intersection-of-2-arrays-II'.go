package _50_intersection_of_2_arrays_II

import "sort"

/*
	给定两个数组，编写一个函数来计算它们的交集。

示例 1：

输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2,2]
示例 2:

输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[4,9]


说明：

输出结果中每个元素出现的次数，应与元素在两个数组中出现次数的最小值一致。
我们可以不考虑输出结果的顺序。
进阶：

如果给定的数组已经排好序呢？你将如何优化你的算法？
如果nums1的大小比nums2小很多，哪种方法更优？
如果nums2的元素存储在磁盘上，内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？

 */

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

func intersect2(nums1 []int, nums2 []int) []int {
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
