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
func findMedianSortedArraysFinal(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	if totalLen % 2 == 0 {
		return (float64(findKthLargestNumInTwoArrays(nums1, nums2, totalLen/2)) + float64(findKthLargestNumInTwoArrays(nums1, nums2, totalLen/2 + 1))) / 2
	} else {
		return float64(findKthLargestNumInTwoArrays(nums1, nums2, totalLen / 2 + 1))
	}
}

func findKthLargestNumInTwoArrays(nums1 []int, nums2 []int, k int) int {
	length1, length2 := len(nums1), len(nums2)
	startIdx1, startIdx2 := 0, 0
	for k > 1 && startIdx1 < length1 && startIdx2 < length2 {
		compareIdx1, compareIdx2 := 0, 0
		if startIdx1 + k / 2 - 1 < length1 {
			compareIdx1 = startIdx1+k/2-1
		} else {
			compareIdx1 = length1-1
		}
		if startIdx2 + k / 2 - 1 < length2 {
			compareIdx2 = startIdx2+k/2-1
		} else {
			compareIdx2 = length2-1
		}
		if nums1[compareIdx1] < nums2[compareIdx2] {
			k -= compareIdx1 - startIdx1 + 1
			startIdx1 = compareIdx1 + 1
		} else {
			k -= compareIdx2 - startIdx2 + 1
			startIdx2 = compareIdx2 + 1
		}
	}
	if startIdx1 >= length1 {
		return nums2[startIdx2+k-1]
	}
	if startIdx2 >= length2 {
		return nums1[startIdx1+k-1]
	}
	if nums1[startIdx1] < nums2[startIdx2] {
		return nums1[startIdx1]
	} else {
		return nums2[startIdx2]
	}
}


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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func findMedianSortedArraysIII(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	if totalLen % 2 == 0 {
		return (float64(findMinK(nums1, nums2, totalLen / 2)) + float64(findMinK(nums1, nums2, totalLen/2 + 1))) / 2
	} else {
		return float64(findMinK(nums1, nums2, totalLen / 2 + 1))
	}

}

func findMinK(nums1 []int, nums2 []int, k int) int {
	m, n := len(nums1), len(nums2)
	i, j := 0, 0
	for k > 1 && i < len(nums1) && j < len(nums2) {
		tmp := k / 2
		idx1, idx2 := 0, 0
		if i + tmp - 1 > m - 1 {
			idx1 = m - 1
		} else {
			idx1 = i + tmp - 1
		}
		if j + tmp - 1 > n - 1 {
			idx2 = n - 1
		} else {
			idx2 = j + tmp - 1
		}
		if nums1[idx1] < nums2[idx2] {
			k = k - (idx1 - i + 1)
			i = idx1 + 1
		} else {
			k = k - (idx2 - j + 1)
			j = idx2 + 1
		}
	}
	if i == m {
		return nums2[j+k-1]
	}
	if j == n {
		return nums1[i+k-1]
	}
	return min(nums1[i], nums2[j])
}