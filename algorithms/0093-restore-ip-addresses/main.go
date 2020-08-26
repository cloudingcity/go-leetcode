package main

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	if len(s) < 4 || len(s) > 12 {
		return nil
	}
	if len(s) == 4 {
		return []string{strings.Join(strings.Split(s, ""), ".")}
	}
	var res []string
	backtrack(s, &res, nil)
	return res
}

func backtrack(s string, res *[]string, track []string) {
	if len(s) == 0 && len(track) == 4 {
		ans := make([]string, 4)
		copy(ans, track)
		*res = append(*res, strings.Join(ans, "."))
		return
	}
	for i := 1; i <= 3; i++ {
		if len(s) < i {
			continue
		}
		c := s[:i]
		if i >= 2 && (strings.HasPrefix(c, "0") || strings.HasPrefix(c, "00")) {
			continue
		}
		n, _ := strconv.Atoi(c)
		if n > 255 {
			continue
		}
		track = append(track, c)
		backtrack(s[i:], res, track)
		track = track[:len(track)-1]
	}
}
