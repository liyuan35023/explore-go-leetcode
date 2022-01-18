package cn
//给两个整数数组 A 和 B ，返回两个数组中公共的、长度最长的子数组的长度。 
//
// 示例： 
//
// 输入：
//A: [1,2,3,2,1]
//B: [3,2,1,4,7]
//输出：3
//解释：
//长度最长的公共子数组是 [3, 2, 1] 。
//
// 提示： 
//
// 1 <= len(A), len(B) <= 1000 
// 0 <= A[i], B[i] < 100 

//leetcode submit region begin(Prohibit modification and deletion)
func findLength(nums1 []int, nums2 []int) int {
	if len(nums1) < len(nums2) {
		return findLength(nums2, nums1)
	}
	m, n := len(nums1), len(nums2)

	ans := 0
	i, j := 0, n - 1
	for j > 0 {
		tmp := getCommonLen(nums1[:i+1], nums2[j:])
		ans = max(ans, tmp)
		i++
		j--
	}
	for i < m {
		tmp := getCommonLen(nums1[i-n+1:i+1], nums2)
		ans = max(ans, tmp)
		i++
	}
	i = i-n+1
	j = n - 1
	for j >= 0 {
		tmp := getCommonLen(nums1[i:], nums2[:j])
		ans = max(ans, tmp)
		i++
		j--
	}
	return ans
}

func getCommonLen(nums1 []int, nums2 []int) int {
	i, j := 0, 0
	ans := 0
	count := 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			count++
		} else {
			count = 0
		}
		ans = max(ans, count)
		i++
		j++
	}
	return ans
}


func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
//leetcode submit region end(Prohibit modification and deletion)
