package main

/*
 *   Stats:
 *   Runtime: 5ms (beats 64.29% of Go solutions)
 *   Memory: 5.31mb (beats 71.43% of Go solutions)
 */

func shuffle(nums []int, n int) []int {
	shuffled := make([]int, len(nums))
	for i := range nums {
		if i%2 == 0 {
			shuffled[i] = nums[i/2]
		} else {
			shuffled[i] = nums[n+i/2]
		}
	}
	return shuffled
}
