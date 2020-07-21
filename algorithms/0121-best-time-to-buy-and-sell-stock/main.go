package main

import "math"

func maxProfit(prices []int) int {
	min := math.MaxInt64
	max := 0
	for _, price := range prices {
		if price < min {
			min = price
		} else if price-min > max {
			max = price - min
		}
	}
	return max
}
