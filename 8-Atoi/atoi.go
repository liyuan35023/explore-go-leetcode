package __Atoi

import (
	"math"
	"unicode"
)

/*
请你来实现一个atoi函数，使其能将字符串转换成整数。

首先，该函数会根据需要丢弃无用的开头空格字符，直到寻找到第一个非空格的字符为止。接下来的转化规则如下：

如果第一个非空字符为正或者负号时，则将该符号与之后面尽可能多的连续数字字符组合起来，形成一个有符号整数。
假如第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成一个整数。
该字符串在有效的整数部分之后也可能会存在多余的字符，那么这些字符可以被忽略，它们对函数不应该造成影响。
注意：假如该字符串中的第一个非空格字符不是一个有效整数字符、字符串为空或字符串仅包含空白字符时，则你的函数不需要进行转换，即无法进行有效转换。

在任何情况下，若函数不能进行有效的转换时，请返回 0 。

提示：

本题中的空白字符只包括空格字符 ' ' 。
假设我们的环境只能存储 32 位大小的有符号整数，那么其数值范围为[−231, 2− 1]。如果数值超过这个范围，请返回 INT_MAX (231− 1) 或INT_MIN (−231) 。

 */

// 状态机来解决

var stateMap = map[string][]string{
	"init":[]string{"init", "signed", "in", "end"},
	"signed":[]string{"end", "end", "in", "end"},
	"in":[]string{"end", "end", "in", "end"},
	"end":[]string{"end", "end", "end", "end"},
}

func getColumn(c rune) int {
	if c == ' ' || c == '	' {
		return 0
	} else if c == '+' || c == '-' {
		return 1
	} else if c >= '0' && c <= '9' {
		return 2
	} else {
		return 3
	}
}


func MyAtoi(str string) int {
	curState := "init"
	ret := 0
	sign := 1
	for _, c := range str {
		curState = stateMap[curState][getColumn(c)]
		if curState == "in" {
			tmp := int(c-'0')
			if sign == 1 {
				if ret > (math.MaxInt32 - tmp) / 10 {
					ret = math.MaxInt32
					return ret
				}
			} else {
				if ret > (-math.MinInt32 - tmp) / 10 {
					ret = math.MinInt32
					return ret
				}
			}
			ret = ret*10 + int(c-'0')
		} else if curState == "signed" && c == '-' {
			sign = -1
		} else if curState == "end" {
			break
		}
	}
	return ret*sign
}


type automation struct {
	state string
	ans   int
	sign  int
	table map[string][]string
}

func newAutomation() *automation {
	return &automation{
		state: "start",
		ans:   0,
		sign:  1,
		table: getTable(),
	}
}

func getTable() map[string][]string {
	m := map[string][]string{
		"start":     {"start", "signed", "in_number", "end"},
		"signed":    {"end", "end", "in_number", "end"},
		"in_number": {"end", "end", "in_number", "end"},
		"end":       {"end", "end", "end", "end"},
	}
	return m
}

func (auto *automation) getCol(r rune) int {
	switch {
	case unicode.IsSpace(r):
		return 0
	case r == '+' || r == '-':
		return 1
	case unicode.IsDigit(r):
		return 2
	default:
		return 3
	}
}
func (auto *automation) get(r rune) {

	col := auto.getCol(r)
	auto.state = auto.table[auto.state][col]
	if auto.state == "in_number" {
		auto.ans = auto.ans*10 + int(r-'0')
		if auto.sign == 1 {
			if auto.ans > math.MaxInt32 {
				auto.ans = math.MaxInt32
			}
		} else {
			if auto.ans * auto.sign < math.MinInt32 {
				auto.ans = -math.MinInt32
			}
		}

	}
	if auto.state == "signed" {
		if r == '+' {
			auto.sign = 1
		} else {
			auto.sign = -1
		}
	}
}

func (auto *automation) myAtoi(str string) int {
	for _, r := range str {
		auto.get(r)
	}
	return auto.sign * auto.ans
}
func MMyAtoi(str string) int {
	auto := newAutomation()
	result := auto.myAtoi(str)
	return result
}

