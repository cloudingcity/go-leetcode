package main

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		idx := abs(nums[i]) - 1
		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		}
	}

	var ans []int
	for i, num := range nums {
		if num > 0 {
			ans = append(ans, i+1)
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
