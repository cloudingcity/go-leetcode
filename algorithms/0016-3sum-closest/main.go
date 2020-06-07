package main

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)
	ans := nums[0] + nums[1] + nums[2]

	for i := 0; i < n-2; i++ {
		a := nums[i]
		start := i + 1
		end := n - 1

		for start < end {
			b := nums[start]
			c := nums[end]
			sum := a + b + c

			if sum == target {
				return sum
			} else if abs(target-ans) > abs(target-sum) {
				ans = sum
			}

			if sum > target {
				end--
			} else {
				start++
			}
		}
	}

	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
