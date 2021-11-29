package cn
//给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。 
//
// 示例 1： 
//
//输入：nums1 = [1,3], nums2 = [2]
//输出：2.00000
//解释：合并数组 = [1,2,3] ，中位数 2
//
// 示例 2： 
//
//输入：nums1 = [1,2], nums2 = [3,4]
//输出：2.50000
//解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
//
// 示例 3： 
//
//输入：nums1 = [0,0], nums2 = [0,0]
//输出：0.00000
//
// 示例 4： 
//
//输入：nums1 = [], nums2 = [1]
//输出：1.00000
//
// 示例 5： 
//
//输入：nums1 = [2], nums2 = []
//输出：2.00000
// 
// nums1.length == m
// nums2.length == n 
// 0 <= m <= 1000 
// 0 <= n <= 1000 
// 1 <= m + n <= 2000 
// -10⁶ <= nums1[i], nums2[i] <= 10⁶ 
// 
// 进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？


//leetcode submit region begin(Prohibit modification and deletion)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	if total % 2 == 0 {
		return (float64(findKthLargestNum(nums1, nums2, total/2)) + float64(findKthLargestNum(nums1, nums2, total/2 + 1))) / 2.0
	} else {
		return float64(findKthLargestNum(nums1, nums2, total / 2 + 1))
	}
}

func findKthLargestNum(nums1 []int, nums2 []int, k int) int {
	m, n := len(nums1), len(nums2)
	l1, l2 := 0, 0
	for k > 1 && l1 < m && l2 < n {
		count := k / 2
		var idx1, idx2 int
		if l1 + count - 1 < m {
			idx1 = l1 + count - 1
		} else {
			idx1 = m - 1
		}
		if l2 + count - 1 < n {
			idx2 = l2 + count - 1
		} else {
			idx2 = n - 1
		}
		if nums1[idx1] < nums2[idx2] {
			k = k - (idx1 - l1 + 1)
			l1 = idx1 + 1
		} else {
			k = k - (idx2 - l2 + 1)
			l2 = idx2 + 1
		}
	}
	if l1 >= m {
		return nums2[l2+k-1]
	}
	if l2 >= n {
		return nums1[l1+k-1]
	}
	return min(nums1[l1], nums2[l2])
}



func min(x, y int) int {
	if x < y {
		return x
	}
	return y

}

//leetcode submit region end(Prohibit modification and deletion)
