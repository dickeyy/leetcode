package main

import "github.com/dickeyy/leetcode/utils"

func main() {
	utils.AssertEqual([]int{0, 1}, twoSum([]int{2, 7, 11, 15}, 9), 1)
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
