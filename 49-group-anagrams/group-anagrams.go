package _9_group_anagrams

import "sort"

/*
	Example:
	Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
	Output:
	[
	  ["ate","eat","tea"],
	  ["nat","tan"],
	  ["bat"]
	]

	Note:
	All inputs will be in lowercase.
	The order of your output does not matter.

	题目大意 #
	给出一个字符串数组，要求对字符串数组里面有 Anagrams 关系的字符串进行分组。
	Anagrams 关系是指两个字符串的字符完全相同，顺序不同，两者是由排列组合组成。

 */

type aliasStr []rune

func (s aliasStr) Len() int {
	return len(s)
}

func (s aliasStr) Less(i, j int) bool {
	if s[i] < s[j] {
		return true
	}
	return false
}

func (s aliasStr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func GroupAnagrams(strs []string) [][]string {
	ans := make([][]string, 0)
	recordMap := make(map[string]int)

	for k, str := range strs {
		tmp := aliasStr(str)
		sort.Sort(tmp)
		str = string(tmp)
		if index, ok := recordMap[str]; ok {
			ans[index] = append(ans[index], strs[k])
		} else {
			recordMap[str] = len(ans)
			ans = append(ans, []string{strs[k]})
		}
	}
	return ans
}