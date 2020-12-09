package offer

import "bytes"

/*
	请实现一个函数，把字符串 s 中的每个空格替换成"%20"。


	示例 1：

	输入：s = "We are happy."
	输出："We%20are%20happy."

	限制：

	0 <= s 的长度 <= 10000


 */

func replaceSpace(s string) string {
	ans := bytes.NewBuffer([]byte{})
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			ans.WriteString("%20")
		} else {
			ans.WriteByte(s[i])
		}
	}
	return ans.String()
}

