package cn

import "sort"

//给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
//
// 示例: 
//
// 输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出:
//[
//  ["ate","eat","tea"],
//  ["nat","tan"],
//  ["bat"]
//] 
//
// 说明： 
//
// 所有输入均为小写字母。 
// 不考虑答案输出的顺序。 

//leetcode submit region begin(Prohibit modification and deletion)

type strrrrr []byte

func (s strrrrr) Len() int {
	return len(s)
}

func (s strrrrr) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s strrrrr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func groupAnagrams(strs []string) [][]string {
	ans := make([][]string, 0)
	wordsMap := make(map[string]int)
	for k, v := range strs {
		alias := strrrrr(v)
		sort.Sort(alias)
		newStr := string(alias)
		if idx, ok := wordsMap[newStr]; !ok {
			ans = append(ans, append([]string{}, strs[k]))
			wordsMap[newStr] = len(ans)-1
		} else {
			ans[idx] = append(ans[idx], strs[k])
		}
	}
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
