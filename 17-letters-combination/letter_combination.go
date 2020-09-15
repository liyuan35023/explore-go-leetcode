package _7_letters_combination

var _map = map[byte]string{
	'2':"abc",
	'3':"def",
	'4':"ghi",
	'5':"jkl",
	'6':"mno",
	'7':"pqrs",
	'8':"tuv",
	'9':"wxyz",
}
var ret []string
func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	backTrace(digits, 0, "")
	return ret
}

func backTrace(digits string, cur int, com string) {
	if cur == len(digits) {
		ret = append(ret, com)
		return
	}
	for _, r := range _map[digits[cur]] {
		backTrace(digits, cur+1, com+string(r))
	}
}

