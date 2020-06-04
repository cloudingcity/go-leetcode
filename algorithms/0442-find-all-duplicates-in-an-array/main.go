package main

func findDuplicates(nums []int) []int {
	var ans []int

	for _, num := range nums {
		idx := abs(num) - 1
		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		} else {
			ans = append(ans, abs(num))
		}
	}
	return ans
}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}
