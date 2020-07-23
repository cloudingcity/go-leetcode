package main

import "strconv"

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	var res []string
	backtrack(digits, "", &res)
	return res
}

var letters = [...]string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func backtrack(digits, track string, res *[]string) {
	if len(digits) == 0 {
		*res = append(*res, track)
		return
	}
	n, _ := strconv.Atoi(string(digits[0]))
	for _, c := range letters[n] {
		backtrack(digits[1:], track+string(c), res)
	}
}
