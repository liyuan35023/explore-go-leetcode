package _65_compare_version

import "strconv"

// 如果 version1 > version2 返回 1，
// 如果 version1 < version2 返回 -1，
// 除此之外返回 0。

func compareVersion(version1 string, version2 string) int {
	i, j := 0, 0
	for i < len(version1) || j < len(version2) {
		num1, num2 := 0, 0
		if i < len(version1) {
			end := i + 1
			for end < len(version1) && version1[end] != '.' {
				end++
			}
			num1, _ = strconv.Atoi(version1[i:end])
			i = end + 1
		}
		if j < len(version2) {
			end := j + 1
			for end < len(version2) && version2[end] != '.' {
				end++
			}
			num2, _ = strconv.Atoi(version2[j:end])
			j = end + 1
		}
		if num1 > num2 {
			return 1
		} else if num1 < num2 {
			return -1
		}
	}
	return 0
}
