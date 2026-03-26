package main

func containsDuplicate(nums []int) bool {
	s := make(map[int]int)
	for _, n := range nums {
		s[n]++
	}
	return len(nums) != len(s)
}
