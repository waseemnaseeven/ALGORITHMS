package main

func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	max := -2147783648
	lowest := prices[0]
	for i := 0; i < n; i++ {
		if prices[i]-lowest > max {
			max = prices[i] - lowest
		}
		if prices[i] < lowest {
			lowest = prices[i]
		}
	}
	return max
}
