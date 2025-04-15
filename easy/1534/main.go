package main

/*
 *   Stats:
 *   Runtime: <26ms (beats 6.89% of Go solutions)
 *   Memory: 4.11mb (beats 13.19% of Go solutions)
 */

func countGoodTriplets(arr []int, a int, b int, c int) int {
	count := 0
	n := len(arr)
	for i := range n {
		for j := range n {
			for k := range n {
				if i < j && j < k &&
					abs(float64(arr[i]-arr[j])) <= float64(a) &&
					abs(float64(arr[j]-arr[k])) <= float64(b) &&
					abs(float64(arr[i]-arr[k])) <= float64(c) {
					count++
				}
			}
		}
	}
	return count
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
