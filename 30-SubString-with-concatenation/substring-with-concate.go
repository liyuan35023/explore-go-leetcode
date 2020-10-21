package _0_SubString_with_concatenation

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
