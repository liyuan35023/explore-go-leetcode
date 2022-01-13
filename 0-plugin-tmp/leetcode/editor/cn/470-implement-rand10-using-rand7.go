package cn
//已有方法 rand7 可生成 1 到 7 范围内的均匀随机整数，试写一个方法 rand10 生成 1 到 10 范围内的均匀随机整数。 
//
// 不要使用系统的 Math.random() 方法。 
//
// 示例 1: 
//
//输入: 1
//输出: [7]
//
// 示例 2: 
//
//输入: 2
//输出: [8,4]
//
// 示例 3: 
//
//输入: 3
//输出: [8,1,10]
//
// 提示: 
//
// rand7 已定义。 
// 传入参数: n 表示 rand10 的调用次数。 
//
// 进阶: 
//
// 
// rand7()调用次数的 期望值 是多少 ? 
// 你能否尽量少调用 rand7() ? 
// 

//leetcode submit region begin(Prohibit modification and deletion)
func rand10() int {

	for {
		rand49 := (rand7() - 1) * 7 + rand7()
		if rand49 <= 40 {
			return rand49 % 10 + 1
		}
		rand63 := (rand7() - 1) * 9 + rand49
		if rand63 <= 60 {
			return rand63 % 10 + 1
		}
		rand21 := (rand7() - 1) * 3 + rand63
		if rand21 <= 20 {
			return rand21 % 10 + 1
		}
	}



}
//func rand10() int {
//	 for {
//	 	x, y := rand7(), rand7()
//	 	rand49 := (x-1) * 7 + y
//	 	if rand49 <= 40 {
//	 		return rand49 % 10 + 1
//		}
//		rand9 := rand49 - 40
//		x = rand7()
//		rand63 := (x-1) * 9 + rand9
//		if rand63 <= 60 {
//			return rand63 % 10 + 1
//		}
//		rand3 := rand63 - 60
//		x = rand7()
//		rand21 := (x-1) * 3 + rand3
//		if rand21 <= 20 {
//			return rand21 % 10 + 1
//		}
//	 }
//
//}
//leetcode submit region end(Prohibit modification and deletion)
