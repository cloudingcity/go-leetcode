package main

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	var ans string
	for i := range s {
		// odd: abcba
		if odd := maxPalindrome(s, i, i); len(odd) > len(ans) {
			ans = odd
		}
		// even: abba
		if even := maxPalindrome(s, i, i+1); len(even) > len(ans) {
			ans = even
		}
	}
	return ans
}

func maxPalindrome(s string, i, j int) string {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}
	return s[i+1 : j]
}
