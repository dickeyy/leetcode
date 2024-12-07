package main

import "github.com/dickeyy/leetcode/utils"

/*
 *   Stats:
 *   Runtime: <1ms (beats 100.00% of Go solutions)
 *   Memory: 4.09mb (beats 10.49% of Go solutions)
 */

func main() {
	utils.AssertEqual(13, scoreOfString("hello"), 1)
	utils.AssertEqual(50, scoreOfString("zaz"), 2)
}

func scoreOfString(s string) int {
	sum := 0
	for i := 0; i < len(s)-1; i++ {
		x := int(s[i]) - int(s[i+1])
		if x < 0 {
			sum += -x
		} else {
			sum += x
		}
	}
	return sum
}
