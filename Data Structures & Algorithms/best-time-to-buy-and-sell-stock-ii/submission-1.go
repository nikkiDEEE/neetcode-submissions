func maxProfit(prices []int) int {
	var profit int
	var i int

	for (i < len(prices)-1) {
		if prices[i+1] > prices[i] {
			profit = profit + prices[i+1] - prices[i]
		}
		i++
	}
	return profit
}

