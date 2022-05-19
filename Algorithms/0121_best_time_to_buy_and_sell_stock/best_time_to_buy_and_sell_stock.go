package best_time_to_buy_and_sell_stock_0121

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	maxProfit := 0
	left := 0
	right := 1
	length := len(prices)
	for right < length {
		if prices[left] < prices[right] {
			currProfit := prices[right] - prices[left]
			maxProfit = max(maxProfit, currProfit)
		} else {
			left = right
		}
		right += 1
	}
	return maxProfit
}
