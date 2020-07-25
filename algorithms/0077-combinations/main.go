package main

func combine(n int, k int) [][]int {
	var res [][]int
	backtrack(n, k, 1, nil, &res)
	return res
}

func backtrack(n, k, start int, track []int, res *[][]int) {
	if k == 0 {
		ans := make([]int, len(track))
		copy(ans, track)
		*res = append(*res, ans)
		return
	}
	for i := start; i <= n-k+1; i++ {
		track = append(track, i)
		backtrack(n, k-1, i+1, track, res)
		track = track[:len(track)-1]
	}
}
