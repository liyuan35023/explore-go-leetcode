package cn
//给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。 
//
// 注意：如果 s 中存在这样的子串，我们保证它是唯一的答案。 
//
// 示例 1： 
//
//输入：s = "ADOBECODEBANC", t = "ABC"
//输出："BANC"
// 
// 示例 2：
//
//输入：s = "a", t = "a"
//输出："a"
//
// 提示： 
//
// 1 <= s.length, t.length <= 105 
// s 和 t 由英文字母组成 
//
//进阶：你能设计一个在 o(n) 时间内解决此问题的算法吗？

//leetcode submit region begin(Prohibit modification and deletion)
func minWindow(s string, t string) string {
	ans := ""
	charMap := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		charMap[t[i]]++
	}
	left, right := 0, 0
	existNum := 0
	for ; right < len(s); right++ {
		if v, ok := charMap[s[right]]; ok {
			charMap[s[right]]--
			if v > 0 {
				existNum++
			}
		}
		for ; existNum == len(t); left++ {
			if ans == "" || right - left + 1 < len(ans) {
				ans = s[left:right+1]
			}

			if v, ok := charMap[s[left]]; ok {
				if v == 0 {
					existNum--
				}
				charMap[s[left]]++
			}
		}
	}
	return ans
}
//func minWindow(s string, t string) string {
//	left, right := 0, 0
//	charMap := make(map[byte]int)
//	for i := 0; i < len(t); i++ {
//		charMap[t[i]]++
//	}
//	exist := 0
//	ans := ""
//	for ; right < len(s); right++ {
//		if num, ok := charMap[s[right]]; ok {
//			if num > 0 {
//				exist++
//			}
//			charMap[s[right]] = num - 1
//		}
//		for exist == len(t) {
//			// move left
//			if len(ans) == 0 || right - left + 1 < len(ans) {
//				ans = s[left:right+1]
//			}
//			if num, ok := charMap[s[left]]; ok {
//				if num == 0 {
//					exist--
//				}
//				charMap[s[left]] = num + 1
//			}
//			left++
//		}
//	}
//	return ans
//}

//leetcode submit region end(Prohibit modification and deletion)
