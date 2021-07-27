package main

func groupAnagrams(strs []string) [][]string {
	var ans [][]string
	m := make(map[string]int)

	for _, str := range strs {
		chars := make([]byte, 26)
		for _, char := range str {
			idx := char - 'a'
			chars[idx] += 1
		}
		key := string(chars)

		if pos, ok := m[key]; !ok {
			ans = append(ans, []string{str})
			m[key] = len(ans) - 1
		} else {
			ans[pos] = append(ans[pos], str)
		}
	}

	return ans
}
