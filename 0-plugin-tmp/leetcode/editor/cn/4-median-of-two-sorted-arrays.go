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
	m, n := len(nums1), len(nums2)
	total := m + n
	if total % 2 == 0 {
		return (float64(findKthLargestNum(nums1, nums2, total/2) + findKthLargestNum(nums1, nums2, total/2 + 1))) / 2
	} else {
		return float64(findKthLargestNum(nums1, nums2, total/2 + 1))
	}
}

func findKthLargestNum(nums1, nums2 []int, k int) int {
	m, n := len(nums1), len(nums2)
	i, j := -1, -1
	for k > 1 && i < m - 1 && j < n - 1 {
		x, y := i + k / 2, j + k / 2
		if x >= m {
			x = m - 1
		}
		if y >= n {
			y = n - 1
		}
		if nums1[x] <= nums2[y] {
			k = k - (x - i)
			i = x
		} else {
			k = k - (y - j)
			j = y
		}
	}
	if i >= m - 1 {
		return nums2[j+k]
	}
	if j >= n - 1 {
		return nums1[i+k]
	}
	return min(nums1[i+1], nums2[j+1])

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y

}

//leetcode submit region end(Prohibit modification and deletion)
