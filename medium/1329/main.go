package main

import (
	"cmp"
	"slices"
)

/*
 *   Stats:
 *   Runtime: 0ms (beats 100.00% of Go solutions)
 *   Memory: 6.74mb (beats 11.11% of Go solutions)
 */

func diagonalSort(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	diag := make(map[int][]int)

	for i := range m {
		for j := range n {
			key := i - j
			diag[key] = append(diag[key], mat[i][j])
		}
	}

	for key, val := range diag {
		slices.SortFunc(val, func(a, b int) int {
			return cmp.Compare(b, a)
		})
		diag[key] = val
	}

	for i := range m {
		for j := range n {
			key := i - j
			mat[i][j] = diag[key][len(diag[key])-1]
			diag[key] = diag[key][:len(diag[key])-1]
		}
	}

	return mat
}
