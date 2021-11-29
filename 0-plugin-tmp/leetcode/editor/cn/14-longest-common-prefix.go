package cn

//编写一个函数来查找字符串数组中的最长公共前缀。
//
// 如果不存在公共前缀，返回空字符串 ""。 
//
// 示例 1：
//
//输入：strs = ["flower","flow","flight"]
//输出："fl"
//
// 示例 2： 
//
//输入：strs = ["dog","racecar","car"]
//输出：""
//解释：输入不存在公共前缀。 
//
//
// 提示： 
//
// 
// 0 <= strs.length <= 200 
// 0 <= strs[i].length <= 200 
// strs[i] 仅由小写英文字母组成 
// 

//leetcode submit region begin(Prohibit modification and deletion)

func longestCommonPrefix(strs []string) string {
	switch len(strs) {
	case 0:
		return ""
	case 1:
		return strs[0]
	}
	min := strs[0]
	max := strs[0]
	for _, s := range strs {
		if s < min {
			min = s
		} else if s > max {
			max = s
		}
	}
	for i := 0; i < len(min) && i < len(max); i++ {
		if max[i] != min[i] {
			return min[:i]
		}
	}
	return min
}
//leetcode submit region end(Prohibit modification and deletion)
