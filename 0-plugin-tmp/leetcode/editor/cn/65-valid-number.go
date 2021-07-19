package cn
//有效数字（按顺序）可以分成以下几个部分： 
//
// 一个 小数 或者 整数 
// （可选）一个 'e' 或 'E' ，后面跟着一个 整数 
// 
//
// 小数（按顺序）可以分成以下几个部分： 
//
// （可选）一个符号字符（'+' 或 '-'） 
// 下述格式之一：
// 
// 至少一位数字，后面跟着一个点 '.' 
// 至少一位数字，后面跟着一个点 '.' ，后面再跟着至少一位数字 
// 一个点 '.' ，后面跟着至少一位数字 
// 
// 整数（按顺序）可以分成以下几个部分：
//
// （可选）一个符号字符（'+' 或 '-'） 
// 至少一位数字 
//
// 部分有效数字列举如下： 
//
// ["2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1",
// "53.5e93", "-123.456e789"] 
//
// 部分无效数字列举如下： 
//
// ["abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53"] 
//
// 给你一个字符串 s ，如果 s 是一个 有效数字 ，请返回 true 。 
//
// 示例 1： 
//
//输入：s = "0"
//输出：true
//
// 示例 2： 
//
//输入：s = "e"
//输出：false
//
// 示例 3： 
//
//输入：s = "."
//输出：false
//
// 示例 4： 
//
//输入：s = ".1"
//输出：true
//
// 提示： 
//
// 1 <= s.length <= 20
// s 仅含英文字母（大写和小写），数字（0-9），加号 '+' ，减号 '-' ，或者点 '.' 。 

//leetcode submit region begin(Prohibit modification and deletion)
type state int
type input int

const (
	Init state = iota
	Signed
	Integer
	NoLeftNumDot
	LeftNumDot
	Decimal
	EState
	ESigned
	EInteger
	End
)

const (
	Num input = iota
	Sign
	E
	Dot
	other
)

func isNumber(s string) bool {
	stateMap := []map[input]state{
		{Num:Integer, Sign:Signed, E:End, Dot: NoLeftNumDot, other:End},
		{Num:Integer, Sign:End, E:End, Dot: NoLeftNumDot, other:End},
		{Num:Integer, Sign:End, E:EState, Dot: LeftNumDot, other:End},
		{Num:Decimal, Sign:End, E:End, Dot: End, other:End},
		{Num:Decimal, Sign:End, E:EState, Dot: End, other:End},
		{Num:Decimal, Sign:End, E:EState, Dot: End, other:End},
		{Num:EInteger, Sign:ESigned, E:End, Dot: End, other:End},
		{Num:EInteger, Sign:End, E:End, Dot: End, other:End},
		{Num:EInteger, Sign:End, E:End, Dot: End, other:End},
	}
	curState := Init
	for i := 0; i < len(s); i++ {
		in := getInput(s[i])
		curState = stateMap[curState][in]
		if curState == End {
			return false
		}
	}
	if curState == Integer || curState == LeftNumDot || curState == Decimal || curState == EInteger {
		return true
	}
	return false
}

func getInput(b byte) input {
	switch {
	case b == '+' || b == '-':
		return Sign
	case b == 'e' || b == 'E':
		return E
	case b == '.':
		return Dot
	case b >= '0' && b <= '9':
		return Num
	default:
		return other
	}
}
//leetcode submit region end(Prohibit modification and deletion)
