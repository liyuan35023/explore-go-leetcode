package _0_SubString_with_concatenation

/*
	给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
	注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。

	示例 1：

	输入：
	  s = "barfoothefoobarman",
	  words = ["foo","bar"]
	输出：[0,9]
	解释：
	从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
	输出的顺序不重要, [9,0] 也是有效答案。
	示例 2：

	输入：
	  s = "wordgoodgoodgoodbestword",
	  words = ["word","good","best","word"]
	输出：[]


 */

func SubString(s string, words []string) []int {
	if len(s) < 1 || len(words) < 1 {
		return []int{}
	}
	wordLen := len(words[0])
	wordsNum := len(words)
	n := len(s)
	if n < wordLen * wordsNum {
		return []int{}
	}

	wordsMap := make(map[string]int)
	for _, word := range words {
		if k, ok := wordsMap[word]; ok {
			wordsMap[word] = k + 1
		} else {
			wordsMap[word] = 1
		}
	}

	ans := make([]int, 0)
	for i := 0; i < n; i++ {
		if n - i < wordLen * wordsNum {
			break
		}

		start := i
		index := 1
		for ; index <= wordsNum; index++ {
			tmp := s[start:start+wordLen]
			if k, ok := wordsMap[tmp]; ok && k > 0 {
				wordsMap[tmp] = k - 1
			} else {
				break
			}
			start = start + wordLen
		}
		if index == wordsNum + 1 {
			ans = append(ans, i)
		}
		//  把已经减去的加回来
		for j := 1; j < index; j++ {
			cur := s[i+wordLen*(j-1):i+wordLen*j]
			wordsMap[cur] = wordsMap[cur] + 1
		}
	}
	return ans
}
