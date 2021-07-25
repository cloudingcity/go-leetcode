package main

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for start <= end {
		mid := (end-start)/2 + start
		switch {
		case target == nums[mid]:
			return mid
		case nums[mid] >= nums[start]:
			if target >= nums[start] && target < nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		case nums[mid] < nums[start]:
			if target <= nums[end] && target > nums[mid] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return -1
}
