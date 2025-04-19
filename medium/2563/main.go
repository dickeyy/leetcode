package main

import "sort"

/*
 *   Stats:
 *   Runtime: 73ms (beats 17.24% of Go solutions)
 *   Memory: 11.07mb (beats 17.24% of Go solutions)
 */

func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	var count int64

	// Use two pointers approach
	for i := range nums {
		// Find the first index where nums[i] + nums[j] >= lower
		left := sort.SearchInts(nums[i+1:], lower-nums[i])
		left += i + 1 // Adjust to the actual index in the original array

		// Find the first index where nums[i] + nums[j] > upper
		right := sort.SearchInts(nums[i+1:], upper-nums[i]+1)
		right += i + 1 // Adjust to the actual index in the original array

		// Add the count of valid pairs for current i
		count += int64(right - left)
	}

	return count
}
