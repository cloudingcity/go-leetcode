package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	var ans [][]int
	n := len(nums)
	if n == 0 || n < 3 {
		return ans
	}

	sort.Ints(nums)

	for i := 0; i < n-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		a := nums[i]
		start := i + 1
		end := n - 1

		for start < end {
			b := nums[start]
			c := nums[end]
			if a+b+c == 0 {
				ans = append(ans, []int{a, b, c})
			}

			if a+b+c < 0 {
				start++
				for start < end && nums[start] == nums[start-1] {
					start++
				}
			} else {
				end--
				for start < end && nums[end] == nums[end+1] {
					end--
				}
			}
		}
	}
	return ans
}
