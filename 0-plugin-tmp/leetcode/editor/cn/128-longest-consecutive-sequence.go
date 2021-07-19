package cn
//给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。 
//
// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。 
//
// 示例 1： 
//
//输入：nums = [100,4,200,1,3,2]
//输出：4
//解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。 
//
// 示例 2： 
//
//输入：nums = [0,3,7,2,5,8,4,6,0,1]
//输出：9
// 
// 提示：
//
// 0 <= nums.length <= 105 
// -109 <= nums[i] <= 109 
// 

//leetcode submit region begin(Prohibit modification and deletion)
func longestConsecutive(nums []int) int {
	numMap := make(map[int][]int)
	ans := 0
	for _, v := range nums {
		if _, ok := numMap[v]; ok {
			continue
		} else {
			section1, ok1 := numMap[v-1]
			section2, ok2 := numMap[v+1]
			if ok1 && ok2 {
				numMap[section1[0]][1] = section2[1]
				numMap[section2[1]][0] = section1[0]
				numMap[v] = []int{section1[0], section2[1]}
			} else if ok1 {
				numMap[section1[0]][1] = v
				numMap[v] = []int{section1[0], v}
			} else if ok2 {
				numMap[section2[1]][0] = v
				numMap[v] = []int{v, section2[1]}
			} else {
				numMap[v] = []int{v, v}
			}
		}
		ans = max(ans, numMap[v][1]-numMap[v][0]+1)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
//leetcode submit region end(Prohibit modification and deletion)
