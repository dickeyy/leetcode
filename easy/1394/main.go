package main

/*
 *   Stats:
 *   Runtime: <1ms (beats 100.00% of Go solutions)
 *   Memory: 4.90mb (beats 52.47% of Go solutions)
 */

func findLucky(arr []int) int {
	freq := make(map[int]int)
	for _, num := range arr {
		freq[num]++
	}

	max := -1
	for num, count := range freq {
		if num == count && num > max {
			max = num
		}
	}

	return max
}
