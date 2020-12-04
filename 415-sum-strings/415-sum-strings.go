package _15_sum_strings

import (
	"bytes"
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	m := len(num1)
	n := len(num2)
	maxLen := 0
	if m > n {
		maxLen = m
	} else {
		maxLen = n
	}

	sum := make([]int, maxLen+1)
	for i := 0; i < maxLen; i++ {
		val1, val2 := 0, 0
		if i < m {
			val1 = int(num1[m-i-1] - '0')
		}
		if i < n {
			val2 = int(num2[n-i-1] - '0')
		}
		sum[i] = val1 + val2
	}
	for i := 0; i < maxLen; i++ {
		sum[i+1] += sum[i] / 10
		sum[i] = sum[i] % 10
	}
	buf := bytes.NewBuffer([]byte{})
	for i := maxLen; i >= 0; i-- {
		if i == maxLen && sum[i] == 0 {
			continue
		}
		buf.WriteString(strconv.Itoa(sum[i]))
	}
	return buf.String()
}

