package dynamic

func maxProfit(prices []int) int {

	if len(prices) < 4 {
		return 0
	}

	max := 0
	for i := 2; i < len(prices)-1; i++ {
		l := maxp(prices[:i])
		if l > 0 {
			r := maxp(prices[i:])
			if r > 0 && l+r > max {
				max = l + r
			}
		}
	}
	return max
}

func maxp(prices []int) int {
	min := prices[0]
	max := 0
	for i, p := range prices {
		if i == 0 {
			min = p
			continue
		}
		if p < min {
			min = p
			continue
		}
		if p-min > max {
			max = p - min
		}
	}
	return max
}
