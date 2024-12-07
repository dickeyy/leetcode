package main

import "github.com/dickeyy/leetcode/utils"

/*
 *   Stats:
 *   Runtime: <1ms (beats 100.00% of Go solutions)
 *   Memory: 4.72mb (beats 8.39% of Go solutions)
 */

func main() {
	utils.AssertEqual(6, theMaximumAchievableX(4, 1), 1)
}

func theMaximumAchievableX(num int, t int) int {
	return num + (t * 2)
}
