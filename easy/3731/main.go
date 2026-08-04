package main

import "slices"

/*
 *   Stats:
 *   Runtime: 0ms (beats 100.00% of Go solutions)
 *   Memory: 5.34mb (beats 41.03% of Go solutions)
 */

func findMissingElements(nums []int) []int {
	missing := []int{}
	slices.Sort(nums[:])
	expected := nums[0]
	for _, num := range nums {
		if num != expected {
			for num != expected {
				missing = append(missing, expected)
				expected++
			}
		}
		expected++
	}
	return missing
}

// func findMissingElements(nums []int) []int {
// 	min, max := nums[0], nums[0]
// 	seen := make(map[int]bool)
// 	for _, num := range nums {
// 		if num < min {
// 			min = num
// 		}
// 		if num > max {
// 			max = num
// 		}
// 		seen[num] = true
// 	}

// 	missing := []int{}
// 	for i := min + 1; i < max; i++ {
// 		if !seen[i] {
// 			missing = append(missing, i)
// 		}
// 	}

// 	return missing
// }
