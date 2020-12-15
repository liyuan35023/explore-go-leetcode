package _8_merge_sort_array

/*
	Example:

	Input:
	nums1 = [1,2,3,0,0,0], m = 3
	nums2 = [2,5,6],       n = 3

	Output: [1,2,2,3,5,6]
	Constraints:

	-10^9 <= nums1[i], nums2[i] <= 10^9
	nums1.length == m + n
	nums2.length == n
	题目大意 #
	合并两个已经有序的数组，结果放在第一个数组中，第一个数组假设空间足够大。要求算法时间复杂度足够低。

 */

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n < 1 {
		return
	}
	totalLen := m + n

	// 双指针，从后向前
	i, j := m-1, n-1
	curIdx := totalLen - 1
	for ; i >= 0 && j >= 0; curIdx-- {
		if nums1[i] < nums2[j] {
			nums1[curIdx] = nums2[j]
			j--
		} else {
			nums1[curIdx] = nums1[i]
			i--
		}
	}
	for j >= 0 {
		nums1[curIdx] = nums2[j]
		curIdx--
		j--
	}
	return
}
