package _2_integer2Roman

import "bytes"

/*
	整数转罗马数字
	贪心算法
 */

var nums = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
var romans = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}


func intToRoman(num int) string {
	ret := bytes.NewBuffer([]byte{})
	for i := 0; i < len(romans) && num > 0; i++ {
		for num-nums[i] >= 0 {
			ret.WriteString(romans[i])
			num = num-nums[i]
		}
	}
	return ret.String()
}
