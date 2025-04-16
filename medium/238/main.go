package main

/*
 *   Stats:
 *   Runtime: 6ms (beats 25.59% of Go solutions)
 *   Memory: 11.62mb (beats 9.02% of Go solutions)
 */

func productExceptSelf(nums []int) []int {
	prefix := make([]int, len(nums))
	suffix := make([]int, len(nums))

	prefix[0] = 1
	suffix[len(nums)-1] = 1

	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}

	for i := len(nums) - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i+1]
	}

	for i := range nums {
		nums[i] = prefix[i] * suffix[i]
	}

	return nums
}
