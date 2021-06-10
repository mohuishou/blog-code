package stack

func removeKdigits(num string, k int) string {
	for next := []byte{}; k > 0 && len(next) < len(num); num = string(next) {
		i := 0
		for ; k > 0 && i < len(num)-1; i++ {
			// 首位数字 0 直接丢弃
			if len(next) == 0 && num[i] == '0' {
				continue
			}
			if num[i] > num[i+1] {
				k = k - 1
			} else {
				next = append(next, num[i])
			}
		}
		if i < len(num) {
			next = append(next, num[i:]...)
		}
	}

	// 移除首位 0
	i := 0
	for ; i < len(num) && num[i] == '0'; i++ {
	}
	num = num[i:]

	if len(num) <= k || i == len(num) {
		return "0"
	}

	return num[k:]
}
