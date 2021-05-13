package __Longest_subString

import (
	"strings"
)

/**
 * Given a string, find the length of the longest substring without repeating characters.
 * Example 1:
 * Input: "abcabcbb"
 * Output: 3
 * Explanation: The answer is "abc", with the length of 3.

 * Example 2:
 * Input: "bbbbb"
 * Output: 1
 * Explanation: The answer is "b", with the length of 1.

 * Example 3:
 * Input: "pwwkew"
 * Output: 3
 * Explanation: The answer is "wke", with the length of 3.
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

 */

func LengthOfLongestSubstringFinal(s string) int {
	ans := 0
	left := 0
	charMap := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		if idx, ok := charMap[s[i]]; ok && idx >= left {
			left = idx + 1
		} else {
			ans = max(ans, i - left + 1)
		}
		charMap[s[i]] = i
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func LengthOfLongestSubstringBruteForce(s string) int {
	n := len(s)
	longest := 0
	i := 0
	j := n
	for ; i < n; i ++ {
		for j = n; j >= i; j--  {
			subString := s[i:j]
			noRepeat := true
			for k, r := range subString {
				if exist := strings.ContainsRune(subString[k+1:], r); exist {

					noRepeat = false
					break
				}
			}
			if noRepeat && j-i > longest {
				longest = j-i
				break
			}
		}

	}

	return longest
}

func Longest(s string) int {
	characterMap := make(map[int32]int, 0)
	longest := 0
	left := 0
	for k, char := range s {
		if oldKey, ok := characterMap[char]; ok && oldKey >= left {
			left = oldKey + 1
		}
		characterMap[char] = k
		if k - left + 1 > longest {
			longest = k - left + 1
		}
	}
	return longest
}

func LengthOfLongestSubstring(s string) int {
	i, j := 0, 1
	longest, cur := 1, 1
	for ;j < len(s); j++ {
		subString := s[i:j]
		if strings.ContainsRune(subString, rune(s[j])) {
			index := strings.IndexRune(subString, rune(s[j]))
			i = index+1+i
		}
		cur = j-i+1
		if cur > longest {
			longest = cur
		}
	}
	return longest
}