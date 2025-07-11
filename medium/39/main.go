package main

/*
 *   Stats:
 *   Runtime: 0ms (beats 100.00% of Go solutions)
 *   Memory: 4.94mb (beats 69.63% of Go solutions)
 */

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var current []int

	backtrack(candidates, target, 0, current, &result)
	return result
}

func backtrack(candidates []int, target int, start int, current []int, result *[][]int) {
	if target == 0 {
		// Make a copy of current to avoid modifying the slice
		combination := make([]int, len(current))
		copy(combination, current)
		*result = append(*result, combination)
		return
	}

	if target < 0 {
		return
	}

	for i := start; i < len(candidates); i++ {
		current = append(current, candidates[i])
		backtrack(candidates, target-candidates[i], i, current, result)
		current = current[:len(current)-1] // backtrack
	}
}
