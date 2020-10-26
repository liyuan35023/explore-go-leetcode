package offer

import "sort"

/*
	输入一个字符串，打印出该字符串中字符的所有排列。
	你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。

	示例:

	输入：s = "abc"
	输出：["abc","acb","bac","bca","cab","cba"]

	1 <= s 的长度 <= 8

 */

var ans []string
var used []bool
var solution []byte

type aliasByte []byte

func (a aliasByte) Len() int {
	return len(a)
}

func (a aliasByte) Less(i , j int) bool {
	if a[i] < a[j] {
		return true
	}
	return false
}

func (a aliasByte) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func Permutation(s string) []string {
	// 回溯法
	if len(s) == 1 {
		return []string{s}
	}
	// 排序

	tmp := []byte(s)
	sort.Sort(aliasByte(tmp))
	s = string(tmp)

	ans = make([]string, 0)
	solution = make([]byte, 0)
	used = make([]bool, len(s))
	backTrace(0, s)
	return ans
}

func backTrace(num int, s string) {
	if num == len(s) {
		tmp := string(solution)
		ans = append(ans, tmp)
		return
	}
	for i := 0; i < len(s); i++ {
		if used[i] || i > 0 && s[i] == s[i-1] && !used[i-1] {
			continue
		}
		solution = append(solution, s[i])
		used[i] = true
		backTrace(num+1, s)
		solution = solution[:len(solution)-1]
		used[i] = false
	}
}

