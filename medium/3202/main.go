package main

/*
 *   Stats:
 *   Runtime: 76ms (beats 100.00% of Go solutions)
 *   Memory: 25.23mb (beats 66.67% of Go solutions)
 */

func maximumLength(nums []int, k int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	// dp[i][r] = maximum length of valid subsequence ending at position i
	// where the last pair sum has remainder r
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, k)
	}

	maxLen := 1

	// For each position i
	for i := 1; i < n; i++ {
		// For each previous position j (must be adjacent in subsequence)
		for j := 0; j < i; j++ {
			remainder := (nums[i] + nums[j]) % k

			// Case 1: Start a new subsequence of length 2
			dp[i][remainder] = max(dp[i][remainder], 2)

			// Case 2: Extend an existing subsequence
			// Look for subsequences ending at j with the same remainder
			if dp[j][remainder] > 0 {
				dp[i][remainder] = max(dp[i][remainder], dp[j][remainder]+1)
			}

			maxLen = max(maxLen, dp[i][remainder])
		}
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
