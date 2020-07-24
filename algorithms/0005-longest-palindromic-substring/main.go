package main

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	var res string
	for i := range s {
		res = maxPalindromic(s, i, i, res)   // odd:  abcba
		res = maxPalindromic(s, i, i+1, res) // even: abba
	}
	return res
}

func maxPalindromic(s string, i int, j int, res string) string {
	var sub string
	for i >= 0 && j <= len(s)-1 && s[i] == s[j] {
		sub = s[i : j+1]
		i--
		j++
	}
	if len(res) < len(sub) {
		return sub
	}
	return res
}
