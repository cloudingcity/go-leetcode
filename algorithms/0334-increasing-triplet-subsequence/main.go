package main

import (
	"math"
)

func increasingTriplet(nums []int) bool {
	first := math.MaxInt32
	second := math.MaxInt32
	for _, num := range nums {
		switch {
		case num <= first:
			first = num
		case num <= second:
			second = num
		default:
			return true
		}
	}
	return false
}
