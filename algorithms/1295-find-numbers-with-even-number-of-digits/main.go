package main

import "strconv"

func findNumbers(nums []int) int {
	var res int
	for _, num := range nums {
		if len(strconv.Itoa(num))%2 == 0 {
			res++
		}
	}
	return res
}
