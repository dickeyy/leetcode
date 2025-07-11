package main

/*
 *   Stats:
 *   Runtime: 0ms (beats 100.00% of Go solutions)
 *   Memory: 9.68mb (beats 82.27% of Go solutions)
 */

func firstMissingPositive(nums []int) int {
	n := len(nums)

	// Step 1: Mark numbers that are out of range [1, n] as n+1
	for i := range nums {
		if nums[i] <= 0 || nums[i] > n {
			nums[i] = n + 1
		}
	}

	// Step 2: Use the array as a hash set
	// For each number x in [1, n], mark nums[x-1] as negative
	for i := range nums {
		val := abs(nums[i])
		if val <= n {
			// Mark nums[val-1] as negative to indicate val is present
			if nums[val-1] > 0 {
				nums[val-1] = -nums[val-1]
			}
		}
	}

	// Step 3: Find the first positive number (index + 1 is the missing positive)
	for i := range nums {
		if nums[i] > 0 {
			return i + 1
		}
	}

	// If all numbers from 1 to n are present, return n+1
	return n + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
