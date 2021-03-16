package _09_fibonacci_number

/*
	斐波那契数，通常用 F(n) 表示，形成的序列称为斐波那契数列。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：

	F(0) = 0,   F(1) = 1
	F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
	给定 N，计算 F(N)。

	提示：0 ≤ N ≤ 30
 */

func fib(N int) int {
	// 递归
	if N == 0 || N == 1 {
		return N
	}
	return fib(N-1) + fib(N-2)
}

func fib2(n int) int {
	if n < 2 {
		return n
	}
	first, second := 0, 1
	for i := 2; i <= n; i++ {
		first, second = second, first+second
	}
	return second
}


