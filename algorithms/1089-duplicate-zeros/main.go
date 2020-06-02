package main

func duplicateZeros(arr []int) {
	var q []int

	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			q = append(q, 0)
		}
		q = append(q, arr[i])
	}
	for i := range arr {
		arr[i] = q[i]
	}
}
