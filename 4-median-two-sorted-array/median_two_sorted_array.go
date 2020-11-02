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
 */

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	sorted := []int{}
	if m == 0 || n == 0 {
		sorted = append(sorted, nums1...)
		sorted = append(sorted, nums2...)
		sorted = sorted[:(m+n)/2+1]
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
		if len(sorted) == (m+n)/2+1 {
			break
		}
		// i++ or j++
	}
	if len(sorted) != (m+n)/2+1 {
		targetNum := (m+n)/2 + 1 - len(sorted)
		if i == m {
			sorted = append(sorted, nums2[j:j+targetNum]...)
		} else if j == n {
			sorted = append(sorted, nums1[i:i+targetNum]...)
		}
	}

	if (m+n)%2 == 0 {
		return (float64(sorted[len(sorted)-2]) + float64(sorted[len(sorted)-1])) / 2.0
	} else {
		return float64(sorted[len(sorted)-1])
	}
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	if totalLen == 1 {
		if len(nums1) != 0 {
			return float64(nums1[0])
		} else {
			return float64(nums2[0])
		}
	}
	if totalLen % 2 != 0 {
		return float64(findMaxKNum(nums1, nums2, totalLen/2 + 1))
	} else {
		return (float64(findMaxKNum(nums1, nums2, totalLen/2)) + float64(findMaxKNum(nums1, nums2, totalLen/2 + 1))) / 2
	}
}

func findMaxKNum(nums1 []int, nums2 []int, k int) int {
	m := len(nums1)
	n := len(nums2)
	startIdx1 := 0
	startIdx2 := 0
	for k > 1 && startIdx1 < m && startIdx2 < n {
		mid1 := -1
		mid2 := -1
		if startIdx1+(k/2-1) < m {
			mid1 = startIdx1 + k/2 - 1
		} else {
			mid1 = m - 1
		}
		if startIdx2+(k/2-1) < n {
			mid2 = startIdx2 + k/2 - 1
		} else {
			mid2 = n - 1
		}
		if nums1[mid1] < nums2[mid2] {
			k = k - (mid1 - startIdx1 + 1)
			startIdx1 = mid1 + 1
		} else {
			k = k - (mid2 - startIdx2 + 1)
			startIdx2 = mid2 + 1
		}
	}
	if startIdx1 == m {
		return nums2[startIdx2+k-1]
	}
	if startIdx2 == n {
		return nums1[startIdx1+k-1]
	}
	return min(nums1[startIdx1], nums2[startIdx2])
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
