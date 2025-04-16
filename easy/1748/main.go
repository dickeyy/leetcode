package main

/*
 *   Stats:
 *   Runtime: <1ms (beats 100.00% of Go solutions)
 *   Memory: 4.23mb (beats 5.56% of Go solutions)
 */

func sumOfUnique(nums []int) int {
	// count the frequencies of each num
	freq := make(map[int]int)

	for _, num := range nums {
		freq[num]++
	}

	sum := 0
	for i, num := range freq {
		if num == 1 {
			sum += i
		}
	}

	return sum
}
