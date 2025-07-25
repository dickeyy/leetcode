package main

/*
 *   Stats:
 *   Runtime: 0ms (beats 100.00% of Go solutions)
 *   Memory: 4.41mb (beats 12.09% of Go solutions)
 */

func maxSum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	seen := make(map[int]bool)
	maxSum := nums[0]
	currentSum := 0
	hasPositive := false

	for _, num := range nums {
		if !seen[num] {
			seen[num] = true
			if num > 0 {
				currentSum += num
				hasPositive = true
			}
			if num > maxSum {
				maxSum = num
			}
		}
	}

	if hasPositive {
		return currentSum
	}
	return maxSum
}
