package _6_plus_one

/*
	给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。最高位数字存放在数组的首位,
	数组中每个元素只存储单个数字。你可以假设除了整数 0 之外，这个整数不会以零开头。

	Example 1:

	Input: [1,2,3]
	Output: [1,2,4]
	Explanation: The array represents the integer 123.
	Example 2:

	Input: [4,3,2,1]
	Output: [4,3,2,2]
	Explanation: The array represents the integer 4321.
 */

func plusOne(digits []int) []int {
	n := len(digits)
	if digits[n-1] + 1 != 10 {
		digits[n-1] += 1
		return digits
	}
	carry := 1
	for i := len(digits) - 1; i >= 0 && carry != 0; i-- {
		digits[i], carry = (digits[i] + carry) % 10, (digits[i] + carry) / 10
	}
	if carry == 1 {
		ans := append([]int{carry}, digits...)
		return ans
	}
	return digits
}
