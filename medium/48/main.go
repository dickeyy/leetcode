package main

/*
 *   Stats:
 *   Runtime: 0ms (beats 100.00% of Go solutions)
 *   Memory: 4.24mb (beats 21.02% of Go solutions)
 */

func rotate(matrix [][]int) {
	n := len(matrix)

	// step 1: transpose the matrix
	for i := range n {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// step 2: reverse each row
	for i := range n {
		for j := range n / 2 {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
}
