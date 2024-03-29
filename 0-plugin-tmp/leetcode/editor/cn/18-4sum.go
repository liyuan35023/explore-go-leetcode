package cn

import "sort"

//给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c +
// d 的值与 target 相等？找出所有满足条件且不重复的四元组。 
//
// 注意：答案中不可以包含重复的四元组。 
//
//
// 示例 1： 
//
//输入：nums = [1,0,-1,0,-2,2], target = 0
//输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
// 
//
// 示例 2： 
//
//输入：nums = [], target = 0
//输出：[]
//
// 提示：
//
// 0 <= nums.length <= 200 
// -109 <= nums[i] <= 109 
// -109 <= target <= 109 
// 


//leetcode submit region begin(Prohibit modification and deletion)
func fourSum(nums []int, target int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(nums)
	for first := 0; first < len(nums); first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		for second := first+1; second < len(nums); second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			third, fourth := second+1, len(nums)-1
			for third < fourth {
				sum := nums[first] + nums[second] + nums[third] + nums[fourth]
				if sum == target {
					ans = append(ans, []int{nums[first], nums[second], nums[third], nums[fourth]})
					third++
					for third < fourth && nums[third] == nums[third-1] {
						third++
					}
				} else if sum < target {
					third++
				} else {
					fourth--
				}
			}
		}
	}
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
