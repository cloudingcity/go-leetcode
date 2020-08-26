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

func climbStairsAll(n int) [][]int {
	var res [][]int
	backtrack(n, &res, nil)
	return res
}

func backtrack(n int, res *[][]int, track []int) {
	if n == 0 {
		ans := make([]int, len(track))
		copy(ans, track)
		*res = append(*res, ans)
		return
	}

	for i := 1; i <= 2; i++ {
		if n-i >= 0 {
			track = append(track, i)
			backtrack(n-i, res, track)
			track = track[:len(track)-1]
		}
	}
}
