package cn

//给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位
//
// 返回滑动窗口中的最大值。 
//
// 示例 1： 
//
//输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
//输出：[3,3,5,5,6,7]
//解释：
//滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
//
// 示例 2： 
//
//输入：nums = [1], k = 1
//输出：[1]
//
// 示例 3： 
//
//输入：nums = [1,-1], k = 1
//输出：[1,-1]
//
// 示例 4： 
//
//输入：nums = [9,11], k = 2
//输出：[11]
//
// 示例 5： 
//
//输入：nums = [4,-2], k = 2
//输出：[4] 
//
// 提示： 
// 1 <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴ 
// 1 <= k <= nums.length 
// 

//leetcode submit region begin(Prohibit modification and deletion)
func maxSlidingWindow(nums []int, k int) []int {
	ans := make([]int, 0)
	dequeue := make([]int, 0)
	for i := 0; i < k-1; i++ {
		for len(dequeue) > 0 && nums[dequeue[len(dequeue)-1]] <= nums[i] {
			dequeue = dequeue[:len(dequeue)-1]
		}
		dequeue = append(dequeue, i)
	}
	l, r := 0, k - 1
	for ; r < len(nums); l, r = l+1, r+1 {
		for len(dequeue) > 0 && nums[dequeue[len(dequeue)-1]] <= nums[r] {
			dequeue = dequeue[:len(dequeue)-1]
		}
		dequeue = append(dequeue, r)
		for len(dequeue) > 0 && dequeue[0] < l {
			dequeue = dequeue[1:]
		}
		ans = append(ans, nums[dequeue[0]])
	}
	return ans
}


//leetcode submit region end(Prohibit modification and deletion)
