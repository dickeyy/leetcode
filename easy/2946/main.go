package main

/*
 *   Stats:
 *   Runtime: 0ms (beats 100.00% of Go solutions)
 *   Memory: 7.03mb (beats 21.74% of Go solutions)
 */

func areSimilar(mat [][]int, k int) bool {
	for i := range mat {
		for j := range mat[i] {
			key := ((j + k) % (len(mat[i])))
			if mat[i][j] != mat[i][key] {
				return false
			}
		}
	}

	return true
}
