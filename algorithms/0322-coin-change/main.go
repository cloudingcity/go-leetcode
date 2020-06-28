package main

import (
	"math"
)

func coinChange(coins []int, amount int) int {
	dict := make(map[int]int)
	return dp(coins, amount, dict)
}

func dp(coins []int, amount int, dict map[int]int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	if res, ok := dict[amount]; ok {
		return res
	}

	res := math.MaxInt64
	for _, coin := range coins {
		sub := dp(coins, amount-coin, dict)
		if sub == -1 {
			continue
		}
		res = min(res, 1+sub)
	}

	if res == math.MaxInt64 {
		dict[amount] = -1
	} else {
		dict[amount] = res
	}
	return dict[amount]
}

func coinChange2(coins []int, amount int) int {
	m := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		minCoin := math.MaxInt32
		for _, coin := range coins {
			if i-coin >= 0 && m[i-coin] != -1 {
				minCoin = min(minCoin, m[i-coin]+1)
			}
		}

		if minCoin == math.MaxInt32 {
			m[i] = -1
		} else {
			m[i] = minCoin
		}
	}

	return m[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
