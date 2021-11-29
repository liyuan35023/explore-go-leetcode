package cn

import "sort"

//给定两个数组，编写一个函数来计算它们的交集。
//
// 示例 1： 
//
// 输入：nums1 = [1,2,2,1], nums2 = [2,2]
//输出：[2]
// 
//
// 示例 2： 
//
// 输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出：[9,4] 
//
// 说明：
//
// 输出结果中的每个元素一定是唯一的。
// 我们可以不考虑输出结果的顺序。 
// 
// Related Topics 数组 哈希表 双指针 二分查找 排序
// 👍 387 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func intersection(nums1 []int, nums2 []int) []int {
	ans := make([]int, 0)
	sort.Ints(nums1)
	sort.Ints(nums2)
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			ans = append(ans, nums1[i])
			i, j = i+1, j+1
			for i < len(nums1) && nums1[i] == nums1[i-1] {
				i++
			}
			for j < len(nums2) && nums2[j] == nums2[j-1] {
				j++
			}
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
