package main

/*
 *   Stats:
 *   Runtime: 0ms (beats 100.00% of Go solutions)
 *   Memory: 4.86mb (beats 21.47% of Go solutions)
 */

func permute(nums []int) [][]int {
	perms := [][]int{}
	permuteHelper(nums, []int{}, &perms)
	return perms
}

func permuteHelper(nums []int, curr []int, perms *[][]int) {
	if len(nums) == 0 {
		*perms = append(*perms, curr)
		return
	}

	for i := range nums {
		next := make([]int, len(nums))
		copy(next, nums)
		next = append(next[:i], next[i+1:]...)
		permuteHelper(next, append(curr, nums[i]), perms)
	}
}
