package _3_roman2Integer

func romanToInt(s string) int {
	_map := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	ret := 0
	pre := 0
	for i := len(s)-1; i >= 0 ; i-- {
		if _map[s[i]] < pre {
			ret -= _map[s[i]]
		} else {
			ret += _map[s[i]]
		}
		pre = _map[s[i]]
	}
	return ret
}
