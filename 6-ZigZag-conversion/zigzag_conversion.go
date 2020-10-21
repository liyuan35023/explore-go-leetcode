package __ZigZag_conversion

import (
	"bytes"
)

/*
	Z字形变换：https://leetcode-cn.com/problems/zigzag-conversion/

*/
func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	bufs := make([]*bytes.Buffer, numRows)
	for k := range bufs {
		bufs[k] = bytes.NewBuffer([]byte{})
	}
	goDown := false
	curRow := 0

	for _, c := range s {
		bufs[curRow].WriteRune(c)
		if curRow == 0 || curRow == numRows-1 {
			goDown = !goDown
		}
		if goDown {
			curRow++
		} else {
			curRow--
		}
	}
	ret := bytes.NewBuffer([]byte{})
	for _, buf := range bufs {
		ret.WriteString(buf.String())
		ret.Write(buf.Bytes())
	}
	return ret.String()
}

func convert2(s string, numRows int) string {
	if len(s) < 2 || numRows < 2 {
		return s
	}
	bufs := make([][]byte, numRows)
	for k := range bufs {
		bufs[k] = make([]byte, 0)
	}
	curRow := 0
	goDown := false
	for i := 0; i < len(s); i++ {
		bufs[curRow] = append(bufs[curRow], s[i])
		if curRow == 0 || curRow == numRows-1 {
			goDown = !goDown
		}
		if goDown {
			curRow++
		} else {
			curRow--
		}
	}
	ans := bytes.NewBuffer([]byte{})
	for _, buf := range bufs {
		ans.Write(buf)
	}
	return ans.String()
}
