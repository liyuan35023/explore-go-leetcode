package _4_lcp

func longestCommonPrefix(strs []string) string {
	switch len(strs) {
	case 0:
		return ""
	case 1:
		return strs[0]
	}

	for i:=0; i<len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[0][i] != strs[j][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

func lcp(strs []string) string {
	switch len(strs) {
	case 0:
		return ""
	case 1:
		return strs[0]
	}
	min, max := strs[0], strs[0]
	for _, s := range strs {
		if s < min {
			min = s
		} else if s > max {
			max = s
		}
	}
	for i := 0; i < len(min) && i < len(max); i++ {
		if max[i] != min[i] {
			return min[:i]
		}
	}
	return min
}