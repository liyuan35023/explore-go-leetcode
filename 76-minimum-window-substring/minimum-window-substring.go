package _6_minimum_window_substring

/*
	Example:

	Input: S = "ADOBECODEBANC", T = "ABC"
	Output: "BANC"

	Note:
	If there is no such window in S that covers all characters in T, return the empty string “”.
	If there is such window, you are guaranteed that there will always be only one unique minimum window in S.
	题目大意 #
	给定一个源字符串 s，再给一个字符串 T，要求在源字符串中找到一个窗口，这个窗口包含由字符串各种排列组合组成的，
	窗口中可以包含 T 中没有的字符，如果存在多个，在结果中输出最小的窗口，如果找不到这样的窗口，输出空字符串。
 */
func MinWindow(s string, t string) string {
	// 滑动窗口
	if len(s) < 1 || len(t) < 1 || len(s) < len(t) {
		return ""
	}
	ans := ""
	update := func(i, j int) {
		if ans == "" || ans != "" && j - i + 1 < len(ans) {
			ans = s[i:j+1]
		}
	}

	left, right := 0, 0
	// 用两个map, cmap 用来记录t中byte出现的次数
	cMap := make(map[byte]int)
	rMap := make(map[byte]int)

	check := func() bool {
		for k, v := range rMap {
			if cMap[k] < v {
				return false
			}
		}
		return true
	}

	for k := range t {
		rMap[t[k]] = rMap[t[k]] + 1
	}

	for ; right < len(s); right++ {
		rByte := s[right]
		if _, ok := rMap[rByte]; ok {
			cMap[rByte]++
		}
		// check s[left:right+1] include all t
		// now move left
		for check() && left <= right {
			if k, ok := cMap[s[left]]; ok {
				cMap[s[left]] = k - 1
			}
			update(left, right)
			left++
		}
	}
	return ans
}

func minWindow(s string, t string) string {
	if len(s) < 1 || len(t) < 1 || len(s) < len(t) {
		return ""
	}

	ans := ""
	n := len(t)
	countMap := make(map[byte]int)
	for i := 0; i < n; i++ {
		countMap[t[i]]++
	}
	left, right := 0, 0

	update := func(i, j int) {
		if ans == "" || j - i + 1 < len(ans) {
			ans = s[i:j+1]
		}
	}

	for ; right < len(s); right++ {
		if v, ok := countMap[s[right]]; ok {
			if v > 0 {
				n--
			}
			countMap[s[right]]--
		}
		// 移动左指针
		for n == 0 {
			update(left, right)
			if v, ok := countMap[s[left]]; ok {
				if v == 0 {
					n++
				}
				countMap[s[left]]++

			}
			left++
		}
	}
	return ans
}
