package _8_strStr


/*
	实现一个查找 substring 的函数。如果在母串中找到了子串，返回子串在母串中出现的下标，如果没有找到，返回 -1，如果子串是空串，则返回 0。
	Example 1:

	Input: haystack = "hello", needle = "ll"
	Output: 2

	Example 2:

	Input: haystack = "aaaaa", needle = "bba"
	Output: -1
 */

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}
	needleLen := len(needle)
	j := 0
	for i := 0; i < len(haystack) - len(needle) + 1; i++ {
		tmp := i
		for haystack[tmp] == needle[j] {
			tmp++
			j++
			if j == needleLen {
				return i
			}
		}
		j = 0
	}
	return -1
}
