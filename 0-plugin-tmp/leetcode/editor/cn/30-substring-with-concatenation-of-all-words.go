package cn
//给定一个字符串 s 和一些 长度相同 的单词 words 。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。 
//
// 注意子串要与 words 中的单词完全匹配，中间不能有其他字符 ，但不需要考虑 words 中单词串联的顺序。 
//
// 示例 1：
//
//输入：s = "barfoothefoobarman", words = ["foo","bar"]
//输出：[0,9]
//解释：
//从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
//输出的顺序不重要, [9,0] 也是有效答案。
// 
// 示例 2：
//
//输入：s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
//输出：[]
//
// 示例 3： 
//
//输入：s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
//输出：[6,9,12]
//
// 提示： 
//
// 1 <= s.length <= 104 
// s 由小写英文字母组成 
// 1 <= words.length <= 5000 
// 1 <= words[i].length <= 30 
// words[i] 由小写英文字母组成 
// 


//leetcode submit region begin(Prohibit modification and deletion)
func findSubstring(s string, words []string) []int {
	ans := make([]int, 0)
	wordLen := len(words[0])
	num := len(words)
	wordMap := make(map[string]int)
	for _, s := range words {
		wordMap[s]++
	}
	count := 0
	for i := 0; i < len(s) - wordLen*num + 1; i++ {
		if _, ok := wordMap[s[i:i+wordLen]]; ok {
			tmp := i
			for count < num {
				if k, ok := wordMap[s[tmp:tmp+wordLen]]; ok && k > 0 {
					wordMap[s[tmp:tmp+wordLen]] = k - 1
					count++
					tmp += wordLen
				} else {
					break
				}
			}
			if count == num {
				ans = append(ans, i)
			}
			tmp = tmp - wordLen
			for count > 0 {
				wordMap[s[tmp:tmp+wordLen]]++
				tmp -= wordLen
				count--
			}
		}
	}
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
