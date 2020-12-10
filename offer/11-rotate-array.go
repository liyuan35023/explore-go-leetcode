package offer

/*
	把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
	输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。例如，数组[3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。

	示例 1：

	输入：[3,4,5,1,2]
	输出：1
	示例 2：

	输入：[2,2,2,0,1]
	输出：0

    == 主站154
 */

func minArray(numbers []int) int {
	left, right := 0, len(numbers)-1
	for left < right {
		mid := (left + right) / 2
		if mid > 0 && numbers[mid] < numbers[mid-1] {
			return mid
		}
		if numbers[left] <= numbers[mid] {
			left = mid
		} else if numbers[left] > numbers[mid] {
			right = mid
		}
	}
	return left
}
