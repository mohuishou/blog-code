package stack

func calculate(s string) int {
	var (
		prev, res int
		num       []int
		operator  []rune
	)
	for _, v := range s {
		if v >= '0' && v <= '9' {
			prev = prev*10 + int(v-'0')
			continue
		}
		num = append(num, prev)

	}
	return res
}
