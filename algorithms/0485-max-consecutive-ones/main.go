package main

func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	cur := 0
	for _, num := range nums {
		switch num {
		case 1:
			cur++
			if cur > max {
				max = cur
			}
		case 0:
			cur = 0
		}
	}
	return max
}
