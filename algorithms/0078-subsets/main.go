package main

func subsets(nums []int) [][]int {
	var res [][]int
	for i := 0; i <= len(nums); i++ {
		backtrack(nums, []int{}, &res, i)
	}
	return res
}

func backtrack(nums, track []int, res *[][]int, max int) {
	if len(track) == max {
		ans := make([]int, len(track))
		copy(ans, track)
		*res = append(*res, ans)
		return
	}
	for _, num := range nums {
		if contains(track, num) {
			continue
		}
		track = append(track, num)
		backtrack(nums, track, res, max)
		track = track[:len(track)-1]
	}
}

func contains(nums []int, target int) bool {
	for _, num := range nums {
		if num == target || num > target {
			return true
		}
	}
	return false
}
