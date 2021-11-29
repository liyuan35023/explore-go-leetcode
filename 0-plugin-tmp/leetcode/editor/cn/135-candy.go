package cn
//老师想给孩子们分发糖果，有 N 个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。 
//
// 你需要按照以下要求，帮助老师给这些孩子分发糖果： 
//
// 
// 每个孩子至少分配到 1 个糖果。 
// 评分更高的孩子必须比他两侧的邻位孩子获得更多的糖果。 
// 
// 那么这样下来，老师至少需要准备多少颗糖果呢？
//
// 示例 1： 
//
//输入：[1,0,2]
//输出：5
//解释：你可以分别给这三个孩子分发 2、1、2 颗糖果。
// 
// 示例 2：
//
//输入：[1,2,2]
//输出：4
//解释：你可以分别给这三个孩子分发 1、2、1 颗糖果。
//     第三个孩子只得到 1 颗糖果，这已满足上述两个条件。 

//leetcode submit region begin(Prohibit modification and deletion)
func candy(ratings []int) int {
	candies := make([]int, len(ratings))
	ans := 0
	candies[0] = 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		} else {
			candies[i] = 1
		}
	}
	for i := len(ratings)-2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
			candies[i] = candies[i+1] + 1
		}
	}
	for i := 0; i < len(ratings); i++ {
		ans += candies[i]
	}
	return ans
	//ans := 1
	//pre := 1
	//inc := 1
	//des := 0
	//for i := 1; i < len(ratings); i++ {
	//	if ratings[i] >= ratings[i-1] {
	//		des = 0
	//		if ratings[i] == ratings[i-1] {
	//			pre = 1
	//		} else {
	//			pre++
	//		}
	//		ans += pre
	//		inc = pre
	//	} else {
	//		des++
	//		if des == inc {
	//			des++
	//		}
	//		ans += des
	//		pre = 1
	//	}
	//}
	//return ans
}
//leetcode submit region end(Prohibit modification and deletion)
