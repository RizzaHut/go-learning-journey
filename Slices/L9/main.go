package main

func createMatrix(rows, cols int) [][]int {
	matrix := [][]int{}
	for r := 0; r < rows; r++ {
		column := []int{}
		
		for c := 0; c < cols; c++ {
			column = append(column, r * c)
		}
		matrix = append(matrix, column)
	}
	return matrix
}
