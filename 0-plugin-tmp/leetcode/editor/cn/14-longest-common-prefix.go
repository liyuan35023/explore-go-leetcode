package cn

import "sort"

//编写一个函数来查找字符串数组中的最长公共前缀。
//
// 如果不存在公共前缀，返回空字符串 ""。 
//
// 
//
// 示例 1： 
//
// 
//输入：strs = ["flower","flow","flight"]
//输出："fl"
// 
//
// 示例 2： 
//
// 
//输入：strs = ["dog","racecar","car"]
//输出：""
//解释：输入不存在公共前缀。 
//
// 
//
// 提示： 
//
// 
// 0 <= strs.length <= 200 
// 0 <= strs[i].length <= 200 
// strs[i] 仅由小写英文字母组成 
// 
// Related Topics 字符串 
// 👍 1643 👎 0



//leetcode submit region begin(Prohibit modification and deletion)
type strSlice []string

func (s *strSlice) Len() int {
	return len(*s)
}

func (s *strSlice) Less(i, j int) bool {
	return (*s)[i] < (*s)[j]
}

func (s *strSlice) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func longestCommonPrefix(strs []string) string {
	s := strSlice(strs)
	sort.Sort(&s)
	sorted := []string(s)
	min, max := sorted[0], sorted[len(sorted)-1]
	right := 0
	for right < len(min) && right < len(min) && min[right] == max[right] {
		right++
	}
	return min[:right]
}
//leetcode submit region end(Prohibit modification and deletion)
