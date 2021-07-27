package main

// Time: O(M*N), Space: O(M*N)
func setZeroes(matrix [][]int) {
	mx := make(map[int]int)
	my := make(map[int]int)
	for x, rows := range matrix {
		for y, v := range rows {
			if v == 0 {
				mx[x] = 0
				my[y] = 0
			}
		}
	}

	for x, rows := range matrix {
		for y, v := range rows {
			if v == 0 {
				continue
			}

			if _, ok := mx[x]; ok {
				matrix[x][y] = 0
			} else if _, ok := my[y]; ok {
				matrix[x][y] = 0
			}
		}
	}
}
