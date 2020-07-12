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

func fibIter(N int) int {
	a, b := 0, 1
	for i := 0; i < N; i++ {
		a, b = b, a+b
	}
	return a
}
