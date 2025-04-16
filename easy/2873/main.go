package main

/*
 *   Stats:
 *   Runtime: 0ms (beats 100.00% of Go solutions)
 *   Memory: 4.23mb (beats 82.60% of Go solutions)
 */

func maximumTripletValue(nums []int) int64 {
	max := int64(0)

	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				triplet := int64(nums[i]-nums[j]) * int64(nums[k])
				if triplet > max {
					max = triplet
				}
			}
		}
	}

	return max
}
