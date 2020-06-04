package main

import (
	"sort"
)

func heightChecker(heights []int) int {
	nums := make([]int, len(heights))
	copy(nums, heights)
	sort.Ints(nums)

	var c int
	for i, num := range nums {
		if num != heights[i] {
			c++
		}
	}
	return c
}
