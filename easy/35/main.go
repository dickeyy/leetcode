package main

import "github.com/dickeyy/leetcode/utils"

/*
 *   Stats:
 *   Runtime: <1ms (beats 100.00% of Go solutions)
 *   Memory: 4.92mb (beats 19.90% of Go solutions)
 */

func main() {
	utils.AssertEqual(2, searchInsert([]int{1, 3, 5, 6}, 5), 1)
}

func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := (l + r) / 2

		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			return mid
		}
	}

	return l
}
