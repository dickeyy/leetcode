package main

/*
 *   Stats:
 *   Runtime: 0s (beats 100% of Go solutions)
 *   Memory: 4.52mb (beats 32.65% of Go solutions)
 */

func countPairs(nums []int, k int) int {
	count := 0

	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] && (i*j)%k == 0 {
				count++
			}
		}
	}

	return count
}
