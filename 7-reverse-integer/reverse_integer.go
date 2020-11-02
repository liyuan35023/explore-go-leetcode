package __reverse_integer

func reverse(x int) int {
	tmp := 0
	for x != 0 {
		if tmp > (1<<31 - 1 - x %10) / 10 || tmp < (-(1<<31)-x%10) / 10 {
			return 0
		}
		tmp = tmp*10 + x % 10
		x = x / 10
	}
	return tmp
}
