package main

func permute(nums []int) [][]int {
	var res [][]int
	backtrack(nums, nil, &res)
	return res
}

func backtrack(nums []int, track []int, res *[][]int) {
	if len(track) == len(nums) {
		ans := make([]int, len(nums))
		copy(ans, track)
		*res = append(*res, ans)
		return
	}
	for _, num := range nums {
		if contains(track, num) {
			continue
		}
		track = append(track, num)
		backtrack(nums, track, res)
		track = track[:len(track)-1]
	}
}

func contains(nums []int, target int) bool {
	for _, num := range nums {
		if num == target {
			return true
		}
	}
	return false
}
