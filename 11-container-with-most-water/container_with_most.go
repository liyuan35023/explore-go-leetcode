package _1_container_with_most_water

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	lVal, rVal, distance := 0, 0, 0
	cur, max := 0, 0

	for left < right {
		lVal = height[left]
		rVal = height[right]
		distance = right-left
		if lVal < rVal {
			cur = lVal * distance
			left++
		} else {
			cur = rVal * distance
			right--
		}
		if cur > max {
			max = cur
		}
	}
	return max
}
