package _1_decode_ways

import "strconv"

func numDecodings(s string) int {
	// dfs 会超时
	if s[0] == '0' {
		return 0
	}
	ans := 0
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx >= len(s) {
			ans++
			return
		}
		if s[idx] == '0' {
			return
		}
		for i := idx+1; i <= idx+2 && i <= len(s); i++ {
			num, _ := strconv.Atoi(s[idx:i])
			if num <= 26 {
				dfs(i)
			}
		}
	}
	dfs(0)
	return ans
}
func numDecodingsII(s string) int {

}
