package main

var empty struct{}

func checkIfExist(arr []int) bool {
	m := make(map[int]struct{}, len(arr))

	for _, num := range arr {
		if _, ok := m[num*2]; ok {
			return true
		}
		if num%2 == 0 {
			if _, ok := m[num/2]; ok {
				return true
			}
		}
		m[num] = empty
	}
	return false
}
