package cn

import "strconv"

//给定一个只包含数字的字符串，用以表示一个 IP 地址，返回所有可能从 s 获得的 有效 IP 地址 。你可以按任何顺序返回答案。
//
// 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。 
//
// 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 
//和 "192.168@1.1" 是 无效 IP 地址。 
//
// 示例 1： 
//
//输入：s = "25525511135"
//输出：["255.255.11.135","255.255.111.35"]
//
// 示例 2： 
//
//输入：s = "0000"
//输出：["0.0.0.0"]
//
// 示例 3： 
//
//输入：s = "1111"
//输出：["1.1.1.1"]
//
// 示例 4： 
//
//输入：s = "010010"
//输出：["0.10.0.10","0.100.1.0"]
// 
// 示例 5：
//
//输入：s = "101023"
//输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
//
// 提示： 
// 0 <= s.length <= 3000
// s 仅由数字组成 

//leetcode submit region begin(Prohibit modification and deletion)
func restoreIpAddresses(s string) []string {
	ans := make([]string, 0)
	var dfs func(dot int, buf string, idx int)
	dfs = func(dot int, buf string, idx int) {
		if idx >= len(s) {
			return
		}
		if dot == 4 {
			if s[idx] == '0' && idx < len(s) - 1 {
				return
			}
			num, _ := strconv.Atoi(s[idx:])
			if num > 255 {
				return
			}
			buf += s[idx:]
			ans = append(ans, buf)
			return
		}
		if s[idx] == '0' {
			dfs(dot+1, buf + s[idx:idx+1] + ".", idx+1)
			return
		}
		for i := idx; i < idx + 3 && i < len(s); i++ {
			num, _ := strconv.Atoi(s[idx:i+1])
			if num > 255 {
				break
			}
			dfs(dot+1, buf + s[idx:i+1] + ".", i+1)
		}
	}
	dfs(1, "", 0)
	return ans










}
//func restoreIpAddresses(s string) []string {
//	// dfs
//	ans := make([]string, 0)
//	if len(s) < 4 {
//		return ans
//	}
//	var dfs func(num int, pos int, solution []byte)
//	dfs = func(num int, pos int, solution []byte) {
//		if num > 3 {
//			value, _ := strconv.Atoi(s[pos:])
//			if value > 255 || s[pos] == '0' && pos < len(s)-1 {
//				return
//			}
//			ans = append(ans, string(append(solution, s[pos:]...)))
//			return
//		}
//		for i := pos+1; i < len(s); i++ {
//			value, _ := strconv.Atoi(s[pos:i])
//			if value > 255 || s[pos] == '0' && i > pos+1 {
//				break
//			}
//			solution = append(solution, s[pos:i]...)
//			solution = append(solution, '.')
//			dfs(num+1, i, solution)
//			solution = solution[:len(solution)-i+pos-1]
//		}
//	}
//	dfs(1, 0, []byte{})
//	return ans
//}
//leetcode submit region end(Prohibit modification and deletion)
