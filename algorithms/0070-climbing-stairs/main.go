package main

var cache = map[int]int{}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if _, ok := cache[n]; ok {
		return cache[n]
	}
	cache[n] = climbStairs(n-1) + climbStairs(n-2)
	return cache[n]
}

func climbStairsIter(n int) int {
	a, b := 1, 2
	for i := 1; i < n; i++ {
		a, b = b, a+b
	}
	return a
}

func climbStairsDP(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
