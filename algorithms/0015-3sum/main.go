package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	var ret [][]int
	length := len(nums)
	if length < 3 {
		return ret
	}

	sort.Ints(nums)

	for i := 0; i < length-2; i++ {
		if nums[i] > 0 {
			break
		} else if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		a := nums[i]
		start := i + 1
		end := length - 1

		for start < end {
			b := nums[start]
			c := nums[end]
			sum := a + b + c

			if sum == 0 {
				ret = append(ret, []int{a, b, c})
			}

			if sum < 0 {
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

	return ret
}
