package main

var cache = map[int]int{}

func fib(N int) int {
	if N < 2 {
		return N
	}
	if _, ok := cache[N]; ok {
		return cache[N]
	}
	cache[N] = fib(N-1) + fib(N-2)
	return cache[N]
}

func fibDP(N int) int {
	if N == 0 || N == 1 {
		return N
	}
	dp := make([]int, N+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= N; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[N]
}

func fibIter(N int) int {
	a, b := 0, 1
	for i := 0; i < N; i++ {
		a, b = b, a+b
	}
	return a
}
