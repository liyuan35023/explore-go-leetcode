package _3_restore_ip

import "strconv"

/*
	给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。
	有效的 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
	例如："0.1.2.201" 和 "192.168.1.1" 是 有效的 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效的 IP 地址。

	示例 1：
	输入：s = "25525511135"
	输出：["255.255.11.135","255.255.111.35"]

	示例 2：
	输入：s = "0000"
	输出：["0.0.0.0"]

	示例 3：
	输入：s = "1111"
	输出：["1.1.1.1"]

	示例 4：
	输入：s = "010010"
	输出：["0.10.0.10","0.100.1.0"]

	示例 5：
	输入：s = "101023"
	输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]

	提示：
	0 <= s.length <= 3000
	s 仅由数字组成

 */

func restoreIpAddresses(s string) []string {
	if len(s) < 1 || len(s) > 12 {
		return nil
	}
	ans := make([]string, 0)
	ip := make([]byte, 0)
	backTrace(s, 1, 0, ip, &ans)
	return ans
}

func backTrace(s string, numOfPoint int, startIndex int, ip []byte, ans *[]string) {
	if numOfPoint == 4 {
		if startIndex > len(s)-1 ||  s[startIndex] == '0' && startIndex != len(s)-1 {
			return
		}
		number, _ := strconv.Atoi(s[startIndex:])
		if number <= 255 {
			ip = append(ip, s[startIndex:]...)
			*ans = append(*ans, string(ip))
		}
		return
	}

	for i := 0; i < 3 && startIndex+i < len(s); i++ {
		number, _ := strconv.Atoi(s[startIndex:startIndex+i+1])
		if number <= 255 {
			ip = append(ip, s[startIndex:startIndex+i+1]...)
			ip = append(ip, '.')
			backTrace(s, numOfPoint+1, startIndex+i+1, ip, ans)
			ip = ip[:startIndex+numOfPoint-1]
		}
		if s[startIndex] == '0' {
			break
		}
	}
}

func restoreIpAddressesII(s string) []string {
	ans := make([]string, 0)
	if len(s) < 4 {
		return ans
	}
	var dfs func(num int, start int)
	solution := []byte{}
	dfs = func(num int, start int) {
		if num == 4 {
			if start >= len(s) || s[start] == '0' && start < len(s) - 1 {
				return
			}
			remain, _ := strconv.Atoi(s[start:])
			if remain <= 255 {
				tmp := append(solution, s[start:]...)
				ans = append(ans, string(tmp))
			}
			return
		}
		if s[start] == '0' {
			solution = append(solution, s[start])
			solution = append(solution, '.')
			dfs(num+1, start+1)
			solution = solution[:len(solution)-2]
		} else {
			for i := 1; i <= 3 && start + i <= len(s); i++ {
				n, _ := strconv.Atoi(s[start:start+i])
				if n <= 255 {
					solution = append(solution, s[start:start+i]...)
					solution = append(solution, '.')
					dfs(num+1, start+i)
					solution = solution[:len(solution)-i-1]
					//solution = solution[:start+num-1]
				}
			}
		}
	}
	dfs(1, 0)
	return ans
}

