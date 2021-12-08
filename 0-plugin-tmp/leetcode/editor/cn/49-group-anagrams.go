package cn

import (
	"sort"
)

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

type strr []byte

func (st strr) Len() int {
	return len(st)
}

func (st strr) Swap(i, j int) {
	st[i], (st)[j] = (st)[j], (st)[i]
}

func (st strr) Less(i, j int) bool {
	return (st)[i] < (st)[j]
}

func groupAnagrams(strs []string) [][]string {
	ans := make([][]string, 0)
	strMap := make(map[string]int)
	for i := 0; i < len(strs); i++ {
		tmp := strs[i]
		s := strr(tmp)
		sort.Sort(s)
		if idx, ok := strMap[string(s)]; ok {
			ans[idx] = append(ans[idx], strs[i])
		} else {
			ans = append(ans, []string{strs[i]})
			strMap[string(s)] = len(ans) - 1
		}
	}

	return ans

}
//leetcode submit region end(Prohibit modification and deletion)
