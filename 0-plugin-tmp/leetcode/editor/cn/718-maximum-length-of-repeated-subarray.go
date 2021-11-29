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
	m, n := len(nums1), len(nums2)
	if m < n {
		return moveWindow(nums1, nums2)
	} else {
		return moveWindow(nums2, nums1)
	}
}

func moveWindow(nums1, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	ans := 0
	start1 := m - 1
	end2 := 0
	for start1 >= 0 {
		ans = max(ans, getMaxLength(nums1[start1:], nums2[:end2+1]))
		start1--
		end2++
	}
	start2 := 1
	for end2 <= n-1 {
		ans = max(ans, getMaxLength(nums1, nums2[start2:end2+1]))
		end2++
		start2++
	}
	end1 := m - 2
	for end1 >= 0 {
		ans = max(ans, getMaxLength(nums1[:end1+1], nums2[start2:]))
		start2++
		end1--
	}
	return ans
}

func getMaxLength(nums1 []int, nums2 []int) int {
	ans := 0
	count := 0
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); i, j = i+1, j+1 {
		if nums1[i] == nums2[j] {
			count++
		} else {
			count = 0
		}
		ans = max(ans, count)
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
