package main

/*
 *   Stats:
 *   Runtime: <1ms (beats 100.00% of Go solutions)
 *   Memory: 8.56mb (beats 9.62% of Go solutions)
 */

func buildArray(nums []int) []int {
	ans := make([]int, len(nums))
	for i, _ := range nums {
		ans[i] = nums[nums[i]]
	}
	return ans
}
