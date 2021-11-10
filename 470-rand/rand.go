package _70_rand

func rand10() int {
	for {
		rand49 := (rand7() - 1) * 7 + rand7()
		if rand49 <= 40 {
			return rand49 % 10 + 1
		}
		rand9 := rand49 - 40
		rand63 := (rand7() - 1) * 9 + rand9
		if rand63 <= 60 {
			return rand63 % 10 + 1
		}
		rand3 := rand63 - 60
		rand21 := (rand7() - 1) * 3 + rand63
		if rand21 <= 20 {
			return rand21 % 10 + 1
		}
	}
}
