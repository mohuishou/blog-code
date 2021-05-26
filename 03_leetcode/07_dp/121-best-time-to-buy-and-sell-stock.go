package dp

func maxProfit(prices []int) int {
	var stack []int
	var res int
	for _, price := range prices {
		if len(stack) == 0 || stack[len(stack)-1] < price {
			stack = append(stack, price)
			continue
		}
		i := len(stack) - 1
		for ; i >= 0 && stack[i] >= price; i-- {
			if stack[i]-stack[0] > res {
				res = stack[i] - stack[0]
			}
		}
		i++
		stack = stack[:i]
		stack = append(stack, price)
	}
	if stack[len(stack)-1]-stack[0] > res {
		res = stack[len(stack)-1] - stack[0]
	}
	return res
}
