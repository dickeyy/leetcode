package main

import "github.com/dickeyy/leetcode/utils"

/*
 *   Stats:
 *   Runtime: <1ms (beats 100.00% of Go solutions)
 *   Memory: 8.56mb (beats 9.62% of Go solutions)
 */

func main() {
	utils.AssertEqual([]int{0, 1, 2, 4, 5, 3}, buildArray([]int{0, 2, 1, 5, 3, 4}), 1)
}

func buildArray(nums []int) []int {
	ans := make([]int, len(nums))
	for i, _ := range nums {
		ans[i] = nums[nums[i]]
	}
	return ans
}
