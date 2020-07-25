package main

func generateParenthesis(n int) []string {
	var res []string
	backtrack(n, 0, 0, "", &res)
	return res
}

func backtrack(max, left, right int, track string, res *[]string) {
	if len(track) == max*2 {
		*res = append(*res, track)
		return
	}
	if left < max {
		backtrack(max, left+1, right, track+"(", res)
	}
	if left > right {
		backtrack(max, left, right+1, track+")", res)
	}
}
