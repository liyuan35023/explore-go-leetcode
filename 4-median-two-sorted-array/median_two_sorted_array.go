package __median_two_sorted_array

/**
 * There are two sorted arrays nums1 and nums2 of size m and n respectively.
 * Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
 * You may assume nums1 and nums2 cannot be both empty.
 *
 * Example 1:
 * nums1 = [1, 3]
 * nums2 = [2]
 * The median is 2.0
 *
 * Example 2:
 * nums1 = [1, 2]
 * nums2 = [3, 4]
 * The median is (2 + 3)/2 = 2.5

 * 给定两个大小为 m 和 n 的正序（从小到大）数组nums1 和nums2。请你找出并返回这两个正序数组的中位数。
 * 进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？
 *
 *
 */

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	sorted := []int{}
	if m == 0 || n == 0 {
		sorted = append(sorted, nums1...)
		sorted = append(sorted, nums2...)
		sorted = sorted[:(m+n)/2 + 1]
	}

	i, j := 0, 0
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			sorted = append(sorted, nums1[i])
			i++
		} else {
			sorted = append(sorted, nums2[j])
			j++
		}
		if len(sorted) == (m+n) / 2 + 1 {
			break
		}
		// i++ or j++
	}
	if len(sorted) != (m+n) / 2 + 1 {
		targetNum := (m+n) / 2 + 1 - len(sorted)
		if i == m {
			sorted = append(sorted, nums2[j:j+targetNum]...)
		} else if j == n {
			sorted = append(sorted, nums1[i:i+targetNum]...)
		}
	}

	if (m+n) % 2 == 0 {
		return (float64(sorted[len(sorted)-2]) + float64(sorted[len(sorted)-1])) / 2.0
	} else {
		return float64(sorted[len(sorted)-1])
	}
}