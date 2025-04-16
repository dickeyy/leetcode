package main

/*
 *   Stats:
 *   Runtime: 8ms (beats 6.42% of Go solutions)
 *   Memory: 8.78mb (beats 5.96% of Go solutions)
 */

func subsetXORSum(nums []int) int {
	// recursively generate all subsets
	subsets := [][]int{}
	generateSubsets(nums, []int{}, &subsets)

	// calculate the XOR sum of each subset
	xorSum := 0
	for _, subset := range subsets {
		subsetXOR := 0
		for _, num := range subset {
			subsetXOR ^= num
		}
		xorSum += subsetXOR
	}
	return xorSum
}

func generateSubsets(nums []int, current []int, subsets *[][]int) {
	if len(nums) == 0 {
		*subsets = append(*subsets, append([]int{}, current...))
		return
	}

	// include the current number in the subset
	generateSubsets(nums[1:], append(current, nums[0]), subsets)

	// exclude the current number from the subset
	generateSubsets(nums[1:], current, subsets)
}
