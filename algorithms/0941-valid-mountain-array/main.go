package main

func validMountainArray(A []int) bool {
	var i int
	l := len(A)

	for i = 1; i < l; i++ {
		if A[i-1] >= A[i] {
			break
		}
	}

	if i == 1 || i == l {
		return false
	}

	for ; i < l; i++ {
		if A[i-1] <= A[i] {
			return false
		}
	}

	return true
}
