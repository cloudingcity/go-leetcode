package main

// Time: O(N), Space: O(N)
func subarraySum(nums []int, k int) int {
	var cnt int
	var sum int
	m := make(map[int]int, len(nums)+1)
	m[0] = 1
	for _, num := range nums {
		sum += num
		cnt += m[sum-k]
		m[sum]++
	}
	return cnt
}

// Time: O(N^2), Space: O(1)
func subarraySum2(nums []int, k int) int {
	var cnt int
	var sum int
	for i := 0; i < len(nums); i++ {
		sum = 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				cnt++
			}
		}
	}
	return cnt
}
