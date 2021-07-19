package cn

import (
	"math"
	"sort"
)

//给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和
//。假定每组输入只存在唯一答案。 
//
// 
//
// 示例： 
//
// 输入：nums = [-1,2,1,-4], target = 1
//输出：2
//解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
// 
//
// 
//
// 提示： 
//
// 
// 3 <= nums.length <= 10^3 
// -10^3 <= nums[i] <= 10^3 
// -10^4 <= target <= 10^4 
// 
// Related Topics 数组 双指针 
// 👍 797 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func threeSumClosest(nums []int, target int) int {
	ans := math.MinInt32
	updateAns := func(a, b, c int) {
		if abs(ans-target) > abs(nums[a]+nums[b]+nums[c]-target) {
			ans = nums[a] + nums[b] + nums[c]
		}
	}
	sort.Ints(nums)
	for left := 0; left < len(nums)-2; left++ {
		remain := target - nums[left]
		if left > 0 && nums[left-1] == nums[left] {
			continue
		}
		mid, right := left+1, len(nums)-1
		for mid < right {
			if nums[mid] + nums[right] == remain {
				return target
			}
			updateAns(left, mid, right)
			if nums[mid] + nums[right] < remain {
				mid++
			} else {
				right--
			}
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}
//leetcode submit region end(Prohibit modification and deletion)
