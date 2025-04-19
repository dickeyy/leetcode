package main

/*
 *   Stats:
 *   Runtime: 0ms (beats 100.00% of Go solutions)
 *   Memory: 8.57mb (beats 83.67% of Go solutions)
 */

func search(nums []int, target int) int {
	h := len(nums) - 1
	l := 0
	for l <= h {
		m := l + (h-l)/2
		if nums[m] == target {
			return m
		} else if nums[m] < target {
			l = m + 1
		} else {
			h = m - 1
		}
	}
	return -1
}
