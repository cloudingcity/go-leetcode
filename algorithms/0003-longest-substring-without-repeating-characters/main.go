package main

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	var ans, j int
	m := make(map[int32]int, len(s))
	for i, c := range s {
		if v, ok := m[c]; ok && v >= j {
			j = v + 1
		}
		m[c] = i
		ans = max(ans, i+1-j)
	}
	return ans
}

func lengthOfLongestSubstring2(s string) int {
	var ans, j int
	index := [128]int{}
	for i, c := range s {
		j = max(index[c], j)
		index[c] = i + 1
		ans = max(ans, i+1-j)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
