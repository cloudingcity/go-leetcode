package main

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	col := len(matrix[0]) - 1
	row := 0

	for col >= 0 && row <= len(matrix)-1 {
		if target == matrix[row][col] {
			return true
		}
		if target > matrix[row][col] {
			row++
		} else {
			col--
		}
	}
	return false
}
