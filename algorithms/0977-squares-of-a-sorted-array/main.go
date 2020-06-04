package main

func sortedSquares(A []int) []int {
	ans := make([]int, len(A))
	i := 0
	j := len(A) - 1
	cur := j
	for i <= j {
		if abs(A[i]) < abs(A[j]) {
			ans[cur] = A[j] * A[j]
			j--
		} else {
			ans[cur] = A[i] * A[i]
			i++
		}
		cur--
	}
	return ans
}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}
