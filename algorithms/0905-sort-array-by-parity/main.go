package main

func sortArrayByParity(A []int) []int {
	var j int
	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			A[j], A[i] = A[i], A[j]
			j++
		}
	}
	return A
}
