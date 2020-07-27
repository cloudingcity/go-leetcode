package main

import (
	"math"
)

// Time: O((n) Space: O(k)
func maxSlidingWindow(nums []int, k int) []int {
	var queue, res []int
	for i := 0; i < len(nums); i++ {
		// if window pass
		if len(queue) != 0 && queue[0] == i-k {
			queue = queue[1:]
		}
		// if current number is greater than before, remove them
		for len(queue) != 0 && nums[queue[len(queue)-1]] < nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
		// if index >= k-1, can write result
		if i >= k-1 {
			res = append(res, nums[queue[0]])
		}
	}
	return res
}

// Time: O((n-k+1)*k) Space: O(1)
func maxSlidingWindow2(nums []int, k int) []int {
	res := make([]int, len(nums)+1-k)
	max := math.MinInt64

	for i := 0; i <= len(nums)-k; i++ {
		if max > math.MinInt64 && nums[i-1] != max && nums[i+k-1] < max {
			res[i] = max
			continue
		} else {
			max = math.MinInt64
		}
		for _, num := range nums[i : i+k] {
			if num > max {
				max = num
			}
		}
		res[i] = max
	}
	return res
}
