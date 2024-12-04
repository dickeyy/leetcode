package main

import "github.com/dickeyy/leetcode/utils"

/*
 *   Stats:
 *   Runtime: <1ms (beats 100.00% of Go solutions)
 *   Memory: 3.90mb (beats 35.73% of Go solutions)
 */

func main() {
	utils.AssertEqual(2, sum(1, 1), 1)
}

func sum(num1 int, num2 int) int {
	return num1 + num2
}
