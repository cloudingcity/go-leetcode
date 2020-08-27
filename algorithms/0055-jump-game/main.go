package main

// time: O(n), space: O(1)
func canJump(nums []int) bool {
	pos := nums[len(nums)-1]
	for i := len(nums) - 1; i >= 0; i-- {
		if i+nums[i] >= pos {
			pos = i
		}
	}
	return pos == 0
}

// time: O(2^n), space: O(n)
func canJumpBacktrack(nums []int) bool {
	return backtrack(0, nums)
}

func backtrack(pos int, nums []int) bool {
	if pos == len(nums)-1 {
		return true
	}
	jump := min(pos+nums[pos], len(nums)-1)
	for i := jump; i > pos; i-- {
		if backtrack(i, nums) {
			return true
		}
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
