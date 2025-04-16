package main

/*
 *   Stats:
 *   Runtime: <1ms (beats 100.00% of Go solutions)
 *   Memory: 4.46mb (beats 96.01% of Go solutions)
 */

func checkIfExist(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i != j && arr[i] == 2*arr[j] {
				return true
			}
		}
	}
	return false
}
